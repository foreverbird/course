pragma solidity ^0.4.21;


contract BirdManager {
    // 定义bird结构体
    struct Bird {
        bytes32 genes;
        uint64 id;
        uint64 birthTime;
        address owner;
        uint64 coin;
        uint32 exp;
        uint32 power;
        uint32 speed;
        uint32 level;
        uint64 eatFruitTime;
    }

    struct Ownership {
        uint64 count;
        // <tokenid:tokenid>，为了快速查找
        mapping(uint64 => uint64) indexMap;
    }
    // 管理员账户
    address internal ceo;
    bool public paused = false;
    // 总bird数量
    uint64 public totalBirds = 0;
    // 盐值
    string salt = "forever_bird";
    // <tokenid:Bird>
    mapping(uint64 => Bird) internal birdsMap;
    // <拥有者:tokenids>
    mapping(address => Ownership) internal ownerMap;
    
    event create(address buyerAddress, uint64 BirdId, uint256 amount, uint64 currentTime);
    event Trade(address buyerAddress, address sellerAddress, uint64 BirdId, uint256 price, uint256 fee, uint64 currentTime);
    event feedFruitEvent(address owner, uint64 BirdId, uint8 fruitId, uint256 amount, uint64 currentTime);
    event PKEvent(address sender, uint64 challengerId, uint64 resisterId, bool isWin, 
        uint32 challengerRewardExp, uint32 resisterRewardExp, uint64 winnerRewardCoin, uint64 attack1, uint64 attack2, uint64 currentTime);
        
   
    modifier onlyCEO() {
        require(msg.sender == ceo);
        _;
    }

    modifier notPaused() {
        require(!paused);
        _;
    }

    function pause()
    external
    onlyCEO
    {
        paused = true;
    }

    // @dev 开启合约
    function unpause()
    external
    onlyCEO
    {
        paused = false;
    }

    function BirdManager()
    public
    {
        ceo = msg.sender;
    }

    // 生成基因序列
    function _generateGenes(uint64 _birthTime)
        internal 
        returns(uint64 id, bytes32 genes)
    {
        //获取 tokenid
        uint64 _id = totalBirds + 1;
        totalBirds += 1;
        bytes32 _genes = keccak256(_birthTime, _id, salt);
        return (_id, _genes);
    }
    // 获取5~15随机值
    function _getNumber(byte b)
    internal pure
    returns (uint8)
    {
        uint8 p = (uint8(b & 0x0f));
        if (p < 5) {
            p = p + 5;
        }
        return p;
    }
    // 创建bird对象
    function _createBird(address _ownerAddress)
    internal
	returns (uint64 id)
    {
        //获取tokenid
        uint64 _id;
        bytes32 _genes;
        uint64 _birthTime = (uint64)(now);
        (_id, _genes) = _generateGenes(_birthTime);
        Bird memory _bird = Bird({
            id : _id,
            owner: _ownerAddress,
            genes : _genes,
            birthTime : _birthTime,
            coin : 1000,
            exp : 0,
            power : _getNumber(_genes[0]),
            speed : _getNumber(_genes[1]),
            level : 0,
            eatFruitTime: 0
            });
        birdsMap[_id] = _bird;
		return _id;
    }

    function _lossBird(address _ownerAddress, uint64 _id)
    private
    {
        Ownership storage owner = ownerMap[_ownerAddress];

        delete owner.indexMap[_id];
        delete birdsMap[_id].owner;

        owner.count -= 1;
    }

    // 添加bird所有权
    function _ownBird(address _ownerAddress, uint64 _id)
    private
    {
        Ownership storage owner = ownerMap[_ownerAddress];
        if(owner.count == 0) {
            ownerMap[_ownerAddress] = Ownership({
                count:0
            });
        }
        owner.count++;
        owner.indexMap[_id] = _id;
    }

    function _ownershipTransfer(
        address _from,
        address _to,
        uint64 _id
    )
    internal
    {
        // bird所有权转移时，原主人与新主人地址不为空
        require(_from != address(0));
        require(_to != address(0));
       
        // 原主人失去所有权
        _lossBird(_from, _id);
        // 新主人获得所有权
        _ownBird(_to, _id);
        // 变更bird主人信息为新主人
        birdsMap[_id].owner = _to;
    }

    function _getBaseLevelExp(uint32 level)
    internal pure
    returns (uint32, uint16)
    {
        if(level < 30)
        {
            if(level < 10)
            {
                return (level * 50, 50);
            }
            else
            {
                return (500 + (level - 10) * 100, 100);
            }
        }
        else
        {
            if(level < 80)
            {
                return (2500 + (level - 30) * 200, 200);
            }
            else
            {
                return (12500 + (level - 80) * 300, 300);
            }
        }
    }

    function _addPower(uint64 _id, uint8 rate_n, uint8 rate_d)
    internal
    {
        Bird storage _bird = birdsMap[_id];
        
        //增加5%力量
        uint32 newPower = uint32(_bird.power * rate_n / rate_d);
        if(newPower == _bird.power)
        {
            _bird.power = _bird.power + 1;
        }
        else
        {
            _bird.power = newPower;
        }
    }

    function _addSpeed(uint64 _id, uint8 rate_n, uint8 rate_d)
    internal 
    {
        Bird storage _bird = birdsMap[_id];
        
        //增加5%力量
        uint32 newSpeed = uint32(_bird.speed * rate_n / rate_d);
        if(newSpeed == _bird.speed)
        {
            _bird.speed = _bird.speed + 1;
        }
        else
        {
            _bird.speed = newSpeed;
        }
    }

    function _addExp(uint64 _id,  uint32 _deltaExp)
    internal 
    {
        Bird storage _bird = birdsMap[_id];
        _bird.exp += _deltaExp;
        
        uint32 baseExp;
        uint16 step;
        (baseExp, step) = _getBaseLevelExp(_bird.level);
        if(_bird.exp - baseExp >= step)
        {
            _bird.level += 1;
            //增加3%力量
            _addPower(_id, 103, 100);
            //增加3%速度
            _addSpeed(_id, 103, 100);
            //增加100 coin
            _bird.coin += 100;
        }
    }

    function _transferWeight(uint64 _fromBirdId, uint64 _toBirdId, uint8 rate_n, uint8 rate_d)
    internal 
    returns (uint64)
    {
        Bird storage _fromBird = birdsMap[_fromBirdId];
        Bird storage _toBird = birdsMap[_toBirdId];
        
        if(_fromBird.coin == 0)
        {
            return 0;
        }

        uint64 deltaCoin = uint64(_fromBird.coin * rate_n / rate_d);
        if(deltaCoin == 0)
        {
            deltaCoin = 1;
        }

        if(deltaCoin > _fromBird.coin)
        {
            deltaCoin = _fromBird.coin;
        }

        _fromBird.coin -= deltaCoin;
        _toBird.coin += deltaCoin;

        return deltaCoin;
    }

    function _reward(uint64 _challengerId, uint64 _resisterId, bool _isWin)
    internal 
    returns (uint32, uint32, uint64)
    {
        Bird memory _c_bird = birdsMap[_challengerId];
        Bird memory _r_bird = birdsMap[_resisterId];

        uint64 deltaWeight = 0;
        int32 deltaLevel = int32(_c_bird.level - _r_bird.level);
        if(deltaLevel >= 0)
        {
            if(deltaLevel > 10)
            {
                if(_isWin)
                {
                    //只奖励经验值
                    _addExp(_challengerId, 10);
                    _addExp(_resisterId, 10);
                    return (10, 10, 0);
                }
                else
                {
                    _addExp(_challengerId, 10);
                    _addExp(_resisterId, 30);
                    deltaWeight = _transferWeight(_challengerId, _resisterId, 5, 100);
                    return (10, 30, deltaWeight);
                }
            }
            else
            {
                if(_isWin)
                {
                    _addExp(_challengerId, 20);
                    _addExp(_resisterId, 10);
                    deltaWeight = _transferWeight(_resisterId, _challengerId, 3, 100);
                    return (20, 10, deltaWeight);
                }
                else
                {
                    _addExp(_challengerId, 10);
                    _addExp(_resisterId, 20);
                    deltaWeight = _transferWeight(_challengerId, _resisterId, 3, 100);
                    return (10, 20, deltaWeight);
                }
            }
        }
        else
        {
            if(deltaLevel > -10)
            {
                if(_isWin)
                {
                    _addExp(_challengerId, 30);
                    _addExp(_resisterId, 10);
                    deltaWeight = _transferWeight(_resisterId, _challengerId, 5, 100);
                    return (30, 10, deltaWeight);
                }
                else
                {
                    _addExp(_challengerId, 10);
                    _addExp(_resisterId, 30);
                    deltaWeight = _transferWeight(_challengerId, _resisterId, 3, 100);
                    return (10, 30, deltaWeight);
                }
            }
            else
            {
                if(_isWin)
                {
                    _addExp(_challengerId, 50);
                    _addExp(_resisterId, 10);
                    deltaWeight = _transferWeight(_resisterId, _challengerId, 5, 100);
                    return (50, 10, deltaWeight);
                }
                else
                {
                    _addExp(_challengerId, 10);
                    _addExp(_resisterId, 10);
                    return (10, 10, 0);
                }
            }
        }
    }

    /**
     ** 为鸟吃水果
     ** @param _id       鸟的编号
     ** @param _fruitType   水果类型
     */
    function _eatFruit(uint64 _id, uint8 _fruitType) 
    internal 
    {
        require(_fruitType >= 1 && _fruitType <= 3);
        Bird storage _bird = birdsMap[_id];
        
        if(_fruitType == 1)
        {
            //增加5%力量
            _addPower(_id, 105, 100);
        }
        else if(_fruitType == 2)
        {
            //增加5%速度
            _addSpeed(_id, 105, 100);
        }
        else
        {
            //增加50经验值
            _addExp(_id, 50);
        }

        _bird.eatFruitTime = (uint64)(now);
    }

    function _getBirdAttack(uint64 _id, uint64 _anotherBirdId) 
    internal view
    returns (uint64, uint64)
    {
        Bird memory _Bird = birdsMap[_id];
        Bird memory _anotherBird = birdsMap[_anotherBirdId];
		
        uint64 attack1 = (_Bird.power * 10 * 2) / _anotherBird.power + (_Bird.coin * 10) / _anotherBird.coin + (_Bird.speed * 10) / _anotherBird.speed;
        uint64 attack2 = (_anotherBird.power * 10 * 2) / _Bird.power + (_anotherBird.coin * 10) / _Bird.coin + (_anotherBird.speed * 10) / _Bird.speed;

        //PKAttackEvent(attack1, attack2);

        uint64 t = (uint64)(now);
        bytes32 _genes = keccak256(t, salt);
        attack1 = attack1 * (100 + _getNumber(_genes[0]));
        attack2 = attack2 * (100 + _getNumber(_genes[1]));
        
        return (attack1, attack2);
    }

    function catchBird()
    external
    payable
    notPaused
    {
		// 用于购买bird的资金不可以低于 0.03 eth
        require(msg.value >= 30000000000000000);
        uint64 t = (uint64)(now);
		//创建Bird
		uint64 _id = _createBird(msg.sender);
		// 记录bird主人地址
        _ownBird(msg.sender, _id);
		//向合约地址转账
        ceo.transfer(msg.value);
		//记录事件
		emit create(msg.sender, _id, msg.value, t);
    }

    // 查询bird数据
    function getBirdInfo(uint64 _id)
    view
    external
    notPaused
    returns (
        bytes32 genes,
        uint64 birthTime,
        address owner,
        uint64 coin,
        uint32 exp,
        uint32 power,
        uint32 speed,
		uint32 level,
        uint64 eatFruitTime
    )
    {
        // 校验, bird需存在
        require(birdsMap[_id].genes != 0);

        Bird memory _bird = birdsMap[_id];

        genes = _bird.genes;
		birthTime = _bird.birthTime;
		owner = _bird.owner;
		coin = _bird.coin;
		exp = _bird.exp;
		power = _bird.power;
		speed = _bird.speed;
		level = _bird.level;
        eatFruitTime = _bird.eatFruitTime;

        return (genes,
                birthTime,
				owner,
				coin,
                exp,
                power,
                speed,
                level,
                eatFruitTime);
    }

    function trade(
        uint64 tokenId,
        uint256 price,
        uint256 fee,
        bytes32 sign
    )
    external
    payable
    notPaused
    {
        // 校验签名
        bytes32 s = keccak256(tokenId, price, fee, salt);
        //require(sign == s);

        // bird存在
        require(birdsMap[tokenId].genes != 0);

        // 买家不是bird主人，不能购买自己的
        require(birdsMap[tokenId].owner != msg.sender);

        // 用于购买bird的资金不可以为 0
        require(msg.value > 0);
        require(price > fee);

        //value == price + fee
        require(msg.value == (price + fee));

        //获取主人地址
        address seller = birdsMap[tokenId].owner;

        // bird所有权转移
        _ownershipTransfer(seller, msg.sender, tokenId);

        // 将购买资金转入bird主人地址
        seller.transfer(price);

        //向合约地址转账
        ceo.transfer(fee);

        emit Trade(msg.sender, seller, tokenId, price, fee, (uint64)(now));
    }

    function feedFruit(
        uint64 tokenId,
        uint8 fruitId,
        bytes32 sign
    )
    external
    payable
    notPaused
    {
        require(msg.value > 0);

        // 校验签名
        bytes32 s = keccak256(tokenId, fruitId, uint256(msg.value), salt);
        require(sign == s);

        // bird存在
        require(birdsMap[tokenId].genes != 0);

        // 买家是bird主人
        require(birdsMap[tokenId].owner == msg.sender);

        //24h内只能吃一次水果
        uint64 t = (uint64)(now);
        require((t - birdsMap[tokenId].eatFruitTime) > 1 days);

        _eatFruit(tokenId, fruitId);

		// 向合约地址转账
        ceo.transfer(msg.value);

		// 记录事件
        emit feedFruitEvent(msg.sender, tokenId, fruitId, msg.value, (uint64)(now));
    }

    function pk(
        uint64 challengerId,
        uint64 resisterId,
        bytes32 sign
    )
    external
    payable
    notPaused
    {
        // 校验签名
        bytes32 s = keccak256(challengerId, resisterId, salt);
        require(sign == s);

        require(birdsMap[challengerId].owner == msg.sender);
        require(birdsMap[resisterId].genes != 0);

        require(birdsMap[challengerId].owner != birdsMap[resisterId].owner);

        bool isWin = true;
        uint32 exp1; 
        uint32 exp2;
        uint64 deltaCoin;
        uint64 attack1;
        uint64 attack2;
        (attack1, attack2) = _getBirdAttack(challengerId, resisterId);
        if(attack1 >= attack2)
        {
            (exp1, exp2, deltaCoin) = _reward(challengerId, resisterId, true);
        }
        else
        {
            isWin = false;
            (exp1, exp2, deltaCoin) = _reward(challengerId, resisterId, false);
        }
        
        emit PKEvent(msg.sender, challengerId, resisterId, isWin, exp1, exp2, deltaCoin, attack1, attack2, (uint64)(now));
    }
}
