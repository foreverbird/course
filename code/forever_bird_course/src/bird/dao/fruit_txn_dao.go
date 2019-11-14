package dao

func InsertFruitTxn(birdId int64, fruitType int, address string, hash string, status int) (bool) {
	sql := "insert into buy_fruit_txn(`bird_id`, `fruit_type`, `address`, `hash`, `status`) values(?,?,?,?,?)"
	args := []interface{}{birdId, fruitType, address, hash, status}

	id := InsertWithArgs(sql, args)
	return id >= 0
}
