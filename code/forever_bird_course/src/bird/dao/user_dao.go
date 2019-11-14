package dao

import (
	"bird/data"
	"strconv"
	"bird/util"
)

func IsUserExist(address string) (exist bool) {
	sql := "select * from user where `address` = " + address + "`"
	result, err := Query(sql)
	if (err == false) || (len(*result) == 0) {
		return false
	}
	return true
}

/**
 ** 按照分页查询数据
 */
func QueryUserByEmail(email string) (*data.User, bool) {
	sql := "select * from user where `email`=\"" + email + "\""
	util.Logger.Debug("QueryUserByAddress:", sql)

	rows, err := Query(sql)
	if err == false {
		return nil, false
	}

	size := len(*rows)
	if size == 0 {
		return nil, true
	}

	user, err := formatUserInfo(&(*rows)[0])
	if !err {
		return nil, false
	}
	return user, true
}

func InsertUser(address string, email string, nick string) (id int64){
	sql := "insert into user(`address`, `email`, `nick`) values(?,?,? )"
	args := []interface{}{address, email, nick}

	id = InsertWithArgs(sql, args)
	return id
}

func QueryUserByAddress(address string) (*data.User, bool) {
	sql := `select * from user where address = "` + address + `"`
	result, err := Query(sql)
	if err == false {
		return nil, false
	}

	if len(*result) == 0 {
		return nil, true
	}

	return formatUserInfo(&(*result)[0])
}

func formatUserInfo(result *interface{}) (*data.User, bool){
	r := (*result).(map[string]string)
	id, e0 := r["id"]
	address, e1 := r["address"]
	nick, e2 := r["nick"]
	email, e3 := r["email"]
	image, e4 := r["image"]

	if e0 && e1 && e2 && e3 && e4 {
		idInt, _ := strconv.ParseInt(id, 10, 64)
		user := data.User{Id:idInt, Address:address, Email:util.EmailDesens(email), Nick:nick, Image:image}
		return &user, true
	}

	return nil, false
}