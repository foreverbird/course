package dao

func InsertBirdPKTxn(challengerId int64, resisterId int64, hash string, status int) (bool) {
	sql := "insert into bird_pk_txn(`challengerId`, `resisterId`, `hash`, `status`) values(?,?,?,?)"
	args := []interface{}{challengerId, resisterId, hash, status}

	id := InsertWithArgs(sql, args)
	return id >= 0
}
