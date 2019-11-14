package logic

import (
	"bird/dao"
	"bird/data"
	"bird/util"
	"strings"
)

/**
 * 判断是否为有效的用户请求
 */
func IsUserValid(address string, hash string) bool {
	// 判断签名
	recoveredAddr, err := util.RecoverTypedSignature("forever bird", hash)
	if !err {
		return false
	}
	return strings.Compare(recoveredAddr, address) == 0
}

func RegisterUser(address string, email string, nick string) (int) {
	//查询当前用户是否存在
	user, err := dao.QueryUserByAddress(address)
	if !err {
		return util.ERROR_CODE_DB_FAILED
	}
	if user != nil {
		//用户存在
		return util.ERROR_USER_EXIST
	}

	//查询当前用户是否存在
	user, err = dao.QueryUserByEmail(email)
	if !err {
		return util.ERROR_CODE_DB_FAILED
	}
	if user != nil {
		//用户存在
		return util.ERROR_USER_EXIST
	}

	// 插入db
	id := dao.InsertUser(address, email, nick)
	if id == 0 {
		return util.ERROR_USER_REGISTER
	}
	return util.ERROR_CODE_SUCCESS
}

func QueryUserInfo(address string) (*data.User, int) {
	user, err := dao.QueryUserByAddress(address)
	if err {
		return user, util.ERROR_CODE_SUCCESS
	}
	return nil, util.ERROR_CODE_USER_QUERY
}