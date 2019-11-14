package data

type CatchBirdRecord struct {
	TokenId  int64  `json:"token_id"`
	BirdName string `json:bird_name`
	Address  string `json:"address"`
	Hash     string `json:"hash"`
	Price    string `json:"price"`
	BirthDay string `json:"day"`
}

type BuyBirdRecord struct {
	TokenId int64  `json:"token_id"`
	BirdName string `json:bird_name`
	Buyer   string `json:"buyer"`
	Seller  string `json:"seller"`
	Price   string `json:"price"`
	Fee     string `json:"fee"`
	Hash    string `json:"hash"`
	Day     string `json:"day"`
}

type BuyFruitRecord struct {
	TokenId int64  `json:"token_id"`
	BirdName string `json:bird_name`
	FruitId int    `json:"-"`
	Address string `json:"address"`
	Price   string `json:"price"`
	Hash    string `json:"hash"`
	Day     string `json:"day"`
}

type PKRecord struct {
	ChallengerId        int64  `json:"challenger_id"`
	ChallengerName      string `json:"challenger_name"`
	ResisterId          int64  `json:"resister_id"`
	ResisterName        string `json:"resister_name"`
	IsWin               bool   `json:"is_win"`
	ChallengerRewardExp int    `json:"challenger_reward_exp"`
	ResisterRewardExp   int    `json:"resister_reward_exp"`
	WinnerRewardCoin    int64  `json:"winner_reward_coin"`
	Hash                string `json:"hash"`
	Day                 string `json:"day"`
}

type PKInfoRecord struct {
	ChallengerBird *Bird     `json:"challenger_bird"`
	ResisterBird   *Bird     `json:"resister_bird"`
	Record         *PKRecord `json:"pk_record"`
}

type ProfitStatistic struct {
	WeekIndex   int    `json:"week_index"`
	BeginTime   int    `json:"begin_time"`
	EndTime     int    `json:"end_time"`
	TotalProfit string `json:"total_profit"`
}

type ProfitInfo struct {
	WeekIndex int    `json:"week_index"`
	BirdId    int64  `json:"bird_id"`
	Address   string `json:"owner"`
	Rank      int    `json:"rank"`
	Weight    int    `json:"weight"`
	Profit    string `json:"profit"`
}
