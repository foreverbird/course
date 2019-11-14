package logic

import (
	"bird/dao"
	"bird/data"
	"bird/util"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func QueryCatchBirds(address string, offset int, count int) (*[]data.CatchTxn, int) {
	catchBirds, err := dao.QueryUnConfirmAndOnChainWithLimit(address, offset, count)
	if !err {
		fmt.Println("query catch bird db failed...")
	}
	return catchBirds, util.ERROR_CODE_SUCCESS
}

func QueryUserNormalBirds(address string, startPage int, pageSize int,) (*map[string]interface{}, int) {
	if startPage <= 0 || (pageSize <= 0) {
		util.Logger.Error("query normal birds count failed.")
		return nil, util.ERROR_CODE_PARAM_INVALID
	}

	//获取market bird数量
	count, err := dao.QueryUserNormalBirdsCount(address)
	if !err {
		util.Logger.Error("query normal birds failed.")
		return nil, util.ERROR_CODE_MARKET_BIRDS
	}

	var result = make(map[string]interface{})
	result["total_number"] = count
	result["start_page"] = startPage
	result["page_size"] = pageSize

	if (count == 0) || ((startPage-1)*pageSize >= count) {
		result["birds"] = []data.Bird{}
		return &result, util.ERROR_CODE_SUCCESS
	}

	size := pageSize
	if pageSize > 50 {
		size := 50
		result["page_size"] = size
	}

	birds, err := dao.QueryNormalBirdsByAddressWithLimit(address, (startPage-1)*pageSize, size)
	if birds == nil {
		util.Logger.Error("dao.QueryNormalBirdsByAddressWithLimit is nil")
		result["birds"] = []data.Bird{}
		return &result, util.ERROR_CODE_SUCCESS
	}

	result["birds"] = birds
	return &result, util.ERROR_CODE_SUCCESS
}

func CatchBirdHash(address string, hash string) (err int) {
	//存入catch表中，记录使用
	id := dao.InsertCatchRecord(address, hash)
	if id == 0 {
		//插入错误，要写入日志
		util.Logger.Error("catchbird db error, address->hash:", address, "->", hash)
	}
	return 0
}

func GetPKListBirds(address string, startPage int, pageSize int, orderType int, sortType int) (*map[string]interface{}, int) {
	if startPage <= 0 || (pageSize <= 0) {
		util.Logger.Error("query pk bird count failed.")
		return nil, util.ERROR_CODE_PARAM_INVALID
	}

	//获取pk bird数量
	count, err := dao.QueryPKListBirdsCount(address)
	if !err {
		util.Logger.Error("query pk bird count failed.")
		return nil, util.ERROR_CODE_MARKET_BIRDS
	}

	var pkBirds = make(map[string]interface{})
	pkBirds["total_number"] = count
	pkBirds["start_page"] = startPage
	pkBirds["page_size"] = pageSize

	if (count == 0) || ((startPage-1)*pageSize >= count) {
		pkBirds["birds"] = []data.Bird{}
		return &pkBirds, util.ERROR_CODE_SUCCESS
	}

	size := pageSize
	if pageSize > 50 {
		size := 50
		pkBirds["page_size"] = size
	}

	birds, err := dao.QueryPKListBirds(address, (startPage-1)*pageSize, size)
	if birds == nil {
		util.Logger.Error("dao.QueryPKListBirds birdAuction is nil")
		pkBirds["birds"] = []data.Bird{}
		return &pkBirds, util.ERROR_CODE_SUCCESS
	}

	pkBirds["birds"] = birds
	return &pkBirds, util.ERROR_CODE_SUCCESS
}

func SellBird(tokenId int64, address string, price string) (int) {
	bird := dao.QueryBirdByAddressAndTokenId(address, tokenId)
	if bird == nil {
		return util.ERROR_CODE_NO_BIRD
	}

	//当前address是否拥有该slot
	if strings.Compare(address, bird.Owner) != 0 {
		return util.ERROR_CODE_OP_OTHER_BIRD
	}

	//判断当前状态
	if bird.Status != data.Bird_Status_Normal {
		return util.ERROR_CODE_SELL_BIRD_STATUS
	}

	//判断db中是否有，如果有应该删除
	birdAuction, err := dao.QueryMarketBirdByTokenId(tokenId)
	if err && (birdAuction != nil) {
		return util.ERROR_CODE_SELLED_BIRD
	}

	auctionBegin := time.Now().Second()
	auctionEnd := auctionBegin + 24*60*60

	//拍卖信息存入数据库
	id := dao.InsertMarketBird(tokenId, address, bird.Rarity, price, auctionBegin, auctionEnd)
	if id == 0 {
		return util.ERROR_CODE_DB_FAILED
	}

	//更新数据库
	err = dao.UpdateBirdStatus(tokenId, data.Bird_Status_Selling)
	if !err {
		util.Logger.Error("update bird status failed", tokenId)
	}

	return util.ERROR_CODE_SUCCESS
}

func SellOff(tokenId int64, address string) (int) {
	bird := dao.QueryBirdByAddressAndTokenId(address, tokenId)
	if bird == nil {
		return util.ERROR_CODE_NO_BIRD
	}

	//当前address是否拥有该slot
	if strings.Compare(address, bird.Owner) != 0 {
		return util.ERROR_CODE_OP_OTHER_BIRD
	}

	//slot状态
	if bird.Status == data.Bird_Status_Selling {
		//更新数据库
		err := dao.UpdateBirdStatus(tokenId, data.Bird_Status_Normal)
		if !err {
			return util.ERROR_CODE_DB_FAILED
		}
		//删除market中bird info
		dao.DeleteMarketBirdByTokenId(tokenId)
		return util.ERROR_CODE_SUCCESS
	} else {
		//状态不对
		util.Logger.Error("SellOffSlot slot.Status:", bird.Status)
		return util.ERROR_CODE_BIRD_STATUS
	}
}


/**
 ** 通过bird token id 获取bird信息
 **/
func QueryBirdById(tokenId int64) (*data.Bird, int) {
	bird := dao.QueryBirdByTokenId(tokenId)
	if bird == nil {
		return nil, util.ERROR_CODE_NO_BIRD
	}

	if bird.Status == data.Bird_Status_Selling {
		//查询拍卖信息
		birdAuction, err := dao.QueryMarketBirdByTokenId(tokenId)
		if err {
			bird.Auction = birdAuction
		} else {
			bird.Auction = nil
		}
	}

	return bird, util.ERROR_CODE_SUCCESS
}

/**
 ** 通过用户地址获取用户所有birds
 */
func QueryBirdsByAddress(address string, startPage int, pageSize int) (*map[string]interface{}, int) {
	// 分页参数是否有效
	if startPage <= 0 || (pageSize <= 0) {
		return nil, util.ERROR_CODE_PARAM_INVALID
	}

	//获取user catch bird 数量
	catchCount, e := dao.QueryUnConfirmAndOnChainCount(address)
	if !e {
		return nil, util.ERROR_CODE_QUERY_CATCH_BIRDS_COUNT
	}

	//获取user bird数量
	count, err := dao.QueryUserBirdsCount(address)
	if !err {
		return nil, util.ERROR_CODE_QUERY_USER_BIRDS_COUNT
	}

	// 总数量
	totalCount := catchCount + count

	var ownerBirds = make(map[string]interface{})
	ownerBirds["total_number"] = totalCount
	ownerBirds["start_page"] = startPage
	ownerBirds["page_size"] = pageSize

	// 判断分页范围是否有数据
	if (totalCount == 0) || ((startPage-1)*pageSize >= totalCount) {
		ownerBirds["birds"] = []data.Bird{}
		ownerBirds["catch_birds"] = []data.CatchTxn{}
		return &ownerBirds, util.ERROR_CODE_SUCCESS
	}

	size := pageSize
	if pageSize > 50 {
		size := 50
		ownerBirds["page_size"] = size
	}

	realCatchOffset := 0
	realCatchCount := 0
	realUserOffset := 0
	realUserCount := 0
	// 计算catch表中查询bird的偏移和数量
	if catchCount > 0 {
		realCatchOffset = (startPage - 1) * pageSize
		if realCatchOffset > catchCount {
			realCatchOffset = 0
			realCatchCount = 0
		} else {
			realCatchCount = catchCount - realCatchOffset
			if realCatchCount > size {
				realCatchCount = size
			}
		}
	}

	// 计算bird表中查询bird的偏移和数量
	if count > 0 {
		if realCatchCount < size {
			if realCatchCount == 0 {
				realUserOffset = (startPage-1)*pageSize - catchCount
			} else {
				realUserOffset = 0
			}

			realUserCount = size - realCatchCount
		}
	}

	// 查询catch表获取bird列表
	if realCatchCount == 0 {
		ownerBirds["catch_birds"] = []data.CatchTxn{}
	} else {
		catchBirds, err := QueryCatchBirds(address, realCatchOffset, realCatchCount)
		if err != util.ERROR_CODE_SUCCESS {
			return nil, err
		} else {
			ownerBirds["catch_birds"] = catchBirds
		}
	}

	// 查询bird表获取bird列表
	if realUserCount == 0 {
		ownerBirds["birds"] = []data.Bird{}
	} else {
		birds, err := QueryUserBirds(address, realUserOffset, realUserCount)
		if err != util.ERROR_CODE_SUCCESS {
			return nil, err
		} else {
			ownerBirds["birds"] = birds
		}
	}

	return &ownerBirds, util.ERROR_CODE_SUCCESS
}

func QueryUserBirds(address string, offset int, count int) (*[]data.Bird, int) {
	// 根据用户地址筛选bird列表
	birds, err := dao.QueryBirdsByAddressWithLimit(address, offset, count)
	if !err {
		return nil, util.ERROR_CODE_DB_FAILED
	}

	if birds == nil {
		return &[]data.Bird{}, util.ERROR_CODE_SUCCESS
	}

	// 使用birdIndex记录bird id列表，然后通过bird id列表查询market数据库
	var keys []int64
	var birdMap = make(map[int64]*data.Bird)
	for birdIndex, bird := range *birds {
		birdMap[bird.TokenId] = &(*birds)[birdIndex]
		keys = append(keys, bird.TokenId)
	}

	// 通过bird id列表查询市场bird信息
	birdAuctions, err := dao.QueryMarketBirdByTokenIds(keys)
	if !err {
		return nil, util.ERROR_CODE_DB_FAILED
	}

	// 返回的bird结构体中填充bird拍卖信息
	if birdAuctions != nil {
		for index, auction := range *birdAuctions {
			tokenId := auction.TokenId
			birdMap[tokenId].Auction = &((*birdAuctions)[index])
		}
	}

	return birds, util.ERROR_CODE_SUCCESS
}

func BuyBird(address string, tokenId int64) (*data.BuyBird, int) {
	//challenger bird存在
	bird := dao.QueryBirdByTokenId(tokenId)
	if bird == nil {
		return nil, util.ERROR_CODE_NO_BIRD
	}

	//本人操作
	if strings.Compare(bird.Owner, address) == 0 {
		return nil, util.ERROR_CODE_BUY_SELF_BIRD
	}

	// bird必须为销售状态
	if bird.Status != data.Bird_Status_Selling {
		return nil, util.ERROR_CODE_NOT_SELL_BIRD
	}

	// 查询market表获取bird的拍卖信息
	birdAuction, _ := dao.QueryMarketBirdByTokenId(tokenId)
	if birdAuction == nil {
		return nil, util.ERROR_CODE_NOT_SELL_BIRD_INFO
	}

	// 计算费率
	fee := util.CountTransactionFee(birdAuction.EthPrice)

	// 构建应答结构体
	buyBird := data.BuyBird{
		Id:    tokenId,
		Price: birdAuction.EthPrice,
		Fee:   fee,
		Sign:  util.Sign2(tokenId, birdAuction.EthPrice, fee),
	}

	return &buyBird, util.ERROR_CODE_SUCCESS
}

func BuyBirdHash(address string, tokenId int64, hash string) (int) {
	//challenger bird存在
	bird := dao.QueryBirdByTokenId(tokenId)
	if bird == nil {
		return util.ERROR_CODE_NO_BIRD
	}

	//本人操作
	if strings.Compare(bird.Owner, address) == 0 {
		return util.ERROR_CODE_BUY_SELF_BIRD
	}

	// bird必须为销售状态
	if bird.Status != data.Bird_Status_Selling {
		return util.ERROR_CODE_NOT_SELL_BIRD
	}

	// 记录txn hash
	err := dao.InsertBirdTxn(tokenId, address, hash, data.Txn_unconfirmed)
	if !err {
		return util.ERROR_CODE_DB_FAILED
	}

	//更新状态
	dao.UpdateBirdStatus(tokenId, data.Slot_status_Selled_Unconfirmed)

	return util.ERROR_CODE_SUCCESS
}

func PKBird(address string, challengerId int64, resisterId int64) (*data.BirdPK, int) {
	//challenger bird存在
	bird := dao.QueryBirdByTokenId(challengerId)
	if bird == nil {
		return nil, util.ERROR_CODE_NO_BIRD
	}

	//本人操作
	if strings.Compare(bird.Owner, address) != 0 {
		return nil, util.ERROR_CODE_OP_OTHER_BIRD
	}

	//challenger bird存在
	bird1 := dao.QueryBirdByTokenId(resisterId)
	if bird1 == nil {
		return nil, util.ERROR_CODE_NO_BIRD
	}

	//不是同一个人的birds
	if strings.Compare(bird1.Owner, address) == 0 {
		return nil, util.ERROR_CODE_PK_SELF_BIRD
	}

	if (bird.Status != data.Bird_Status_Normal) || (bird1.Status != data.Bird_Status_Normal) {
		return nil, util.ERROR_CODE_BIRD_STATUS
	}

	// 构建响应结构体
	pkBird := data.BirdPK{
		ChallengerId: challengerId,
		ResisterId:   resisterId,
		Sign:         util.Sign2Int64(challengerId, resisterId),
	}

	return &pkBird, util.ERROR_CODE_SUCCESS
}

func PKBirdHash(address string, challengerId int64, resisterId int64, hash string) int {
	//challenger bird存在
	bird := dao.QueryBirdByTokenId(challengerId)
	if bird == nil {
		return util.ERROR_CODE_NO_BIRD
	}

	//本人操作
	if strings.Compare(bird.Owner, address) != 0 {
		return util.ERROR_CODE_OP_OTHER_BIRD
	}

	//challenger bird存在
	bird1 := dao.QueryBirdByTokenId(resisterId)
	if bird1 == nil {
		return util.ERROR_CODE_NO_BIRD
	}

	//不是同一个人的birds
	if strings.Compare(bird1.Owner, address) == 0 {
		return util.ERROR_CODE_PK_SELF_BIRD
	}

	if (bird.Status != data.Bird_Status_Normal) || (bird1.Status != data.Bird_Status_Normal) {
		return util.ERROR_CODE_BIRD_STATUS
	}

	// 记录txn hash
	err := dao.InsertBirdPKTxn(challengerId, resisterId, hash, data.Txn_unconfirmed)
	if !err {
		return util.ERROR_CODE_DB_FAILED
	}

	//更新状态
	dao.UpdateBirdStatus(challengerId, data.Bird_Status_PK_Unconfirmed)
	dao.UpdateBirdStatus(resisterId, data.Bird_Status_PK_Unconfirmed)

	return util.ERROR_CODE_SUCCESS
}

func RankBird(startPage int, pageSize int) (*map[string]interface{}, int) {
	if startPage <= 0 || (pageSize <= 0) {
		util.Logger.Error("query rank bird count failed.")
		return nil, util.ERROR_CODE_PARAM_INVALID
	}

	//获取birds数量
	count, err := dao.QueryBirdsCount()
	if !err {
		util.Logger.Error("query bird count failed.")
		return nil, util.ERROR_CODE_BIRDS_COUNT
	}

	var rankBirds = make(map[string]interface{})
	rankBirds["total_number"] = count
	rankBirds["start_page"] = startPage
	rankBirds["page_size"] = pageSize

	if (count == 0) || ((startPage-1)*pageSize >= count) {
		rankBirds["birds"] = []data.Bird{}
		return &rankBirds, util.ERROR_CODE_SUCCESS
	}

	size := pageSize
	if pageSize > 50 {
		size := 50
		rankBirds["page_size"] = size
	}

	// 已经bird体重排序
	birds, err := dao.QueryBirdsByRank((startPage-1)*pageSize, size)
	if birds == nil {
		rankBirds["birds"] = []data.Bird{}
		return &rankBirds, util.ERROR_CODE_SUCCESS
	}

	rankBirds["birds"] = birds
	return &rankBirds, util.ERROR_CODE_SUCCESS
}

func GetMarketBirds(startPage int, pageSize int, orderType int, sortType int) (*map[string]interface{}, int) {
	if startPage <= 0 || (pageSize <= 0) {
		return nil, util.ERROR_CODE_PARAM_INVALID
	}

	//获取market bird数量
	count, err := dao.QueryMarketBirdsCount()
	if !err {
		return nil, util.ERROR_CODE_MARKET_BIRDS
	}

	var marketBirds = make(map[string]interface{})
	marketBirds["total_number"] = count
	marketBirds["start_page"] = startPage
	marketBirds["page_size"] = pageSize

	if (count == 0) || ((startPage-1)*pageSize >= count) {
		marketBirds["birds"] = []data.Bird{}
		return &marketBirds, util.ERROR_CODE_SUCCESS
	}

	size := pageSize
	if pageSize > 50 {
		size := 50
		marketBirds["page_size"] = size
	}

	var keys []int64
	var auctionMap = make(map[int64]data.BirdAuction)
	// 获取市场的bird拍卖信息列表
	birdAuction, err := dao.QueryAllMarketBirds(orderType, sortType, (startPage-1)*pageSize, size)
	if birdAuction == nil {
		marketBirds["birds"] = []data.Bird{}
		return &marketBirds, util.ERROR_CODE_SUCCESS
	}
	// keys记录bird id，用于下一步查询bird列表
	for _, auction := range *birdAuction {
		auctionMap[auction.TokenId] = auction
		keys = append(keys, auction.TokenId)
	}

	//根据tokenIds查询所有bird
	birds, err := dao.QueryBirdByTokenIds(keys)
	if !err {
		return nil, util.ERROR_CODE_MARKET_BIRDS
	}

	//填充bird结构体
	for key, bird := range *birds {
		birdAuction := auctionMap[bird.TokenId]
		(*birds)[key].Auction = &birdAuction
	}
	marketBirds["birds"] = birds
	return &marketBirds, util.ERROR_CODE_SUCCESS
}

/**
 * 生成bird信息：图片、name等
 */
func GenerateBird(tokenId int64) (*data.Bird, int) {
	//查询数据库
	bird := dao.QueryBirdByTokenId(tokenId)
	if bird == nil {
		return nil, util.ERROR_CODE_NO_BIRD
	}

	// bird的状态是否为初始化状态
	if bird.Status == data.Bird_Status_Init {
		//生成图片等信息
		err := generateBirdPic(bird)
		if err {
			//更新db
			e := dao.UpdateBirdGenerateInfo(data.Bird_Status_Normal, bird)
			if e {
				return bird, util.ERROR_CODE_SUCCESS
			} else {
				return nil, util.ERROR_CODE_DB_FAILED
			}
		}
		return nil, util.ERROR_CODE_BIRD_CREATE
	}

	return nil, util.ERROR_CODE_BIRD_STATUS
}

func generateBirdPic(bird *data.Bird) (bool) {
	// 生成svg图片
	svgPath := BirdGeneratorInstance().GenerateSvg(bird.TokenId, bird.Genes)
	bird.SvgPath = svgPath
	// 生成bird name
	bird.Name = giveName(bird.TokenId)

	// todo 计算优先级，同学们自己完成
	// calcRarity(...)

	return true
}

func giveName(tokenId int64) (string) {
	return "bird.no" + strconv.FormatInt(tokenId, 10)
}

func calcRarity(rareCount int) string {
	switch rareCount {
	case 0, 1:
		return "普通"
	case 2:
		return "稀有"
	case 3, 4:
		return "卓越"
	case 5, 6:
		return "神话"
	case 7, 8:
		return "传说"
	}
	return ""
}

