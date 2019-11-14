package dao

import (
	"strconv"
	"bird/data"
	"bird/util"
)

const (
	Order_by_time     = 0
	Order_by_price    = 1
	Order_by_rarity   = 2
	Order_by_propType = 2
)

var OrderKey = map[int]string{
	Order_by_time:   "CREATED",
	Order_by_price:  "price",
	Order_by_rarity: "rarity",
}

var OrderPropKey = map[int]string{
	Order_by_time:     "created",
	Order_by_price:    "price",
	Order_by_propType: "type",
}

const (
	Sort_asc  = 0
	Sort_desc = 1
)

var SortKey = map[int]string{
	Sort_asc:  "asc",
	Sort_desc: "desc",
}

/**
 ** 查询市场bird
 */
func QueryMarketBirdByTokenId(tokenId int64) (*data.BirdAuction, bool) {
	sql := `select * from market_bird where bird_token_id = ` + strconv.FormatInt(tokenId, 10)
	rows, err := Query(sql)
	if err == false {
		return nil, false
	}

	size := len(*rows)
	if size == 0 {
		return nil, true
	}

	birdAuction, err := formatMarketBirdInfo(&(*rows)[0])
	if !err {
		util.Logger.Error("QueryMarketBirdByTokenId failed ...", sql)
		return nil, true
	}

	return birdAuction, true
}

/**
 ** 按照分页查询数据
 */
func QueryAllMarketBirds(order int, sort int, offset int, count int) (*[]data.BirdAuction, bool) {
	o, err := OrderKey[order]
	if !err {
		o = "created"
	}
	s, err := SortKey[sort]
	if !err {
		s = "asc"
	}

	sql := `select * from market_bird order by ` + o + ` ` + s + ` limit ` + strconv.Itoa(offset) + `,` + strconv.Itoa(count)
	util.Logger.Debug("QueryAllMarketBirds", sql)

	rows, err := Query(sql)
	if err == false {
		return nil, false
	}

	size := len(*rows)
	if size == 0 {
		return nil, true
	}

	var record []data.BirdAuction
	for _, row := range *rows {
		if row != nil {
			birdAuction, err := formatMarketBirdInfo(&row)
			if err {
				record = append(record, *birdAuction)
			}
		}
	}
	return &record, true
}

func QueryUserBirdsCount(address string) (int, bool) {
	sql := "select count(1) from bird where `address` = \"" + address + "\""
	rows, err := Query(sql)
	if err == false {
		return 0, false
	}

	r := ((*rows)[0]).(map[string]string)
	count, _ := r["count(1)"]
	c, _ := strconv.Atoi(count)

	return c, true
}

func QueryMarketBirdsCount() (int, bool) {
	sql := `select count(1) from market_bird`
	rows, err := Query(sql)
	if err == false {
		return 0, false
	}

	r := ((*rows)[0]).(map[string]string)
	count, _ := r["count(1)"]
	c, _ := strconv.Atoi(count)

	return c, true
}

/**
 ** 查询birds
 */
func QueryMarketBirdByTokenIds(tokenIds []int64) (*[]data.BirdAuction, bool) {
	sql := `select * from market_bird where bird_token_id in(` + formatSql(&tokenIds) + `)`
	util.Logger.Info("QueryMarketBirdByTokenIds sql: ", sql)

	rows, err := Query(sql)
	if err == false {
		return nil, false
	}

	size := len(*rows)
	if size == 0 {
		return nil, true
	}

	var record []data.BirdAuction
	for _, row := range *rows {
		if row != nil {
			birdAuction, err := formatMarketBirdInfo(&row)
			if err {
				record = append(record, *birdAuction)
			}
		}
	}
	return &record, true
}

/**
 ** 查询市场bird
 */
func QueryMarketBirdByAddress(address string) (*[]data.BirdAuction, bool) {
	sql := `select * from market_bird where seller = "` + address + `"`
	rows, err := Query(sql)
	if err == false {
		return nil, false
	}

	size := len(*rows)
	if size == 0 {
		return nil, true
	}

	var record []data.BirdAuction
	for _, row := range *rows {
		if row != nil {
			birdAuction, err := formatMarketBirdInfo(&row)
			if err {
				record = append(record, *birdAuction)
			}
		}
	}
	return &record, true
}

func InsertMarketBird(tokenId int64, address string, rarity string, price string, auctionBegin int, auctionEnd int) (id int64) {
	sql := "insert into market_bird(`bird_token_id`, `seller`,`rarity`, `price`, `auction_begin`, `auction_end`) values(?,?,?,?,?,?)"
	args := []interface{}{tokenId, address, rarity, price, auctionBegin, auctionEnd,}

	id = InsertWithArgs(sql, args)
	return id
}

func DeleteMarketBirdByTokenId(tokenId int64) (bool) {
	sql := `delete from market_bird where bird_token_id = ` + strconv.FormatInt(tokenId, 10)
	err := Delete(sql)
	if err == 0 {
		return false
	}
	return true
}

func formatMarketBirdInfo(result *interface{}) (*data.BirdAuction, bool) {
	r := (*result).(map[string]string)
	id, e0 := r["bird_token_id"]
	seller, e1 := r["seller"]
	price, e2 := r["price"]
	begin, e3 := r["auction_begin"]
	end, e4 := r["auction_end"]

	if e0 && e1 && e2 && e3 && e4 {
		tokenId, _ := strconv.ParseInt(id, 10, 64)
		auctionBegin, _ := strconv.ParseInt(begin, 10, 64)
		auctionEnd, _ := strconv.ParseInt(end, 10, 64)

		birdAuction := data.BirdAuction{
			TokenId:      tokenId,
			Seller:       seller,
			EthPrice:     price,
			Price:        util.Wei2Eth(price),
			AuctionBegin: auctionBegin,
			AuctionEnd:   auctionEnd,
		}
		return &birdAuction, true
	}

	return nil, false
}
