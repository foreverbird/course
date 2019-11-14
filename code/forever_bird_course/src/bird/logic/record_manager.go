package logic

import (
	"bird/dao"
	"bird/data"
	"bird/util"
	"math/big"
)

func QueryQueryCatchRecordInWeek(startPage int, pageSize int) (*map[string]interface{}, int) {
	beginTime := util.CountCurWeekStartingTime()
	return QueryCatchRecordByTime(beginTime, startPage, pageSize)
}

func QueryCatchRecordByTime(beginTime int, startPage int, pageSize int) (*map[string]interface{}, int) {
	if startPage <= 0 || (pageSize <= 0) {
		util.Logger.Error("query catch bird record count failed.")
		return nil, util.ERROR_CODE_PARAM_INVALID
	}

	//获取catch bird数量
	count, err := dao.QueryCatchRecordCountByTime(beginTime)
	if !err {
		util.Logger.Error("query catch bird record count failed.")
		return nil, util.ERROR_CODE_MARKET_BIRDS
	}

	var catchInfos = make(map[string]interface{})
	catchInfos["total_number"] = count
	catchInfos["start_page"] = startPage
	catchInfos["page_size"] = pageSize

	// 判断分页范围是否有数据
	if (count == 0) || ((startPage-1)*pageSize >= count) {
		catchInfos["records"] = []data.CatchBirdRecord{}
		return &catchInfos, util.ERROR_CODE_SUCCESS
	}

	size := pageSize
	if pageSize > 50 {
		size := 50
		catchInfos["page_size"] = size
	}

	// 按照时间请求记录
	records, err := dao.QueryCatchBirdRecordsByTime(beginTime, (startPage-1)*pageSize, size)
	if !err {
		return nil, util.ERROR_CODE_DB_FAILED
	}

	if records == nil {
		catchInfos["records"] = []data.CatchBirdRecord{}
		return &catchInfos, util.ERROR_CODE_SUCCESS
	}

	catchInfos["records"] = records
	return &catchInfos, util.ERROR_CODE_SUCCESS
}

func QueryCatchRecordByAddress(address string, startPage int, pageSize int) (*map[string]interface{}, int) {
	if startPage <= 0 || (pageSize <= 0) {
		util.Logger.Error("query catch bird record count failed.")
		return nil, util.ERROR_CODE_PARAM_INVALID
	}

	//获取catch bird数量
	count, err := dao.QueryCatchRecordCountByAddress(address)
	if !err {
		util.Logger.Error("query catch bird record count failed.")
		return nil, util.ERROR_CODE_MARKET_BIRDS
	}

	var ownerBirds = make(map[string]interface{})
	ownerBirds["total_number"] = count
	ownerBirds["start_page"] = startPage
	ownerBirds["page_size"] = pageSize

	if (count == 0) || ((startPage-1)*pageSize >= count) {
		ownerBirds["records"] = []data.CatchBirdRecord{}
		return &ownerBirds, util.ERROR_CODE_SUCCESS
	}

	size := pageSize
	if pageSize > 50 {
		size := 50
		ownerBirds["page_size"] = size
	}

	records, err := dao.QueryCatchBirdRecordByAddress(address, (startPage-1)*pageSize, size)
	if !err {
		return nil, util.ERROR_CODE_DB_FAILED
	}

	if records == nil {
		ownerBirds["records"] = []data.CatchBirdRecord{}
		return &ownerBirds, util.ERROR_CODE_SUCCESS
	}

	ownerBirds["records"] = records
	return &ownerBirds, util.ERROR_CODE_SUCCESS
}

func QueryBuyBirdRecordInWeek(startPage int, pageSize int) (*map[string]interface{}, int) {
	beginTime := util.CountCurWeekStartingTime()
	return QueryBuyBirdRecordByTime(beginTime, startPage, pageSize)
}

func QueryBuyBirdRecordByTime(beginTime int, startPage int, pageSize int) (*map[string]interface{}, int) {
	if startPage <= 0 || (pageSize <= 0) {
		util.Logger.Error("query buy bird record count failed.")
		return nil, util.ERROR_CODE_PARAM_INVALID
	}

	//buy bird数量
	count, err := dao.QueryBuyBirdRecordCountByTime(beginTime)
	if !err {
		util.Logger.Error("query buy bird record count failed.")
		return nil, util.ERROR_CODE_MARKET_BIRDS
	}

	var ownerBirds = make(map[string]interface{})
	ownerBirds["total_number"] = count
	ownerBirds["start_page"] = startPage
	ownerBirds["page_size"] = pageSize

	if (count == 0) || ((startPage-1)*pageSize >= count) {
		ownerBirds["records"] = []data.BuyBirdRecord{}
		return &ownerBirds, util.ERROR_CODE_SUCCESS
	}

	size := pageSize
	if pageSize > 50 {
		size := 50
		ownerBirds["page_size"] = size
	}

	records, err := dao.QueryBuyBirdRecordByTime(beginTime, (startPage-1)*pageSize, size)
	if !err {
		return nil, util.ERROR_CODE_DB_FAILED
	}

	if records == nil {
		ownerBirds["records"] = []data.BuyBirdRecord{}
		return &ownerBirds, util.ERROR_CODE_SUCCESS
	}

	ownerBirds["records"] = records
	return &ownerBirds, util.ERROR_CODE_SUCCESS
}

func QueryBuyBirdRecordByTokenId(tokenId int64, startPage int, pageSize int) (*map[string]interface{}, int) {
	if startPage <= 0 || (pageSize <= 0) {
		util.Logger.Error("query buy bird record count failed.")
		return nil, util.ERROR_CODE_PARAM_INVALID
	}

	//buy bird数量
	count, err := dao.QueryBuyBirdRecordCountByTokenId(tokenId)
	if !err {
		util.Logger.Error("query buy bird record count failed.")
		return nil, util.ERROR_CODE_MARKET_BIRDS
	}

	var ownerBirds = make(map[string]interface{})
	ownerBirds["total_number"] = count
	ownerBirds["start_page"] = startPage
	ownerBirds["page_size"] = pageSize

	if (count == 0) || ((startPage-1)*pageSize >= count) {
		ownerBirds["records"] = []data.BuyBirdRecord{}
		return &ownerBirds, util.ERROR_CODE_SUCCESS
	}

	size := pageSize
	if pageSize > 50 {
		size := 50
		ownerBirds["page_size"] = size
	}

	records, err := dao.QueryBuyBirdRecordByTokenId(tokenId, (startPage-1)*pageSize, size)
	if !err {
		return nil, util.ERROR_CODE_DB_FAILED
	}

	if records == nil {
		ownerBirds["records"] = []data.BuyBirdRecord{}
		return &ownerBirds, util.ERROR_CODE_SUCCESS
	}

	ownerBirds["records"] = records
	return &ownerBirds, util.ERROR_CODE_SUCCESS
}

func QueryBuyFruitRecordInWeek(startPage int, pageSize int) (*map[string]interface{}, int) {
	beginTime := util.CountCurWeekStartingTime()
	return QueryBuyFruitRecordByTime(beginTime, startPage, pageSize)
}

func QueryTxStatistics() (int, string, int) {
	beginTime := util.CountCurWeekStartingTime()

	c1, v1, e1 := QueryBuyFruitTx(beginTime)
	if e1 != 0 {
		return 0, "0", e1
	}

	c2, v2, e2 := QueryBuyBirdTx(beginTime)
	if e2 != 0 {
		return 0, "0", e2
	}

	c3, v3, e3 := QueryCatchTx(beginTime)
	if e3 != 0 {
		return 0, "0", e3
	}

	c := c1 + c2 + c3
	value := v1.Add(v1, v2)
	value = value.Add(value, v3)

	return c, util.Wei2Eth(value.String()), util.ERROR_CODE_SUCCESS
}

func QueryCatchTx(beginTime int) (int, *big.Int, int) {
	count := 0
	value := big.NewInt(256)
	value.SetString("0", 10)

	records, err := dao.QueryCatchBirdTxPriceByTime(beginTime)
	if !err {
		return 0, value, util.ERROR_CODE_DB_FAILED
	}

	if records == nil {
		return 0, value, 0
	}

	count = len(*records)
	for _, row := range *records {
		v := big.NewInt(256)
		v.SetString(row, 10)
		value = value.Add(value, v)
	}

	return count, value, 0
}

func QueryBuyBirdTx(beginTime int) (int, *big.Int, int) {
	count := 0
	value := big.NewInt(256)
	value.SetString("0", 10)

	records, err := dao.QueryBuyBirdTxPriceByTime(beginTime)
	if !err {
		return 0, value, util.ERROR_CODE_DB_FAILED
	}

	if records == nil {
		return 0, value, 0
	}

	count = len(*records)
	for _, row := range *records {
		v := big.NewInt(256)
		v.SetString(row, 10)
		value = value.Add(value, v)
	}

	return count, value, 0
}

func QueryBuyFruitTx(beginTime int) (int, *big.Int, int) {
	count := 0
	value := big.NewInt(256)
	value.SetString("0", 10)

	records, err := dao.QueryBuyFruitTxPriceByTime(beginTime)
	if !err {
		return 0, value, util.ERROR_CODE_DB_FAILED
	}

	if records == nil {
		return 0, value, 0
	}
	count = len(*records)
	for _, row := range *records {
		v := big.NewInt(256)
		v.SetString(row, 10)
		value = value.Add(value, v)
	}

	return count, value, 0
}

func QueryBuyFruitRecordByTime(beginTime int, startPage int, pageSize int) (*map[string]interface{}, int) {
	if startPage <= 0 || (pageSize <= 0) {
		util.Logger.Error("query buy fruit record count failed.")
		return nil, util.ERROR_CODE_PARAM_INVALID
	}

	//获取catch bird数量
	count, err := dao.QueryBuyFruitRecordCountByTime(beginTime)
	if !err {
		util.Logger.Error("query buy fruit record count failed.")
		return nil, util.ERROR_CODE_MARKET_BIRDS
	}

	var ownerBirds = make(map[string]interface{})
	ownerBirds["total_number"] = count
	ownerBirds["start_page"] = startPage
	ownerBirds["page_size"] = pageSize

	if (count == 0) || ((startPage-1)*pageSize >= count) {
		ownerBirds["records"] = []data.BuyFruitRecord{}
		return &ownerBirds, util.ERROR_CODE_SUCCESS
	}

	size := pageSize
	if pageSize > 50 {
		size := 50
		ownerBirds["page_size"] = size
	}

	records, err := dao.QueryBuyFruitRecordByTime(beginTime, (startPage-1)*pageSize, size)
	if !err {
		return nil, util.ERROR_CODE_DB_FAILED
	}

	if records == nil {
		ownerBirds["records"] = []data.BuyFruitRecord{}
		return &ownerBirds, util.ERROR_CODE_SUCCESS
	}

	ownerBirds["records"] = records
	return &ownerBirds, util.ERROR_CODE_SUCCESS
}

func QueryBuyFruitRecordByTokenId(tokenId int64, startPage int, pageSize int) (*map[string]interface{}, int) {
	if startPage <= 0 || (pageSize <= 0) {
		util.Logger.Error("query buy fruit record count failed.")
		return nil, util.ERROR_CODE_PARAM_INVALID
	}

	//获取catch bird数量
	count, err := dao.QueryBuyFruitRecordCountByTokenId(tokenId)
	if !err {
		util.Logger.Error("query buy fruit record count failed.")
		return nil, util.ERROR_CODE_MARKET_BIRDS
	}

	var ownerBirds = make(map[string]interface{})
	ownerBirds["total_number"] = count
	ownerBirds["start_page"] = startPage
	ownerBirds["page_size"] = pageSize

	if (count == 0) || ((startPage-1)*pageSize >= count) {
		ownerBirds["records"] = []data.BuyFruitRecord{}
		return &ownerBirds, util.ERROR_CODE_SUCCESS
	}

	size := pageSize
	if pageSize > 50 {
		size := 50
		ownerBirds["page_size"] = size
	}

	records, err := dao.QueryBuyFruitRecordByTokenId(tokenId, (startPage-1)*pageSize, size)
	if !err {
		return nil, util.ERROR_CODE_DB_FAILED
	}

	if records == nil {
		ownerBirds["records"] = []data.BuyFruitRecord{}
		return &ownerBirds, util.ERROR_CODE_SUCCESS
	}

	ownerBirds["records"] = records
	return &ownerBirds, util.ERROR_CODE_SUCCESS
}

func QueryPKRecordByTokenId(tokenId int64, startPage int, pageSize int) (*map[string]interface{}, int) {
	if startPage <= 0 || (pageSize <= 0) {
		util.Logger.Error("query bird pk record count failed.")
		return nil, util.ERROR_CODE_PARAM_INVALID
	}

	//获取catch bird数量
	count, err := dao.QueryPKRecordCountByTokenId(tokenId)
	if !err {
		util.Logger.Error("query bird pk record count failed.")
		return nil, util.ERROR_CODE_MARKET_BIRDS
	}

	var ownerBirds = make(map[string]interface{})
	ownerBirds["total_number"] = count
	ownerBirds["start_page"] = startPage
	ownerBirds["page_size"] = pageSize

	if (count == 0) || ((startPage-1)*pageSize >= count) {
		ownerBirds["birds"] = []data.PKRecord{}
		return &ownerBirds, util.ERROR_CODE_SUCCESS
	}

	size := pageSize
	if pageSize > 50 {
		size := 50
		ownerBirds["page_size"] = size
	}

	birds, err := dao.QueryPKRecordByTokenId(tokenId, (startPage-1)*pageSize, size)
	if !err {
		return nil, util.ERROR_CODE_DB_FAILED
	}

	if birds == nil {
		ownerBirds["birds"] = []data.PKRecord{}
		return &ownerBirds, util.ERROR_CODE_SUCCESS
	}

	ownerBirds["birds"] = birds
	return &ownerBirds, util.ERROR_CODE_SUCCESS
}

func QueryPKInfoRecordByHash(hash string) (*data.PKInfoRecord, int) {
	record, err := dao.QueryPKRecordByHash(hash)
	if !err {
		return nil, util.ERROR_CODE_DB_FAILED
	}

	if record == nil {
		return nil, util.ERROR_CODE_NO_RECORD
	}

	cBird := dao.QueryBirdByTokenId(record.ChallengerId)
	if cBird == nil {
		return nil, util.ERROR_CODE_NO_BIRD
	}

	rBird := dao.QueryBirdByTokenId(record.ResisterId)
	if rBird == nil {
		return nil, util.ERROR_CODE_NO_BIRD
	}

	info := data.PKInfoRecord{
		ChallengerBird: cBird,
		ResisterBird:   rBird,
		Record:         record,
	}
	return &info, util.ERROR_CODE_SUCCESS
}

func QueryProfitStatisticsRecord(week int) (*data.ProfitStatistic, int) {
	lastWeek, err := dao.QueryProfitStaticsMaxWeek()
	if !err {
		return nil, util.ERROR_CODE_DB_FAILED
	}

	if (week <= 0) || (week > lastWeek) {
		week = lastWeek
	}

	ps, err := dao.QueryProfitStatics(week)
	if !err {
		return nil, util.ERROR_CODE_DB_FAILED
	}

	return ps, util.ERROR_CODE_SUCCESS
}

func QueryProfitListRecord(week int) (*[]data.ProfitInfo, int) {
	lastWeek, err := dao.QueryProfitStaticsMaxWeek()
	if !err {
		return nil, util.ERROR_CODE_DB_FAILED
	}

	if (week <= 0) || (week > lastWeek) {
		week = lastWeek
	}

	pl, err := dao.QueryProfitList(week)
	if !err {
		return nil, util.ERROR_CODE_DB_FAILED
	}

	return pl, util.ERROR_CODE_SUCCESS
}

func QueryProfitRecord(week int) (*map[string]interface{}, int) {
	lastWeek, err := dao.QueryProfitStaticsMaxWeek()
	if !err {
		return nil, util.ERROR_CODE_DB_FAILED
	}

	if (week <= 0) || (week > lastWeek) {
		week = lastWeek
	}

	ps, e := dao.QueryProfitStatics(week)
	if !e {
		return nil, util.ERROR_CODE_DB_FAILED
	}

	pl, err := dao.QueryProfitList(week)
	if !err {
		return nil, util.ERROR_CODE_DB_FAILED
	}

	var profit = make(map[string]interface{})
	profit["statistics"] = ps
	profit["list"] = pl

	return &profit, util.ERROR_CODE_SUCCESS
}
