package dao

import (
	"strconv"
	"bird/data"
)

func InsertCatchRecord(address string, hash string) (id int64) {
	sql := `insert into catch(address, hash, status) values(?,?,? )`
	args := []interface{}{address, hash, data.Txn_unconfirmed}

	id = InsertWithArgs(sql, args)
	return id
}

/**
 ** 更新
 */
func UpdateCatchData(address string, hash string, status int, transReceipt string) (bool) {
	sql := `update catch set status=?, trans_receipt=? where address=? and hash=?`
	args := []interface{}{status, transReceipt, address, hash}

	count := UpdateWithArgs(sql, args)

	return count == 1
}

func QueryUnConfirmAndOnChainCount(address string) (int, bool) {
	sql := "select count(1) from catch where `status` <= 1 and `address` = \"" + address + "\""
	rows, err := Query(sql)
	if err == false {
		return 0, false
	}

	r := ((*rows)[0]).(map[string]string)
	count, _ := r["count(1)"]
	c, _ := strconv.Atoi(count)

	return c, true
}

func QueryUnConfirmAndOnChain(address string) (*[]data.CatchTxn, bool) {
	sql := "select * from catch where `status` <= 1 and `address` = \"" + address + "\" order by `id` desc"
	rows, err := Query(sql)
	if err == false {
		return nil, false
	}

	size := len(*rows)
	if size == 0 {
		return nil, true
	}

	var record []data.CatchTxn
	for _, row := range *rows {
		if row != nil {
			catchTxn, err := formatCatchTxnInfo(&row)
			if err {
				record = append(record, *catchTxn)
			}
		}
	}
	return &record, true
}

func QueryUnConfirmAndOnChainWithLimit(address string, offset int, count int) (*[]data.CatchTxn, bool) {
	sql := "select * from catch where `status` <= 1 and `address` = \"" + address + "\" order by `id` desc limit " + strconv.Itoa(offset) + `,` + strconv.Itoa(count)
	rows, err := Query(sql)
	if err == false {
		return nil, false
	}

	size := len(*rows)
	if size == 0 {
		return nil, true
	}

	var record []data.CatchTxn
	for _, row := range *rows {
		if row != nil {
			catchTxn, err := formatCatchTxnInfo(&row)
			if err {
				record = append(record, *catchTxn)
			}
		}
	}
	return &record, true
}

/**
 ** 获取捕获未处理birds对应hash列表
 */
func QueryCatchTxnByStatus(status int) (*[]data.CatchTxn, bool) {
	sql := `select * from catch where status = ` + strconv.Itoa(status)
	rows, err := Query(sql)
	if err == false {
		return nil, false
	}

	size := len(*rows)
	if size == 0 {
		return nil, true
	}

	var record []data.CatchTxn
	for _, row := range *rows {
		if row != nil {
			catchTxn, err := formatCatchTxnInfo(&row)
			if err {
				record = append(record, *catchTxn)
			}
		}
	}
	return &record, true
}

func QueryUnconfirmedCatchTxn() (*[]data.CatchTxn, bool) {
	return QueryCatchTxnByStatus(data.Txn_unconfirmed)
}

func QueryMiningCatchTxn() (*[]data.CatchTxn, bool) {
	return QueryCatchTxnByStatus(data.Txn_mining)
}

/**
 ** 获取用户捕获birds对应hash列表
 */
func QueryCatchTxnByAddress(address string, status int) (*[]data.CatchTxn, bool) {
	sql := `select * from catch where status = ` + strconv.Itoa(status) + ` and address = "` + address + `"`
	rows, err := Query(sql)
	if err == false {
		return nil, false
	}

	size := len(*rows)
	if size == 0 {
		return nil, true
	}

	var record []data.CatchTxn
	for _, row := range *rows {
		if row != nil {
			catchTxn, err := formatCatchTxnInfo(&row)
			if err {
				record = append(record, *catchTxn)
			}
		}
	}
	return &record, true
}

func QueryUnconfirmedCatchTxnByAddress(address string) (*[]data.CatchTxn, bool) {
	return QueryCatchTxnByAddress(address, data.Txn_unconfirmed)
}

func QueryMiningCatchTxnByAddress(address string) (*[]data.CatchTxn, bool) {
	return QueryCatchTxnByAddress(address, data.Txn_mining)
}

func formatCatchTxnInfo(result *interface{}) (*data.CatchTxn, bool){
	r := (*result).(map[string]string)
	hash, e0 := r["hash"]
	address, e1 := r["address"]
	status, e2 := r["status"]
	if e0 && e1 && e2 {
		s, _ := strconv.Atoi(status)
		return &data.CatchTxn{
			Address:address,
			Hash:hash,
			Status:s,
		}, true
	}

	return nil, false
}