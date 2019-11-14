package data

const (
	Bird_Status_Init               = 0 //初始状态
	Bird_Status_Mining             = 1 //正在上链
	Bird_Status_Normal             = 2 //正常状态
	Bird_Status_Selling            = 3 //正常销售
	Bird_Status_Pk                 = 4 //正在PK
	Bird_Status_PK_Unconfirmed     = 5 //订单未确定状态
	Slot_status_Selled_Unconfirmed = 6 //销售订单未确定状态
	Bird_Status_Selled_Mining      = 7 //销售订单确定，正在上链
	Bird_Status_Eating             = 8 //正在吃水果
	Bird_Status_Eat_Unconfirmed    = 9 //订单未确定状态
)

type EthBird struct {
	//token id
	TokenId int64
	//拥有者地址
	Address string
	//基因组 256位二进制字符串
	Genome string
	//体重，微分作为单位
	Weight int64
	//经验
	Experience int64
	//生日
	BirthDay uint
	//交易凭证
	Receipt string
}

/**
 ** bird基本信息
 */
type Bird struct {
	//pet id
	TokenId int64 `json:"token_id"`
	//pet name
	Name string `json:"name"`
	//bird describe
	Desc string `json:"desc"`
	//birthday
	Birthday uint `json:"birth_date"`
	//稀有程度
	Rarity string `json:"rarity"`
	//速度
	Speed int `json:"speed"`
	//力量
	Power int `json:"power"`
	//级别
	Level int `json:"level"`
	//重量
	Weight int64 `json:"weight"`
	//经验值
	Experience int64 `json:"exp"`
	//owner address
	Owner string `json:"owner"`
	//状态:未领养、上架、领养、销售中、上链中
	Status int `json:"status"`
	//基因
	Genes string `json:"-"`
	//svg路径
	SvgPath string `json:"svg_path"`
	//拍卖信息
	Auction      *BirdAuction `json:"auction_data"`
	EatFruitTime int64        `json:"last_eat_fruit"`
}

type RankBird struct {
	RowNo int `json:"row_no"`
	//pet id
	TokenId int64 `json:"token_id"`
	//pet name
	Name string `json:"name"`
	//bird describe
	Desc string `json:"desc"`
	//birthday
	Birthday uint `json:"birth_date"`
	//稀有程度
	Rarity string `json:"rarity"`
	//速度
	Speed int `json:"speed"`
	//力量
	Power int `json:"power"`
	//级别
	Level int `json:"level"`
	//重量
	Weight int64 `json:"weight"`
	//经验值
	Experience int64 `json:"exp"`
	//owner address
	Owner string `json:"owner"`
	//状态:未领养、上架、领养、销售中、上链中
	Status int `json:"status"`
	//基因组 256位二进制字符串
	Genes string `json:"-"`
	//svg路径
	SvgPath string `json:"svg_path"`
	//拍卖信息
	Auction *BirdAuction `json:"auction_data"`
}

/**
 ** bird 拍卖信息
 */
type BirdAuction struct {
	TokenId      int64  `json:"token_id"`
	Price        string `json:"price"`
	EthPrice     string `json:"-"`
	Seller       string `json:"seller"`
	AuctionBegin int64  `json:"auction_begin"`
	AuctionEnd   int64  `json:"auction_end"`
}

/**
 ** 交易信息
 */
type CatchTxn struct {
	Address string `json:"address"`
	Hash    string `json:"hash"`
	Status  int    `json:"status"`
}

/**
 ** bird pk信息
 */
type BirdPK struct {
	ChallengerId int64  `json:"challengerId"`
	ResisterId   int64  `json:"resisterId"`
	Sign         string `json:"sign"`
}

type BuyBird struct {
	Id    int64  `json:"bird_id"`
	Price string `json:"price"`
	Fee   string `json:"fee"`
	Sign  string `json:"sign"`
}
