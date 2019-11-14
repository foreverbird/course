package data

const (
	Fruit_type_power = 1	//力量水果
	Fruit_type_speed = 2	//速度水果
	Fruit_type_exp   = 3	//经验水果
)

/**
 * 水果基本结构
 */
type Fruit struct {
	Id       int    `json:"id"`
	Type     int    `json:"type"`
	Desc     string `json:"desc"`
	Price    string `json:"price"`
	Strength int    `json:"strength"`
}

/**
 * 购买水果返回的结构体
 */
type BuyFruit struct {
	Id    int64  `json:"bird_id"`
	Type  int    `json:"type"`
	Price string `json:"price"`
	Sign  string `json:"sign"`
}
