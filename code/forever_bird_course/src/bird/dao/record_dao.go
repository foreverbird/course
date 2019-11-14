package dao

import (
	"strconv"
	"bird/data"
	"bird/util"
)

func QueryCatchRecordCountByAddress(address string) (int, bool) {
	sql := "select count(1) from catch_record where `address` = " + address
	rows, err := Query(sql)
	if err == false {
		return 0, false
	}

	r := ((*rows)[0]).(map[string]string)
	count, _ := r["count(1)"]
	c, _ := strconv.Atoi(count)

	return c, true
}

func QueryCatchBirdRecordByTokenId(tokenId int64, offset int, count int) (*data.CatchBirdRecord, bool) {
	sql := `select * from catch_record where token_id = ` + strconv.FormatInt(tokenId, 10) + " order by created desc limit " + strconv.Itoa(offset) + `,` + strconv.Itoa(count)
	rows, err := Query(sql)
	if err == false {
		return nil, false
	}

	size := len(*rows)
	if size == 0 {
		return nil, true
	}

	record := formatCatchBirdRecordInfo(&(*rows)[0])
	return record, true
}

func QueryCatchBirdRecordByAddress(address string, offset int, count int) (*[]data.CatchBirdRecord, bool) {
	sql := `select * from catch_record where address = "` + address + `"` + " order by created desc limit " + strconv.Itoa(offset) + `,` + strconv.Itoa(count)
	rows, err := Query(sql)
	if err == false {
		return nil, false
	}

	size := len(*rows)
	if size == 0 {
		return nil, true
	}

	var record []data.CatchBirdRecord
	for _, row := range *rows {
		if row != nil {
			bird := formatCatchBirdRecordInfo(&row)
			record = append(record, *bird)
		}
	}
	return &record, true
}

func QueryCatchRecordCountByTime(beginTime int) (int, bool) {
	sql := "select count(1) from catch_record where `real_time` > " + strconv.Itoa(beginTime)
	rows, err := Query(sql)
	if err == false {
		return 0, false
	}

	r := ((*rows)[0]).(map[string]string)
	count, _ := r["count(1)"]
	c, _ := strconv.Atoi(count)

	return c, true
}

func QueryCatchBirdTxPriceByTime(beginTime int) (*[]string, bool) {
	sql := "select `price` from catch_record where `real_time` > " + strconv.Itoa(beginTime)
	rows, err := Query(sql)
	if err == false {
		return nil, false
	}

	size := len(*rows)
	if size == 0 {
		return nil, true
	}

	var record []string
	for _, row := range *rows {
		if row != nil {
			r := (row).(map[string]string)
			price, _ := r["price"]
			record = append(record, price)
		}
	}

	return &record, true
}

func QueryBuyBirdTxPriceByTime(beginTime int) (*[]string, bool) {
	sql := "select `price` from buy_bird_record where `real_time` > " + strconv.Itoa(beginTime)
	rows, err := Query(sql)
	if err == false {
		return nil, false
	}

	size := len(*rows)
	if size == 0 {
		return nil, true
	}

	var record []string
	for _, row := range *rows {
		if row != nil {
			r := (row).(map[string]string)
			price, _ := r["price"]
			record = append(record, price)
		}
	}

	return &record, true
}

func QueryBuyFruitTxPriceByTime(beginTime int) (*[]string, bool) {
	sql := "select `price` from buy_fruit_record where `real_time` > " + strconv.Itoa(beginTime)
	rows, err := Query(sql)
	if err == false {
		return nil, false
	}

	size := len(*rows)
	if size == 0 {
		return nil, true
	}

	var record []string
	for _, row := range *rows {
		if row != nil {
			r := (row).(map[string]string)
			price, _ := r["price"]
			record = append(record, price)
		}
	}

	return &record, true
}

func QueryCatchBirdRecordsByTime(beginTime int, offset int, count int) (*[]data.CatchBirdRecord, bool) {
	sql := "select * from catch_record where `real_time` > " + strconv.Itoa(beginTime) + " order by `real_time` desc limit " + strconv.Itoa(offset) + `,` + strconv.Itoa(count)
	rows, err := Query(sql)
	if err == false {
		return nil, false
	}

	size := len(*rows)
	if size == 0 {
		return nil, true
	}

	var keys []int64
	var record []data.CatchBirdRecord
	for _, row := range *rows {
		if row != nil {
			bird := formatCatchBirdRecordInfo(&row)
			keys = append(keys, bird.TokenId)
			record = append(record, *bird)
		}
	}

	//遍历bird名称
	birds, err := QueryBirdsMapByTokenIds(keys)
	for index, r := range record {
		record[index].BirdName = getBirdName(birds, r.TokenId)
	}
	return &record, true
}


func QueryBuyBirdRecordCountByTime(beginTime int) (int, bool) {
	sql := "select count(1) from buy_bird_record where `real_time` > " + strconv.Itoa(beginTime)
	rows, err := Query(sql)
	if err == false {
		return 0, false
	}

	r := ((*rows)[0]).(map[string]string)
	count, _ := r["count(1)"]
	c, _ := strconv.Atoi(count)

	return c, true
}

func QueryBuyBirdRecordByTime(beginTime int, offset int, count int) (*[]data.BuyBirdRecord, bool) {
	sql := "select * from buy_bird_record where `real_time` > " + strconv.Itoa(beginTime) + " order by `real_time` desc limit " + strconv.Itoa(offset) + `,` + strconv.Itoa(count)
	rows, err := Query(sql)
	if err == false {
		return nil, false
	}

	size := len(*rows)
	if size == 0 {
		return nil, true
	}

	var keys []int64
	var record []data.BuyBirdRecord
	for _, row := range *rows {
		if row != nil {
			bird := formatBuyBirdRecordInfo(&row)
			keys = append(keys, bird.TokenId)
			record = append(record, *bird)
		}
	}

	//遍历bird名称
	birds, err := QueryBirdsMapByTokenIds(keys)
	for index, r := range record {
		record[index].BirdName = getBirdName(birds, r.TokenId)
	}
	return &record, true
}

func QueryBuyBirdRecordCountByTokenId(tokenId int64) (int, bool) {
	sql := "select count(1) from buy_bird_record where `bird_id` = " + strconv.FormatInt(tokenId, 10)
	rows, err := Query(sql)
	if err == false {
		return 0, false
	}

	r := ((*rows)[0]).(map[string]string)
	count, _ := r["count(1)"]
	c, _ := strconv.Atoi(count)

	return c, true
}

func QueryBuyBirdRecordByTokenId(tokenId int64, offset int, count int) (*[]data.BuyBirdRecord, bool) {
	sql := "select * from buy_bird_record where `bird_id` = " + strconv.FormatInt(tokenId, 10) + " order by created desc limit " + strconv.Itoa(offset) + `,` + strconv.Itoa(count)
	rows, err := Query(sql)
	if err == false {
		return nil, false
	}

	size := len(*rows)
	if size == 0 {
		return nil, true
	}

	var record []data.BuyBirdRecord
	for _, row := range *rows {
		if row != nil {
			bird := formatBuyBirdRecordInfo(&row)
			record = append(record, *bird)
		}
	}

	return &record, true
}

func QueryBuyFruitRecordCountByTime(beginTime int) (int, bool) {
	sql := "select count(1) from buy_fruit_record where `real_time` > " + strconv.Itoa(beginTime)
	rows, err := Query(sql)
	if err == false {
		return 0, false
	}

	r := ((*rows)[0]).(map[string]string)
	count, _ := r["count(1)"]
	c, _ := strconv.Atoi(count)

	return c, true
}

func QueryBuyFruitRecordByTime(beginTime int, offset int, count int) (*[]data.BuyFruitRecord, bool) {
	sql := "select * from buy_fruit_record where `real_time` > " + strconv.Itoa(beginTime) + " order by `real_time` desc limit " + strconv.Itoa(offset) + `,` + strconv.Itoa(count)
	rows, err := Query(sql)
	if err == false {
		return nil, false
	}

	size := len(*rows)
	if size == 0 {
		return nil, true
	}

	var keys []int64
	var record []data.BuyFruitRecord
	for _, row := range *rows {
		if row != nil {
			bird := formatBuyFruitRecordInfo(&row)
			keys = append(keys, bird.TokenId)
			record = append(record, *bird)
		}
	}

	//遍历bird名称
	birds, err := QueryBirdsMapByTokenIds(keys)
	for index, r := range record {
		record[index].BirdName = getBirdName(birds, r.TokenId)
	}
	return &record, true
}

func QueryBuyFruitRecordCountByTokenId(tokenId int64) (int, bool) {
	sql := "select count(1) from buy_fruit_record where `bird_id` = " + strconv.FormatInt(tokenId, 10)
	rows, err := Query(sql)
	if err == false {
		return 0, false
	}

	r := ((*rows)[0]).(map[string]string)
	count, _ := r["count(1)"]
	c, _ := strconv.Atoi(count)

	return c, true
}

func QueryBuyFruitRecordByTokenId(tokenId int64, offset int, count int) (*[]data.BuyFruitRecord, bool) {
	sql := "select * from buy_fruit_record where `bird_id` = " + strconv.FormatInt(tokenId, 10) + " order by created desc limit " + strconv.Itoa(offset) + `,` + strconv.Itoa(count)
	rows, err := Query(sql)
	if err == false {
		return nil, false
	}

	size := len(*rows)
	if size == 0 {
		return nil, true
	}

	var record []data.BuyFruitRecord
	for _, row := range *rows {
		if row != nil {
			bird := formatBuyFruitRecordInfo(&row)
			record = append(record, *bird)
		}
	}
	return &record, true
}

func QueryPKRecordCountByTokenId(tokenId int64) (int, bool) {
	id := strconv.FormatInt(tokenId, 10)
	sql := "select count(1) from bird_pk_record where `challengerId` = " + id + " or `resisterId` = " + id
	rows, err := Query(sql)
	if err == false {
		return 0, false
	}

	r := ((*rows)[0]).(map[string]string)
	count, _ := r["count(1)"]
	c, _ := strconv.Atoi(count)

	return c, true
}

func QueryPKRecordByTokenId(tokenId int64, offset int, count int) (*[]data.PKRecord, bool) {
	id := strconv.FormatInt(tokenId, 10)
	sql := "select * from bird_pk_record where `challengerId` = " + id + " or `resisterId` = " + id + " order by created desc limit " + strconv.Itoa(offset) + `,` + strconv.Itoa(count)
	rows, err := Query(sql)
	if err == false {
		return nil, false
	}

	size := len(*rows)
	if size == 0 {
		return nil, true
	}

	var keys []int64
	var record []data.PKRecord
	for _, row := range *rows {
		if row != nil {
			bird := formatPKRecordInfo(&row)
			keys = append(keys, bird.ChallengerId)
			keys = append(keys, bird.ResisterId)
			record = append(record, *bird)
		}
	}

	//遍历bird名称
	birds, err := QueryBirdsMapByTokenIds(keys)
	for index, r := range record {
		record[index].ChallengerName = getBirdName(birds, r.ChallengerId)
		record[index].ResisterName = getBirdName(birds, r.ResisterId)
	}

	return &record, true
}

func getBirdName(birds *map[int64]*data.Bird, tokenId int64) string {
	bird, err := (*birds)[tokenId]
	if !err {
		return "bird.no" + strconv.FormatInt(tokenId, 10)
	}
	return bird.Name
}

func QueryPKRecordByHash(hash string) (*data.PKRecord, bool) {
	sql := "select * from bird_pk_record where `hash` = \"" + hash + "\""
	rows, err := Query(sql)
	if err == false {
		return nil, false
	}

	size := len(*rows)
	if size == 0 {
		return nil, true
	}

	bird := formatPKRecordInfo(&(*rows)[0])
	return bird, true
}

func QueryProfitStaticsMaxWeek() (int, bool) {
	sql := "select max(`week_index`) from profit_statistics_record"
	rows, err := Query(sql)
	if err == false {
		return 0, false
	}

	r := ((*rows)[0]).(map[string]string)
	count, _ := r["max(`week_index`)"]
	c, _ := strconv.Atoi(count)

	return c, true
}

func QueryProfitStatics(week int) (*data.ProfitStatistic, bool) {
	sql := "select * from profit_statistics_record where `week_index` = " + strconv.Itoa(week)
	rows, err := Query(sql)
	if err == false {
		return nil, false
	}

	size := len(*rows)
	if size == 0 {
		return nil, true
	}

	profitStatistic := formatProfitStatisRecordInfo(&(*rows)[0])
	return profitStatistic, true
}

func QueryProfitList(week int) (*[]data.ProfitInfo, bool) {
	sql := "select * from profit_detail_record where `week_index` = " + strconv.Itoa(week) + " order by `rank_no` asc"
	rows, err := Query(sql)
	if err == false {
		return nil, false
	}

	size := len(*rows)
	if size == 0 {
		return nil, true
	}

	var record []data.ProfitInfo
	for _, row := range *rows {
		if row != nil {
			pi := formatProfitInfo(&row)
			record = append(record, *pi)
		}
	}
	return &record, true
}

func formatCatchBirdRecordInfo(result *interface{}) (*data.CatchBirdRecord) {
	r := (*result).(map[string]string)
	id, _ := r["bird_id"]
	address, _ := r["address"]
	hash, _ := r["hash"]
	price, _ := r["price"]
	birth, _ := r["real_time"]

	tokenId, _ := strconv.ParseInt(id, 10, 64)

	record := data.CatchBirdRecord{
		TokenId:  tokenId,
		Address:  address,
		Hash:     hash,
		Price:    util.Wei2Eth(price),
		BirthDay: birth,
	}
	return &record
}

func formatBuyBirdRecordInfo(result *interface{}) (*data.BuyBirdRecord) {
	r := (*result).(map[string]string)
	id, _ := r["bird_id"]
	buyer, _ := r["buyer"]
	seller, _ := r["seller"]
	hash, _ := r["hash"]
	price, _ := r["price"]
	fee, _ := r["fee"]
	day, _ := r["real_time"]

	tokenId, _ := strconv.ParseInt(id, 10, 64)

	record := data.BuyBirdRecord{
		TokenId: tokenId,
		Buyer:   buyer,
		Seller:  seller,
		Price:   util.Wei2Eth(price),
		Fee:     util.Wei2Eth(fee),
		Hash:    hash,
		Day:     day,
	}
	return &record
}

func formatBuyFruitRecordInfo(result *interface{}) (*data.BuyFruitRecord) {
	r := (*result).(map[string]string)
	id, _ := r["bird_id"]
	fId, _ := r["fruit_id"]
	address, _ := r["address"]
	hash, _ := r["hash"]
	price, _ := r["price"]
	day, _ := r["real_time"]

	tokenId, _ := strconv.ParseInt(id, 10, 64)
	fruitId, _ := strconv.Atoi(fId)

	record := data.BuyFruitRecord{
		TokenId: tokenId,
		FruitId: fruitId,
		Address: address,
		Price:   util.Wei2Eth(price),
		Hash:    hash,
		Day:     day,
	}
	return &record
}

func formatPKRecordInfo(result *interface{}) (*data.PKRecord) {
	r := (*result).(map[string]string)
	challengerId, _ := r["challengerId"]
	resisterId, _ := r["resisterId"]
	isWin, _ := r["isWin"]
	challengerRewardExp, _ := r["challengerRewardExp"]
	resisterRewardExp, _ := r["resisterRewardExp"]
	winnerRewardCoin, _ := r["winnerRewardCoin"]
	hash, _ := r["hash"]
	day, _ := r["real_time"]

	cId, _ := strconv.ParseInt(challengerId, 10, 64)
	rId, _ := strconv.ParseInt(resisterId, 10, 64)
	cWin, _ := strconv.Atoi(isWin)
	cExp, _ := strconv.Atoi(challengerRewardExp)
	rExp, _ := strconv.Atoi(resisterRewardExp)
	wCoin, _ := strconv.ParseInt(winnerRewardCoin, 10, 64)

	win := true
	if cWin <= 0 {
		win = false
	}

	record := data.PKRecord{
		ChallengerId:        cId,
		ResisterId:          rId,
		IsWin:               win,
		ChallengerRewardExp: cExp,
		ResisterRewardExp:   rExp,
		WinnerRewardCoin:    wCoin,
		Hash:                hash,
		Day:                 day,
	}
	return &record
}

func formatProfitStatisRecordInfo(result *interface{}) (*data.ProfitStatistic) {
	r := (*result).(map[string]string)
	week, _ := r["week_index"]
	beginTime, _ := r["begin_time"]
	endTime, _ := r["end_time"]
	totalProfit, _ := r["total_profit"]

	wi, _ := strconv.Atoi(week)
	bt, _ := strconv.Atoi(beginTime)
	et, _ := strconv.Atoi(endTime)

	record := data.ProfitStatistic{
		WeekIndex:   wi,
		BeginTime:   bt,
		EndTime:     et,
		TotalProfit: util.Wei2Eth(totalProfit),
	}
	return &record
}

func formatProfitInfo(result *interface{}) (*data.ProfitInfo) {
	r := (*result).(map[string]string)
	bid, _ := r["bird_id"]
	address, _ := r["address"]
	rank, _ := r["rank_no"]
	weight, _ := r["weight_record"]
	week, _ := r["week_index"]
	profit, _ := r["profit"]

	birdId, _ := strconv.ParseInt(bid, 10, 64)
	rno, _ := strconv.Atoi(rank)
	wi, _ := strconv.Atoi(week)
	w, _ := strconv.Atoi(weight)

	record := data.ProfitInfo{
		WeekIndex: wi,
		BirdId:    birdId,
		Address:   address,
		Weight:    w,
		Rank:      rno,
		Profit:    util.Wei2Eth(profit),
	}
	return &record
}
