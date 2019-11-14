package dao

func InsertBirdTxn(birdId int64, address string, hash string, status int) (bool) {
	sql := "insert into buy_bird_txn(`bird_id`, `address`, `hash`, `status`) values(?,?,?,?)"
	args := []interface{}{birdId, address, hash, status}

	id := InsertWithArgs(sql, args)
	return id >= 0
}
