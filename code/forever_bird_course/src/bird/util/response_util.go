package util

import (
	"encoding/json"
	"strconv"
	"bird/log"
)

const ERROR_CODE_SUCCESS = 0
const ERROR_CODE_NOTLOGIN = 1
const ERROR_CODE_PARAM_INVALID = 2
const ERROR_CODE_DB_FAILED = 3
const ERROR_CODE_OP_INVAILD = 4
const ERROR_USER_INVAILD = 9
const ERROR_USER_EXIST = 10
const ERROR_EMAIL_EXIST = 11
const ERROR_USER_REGISTER = 12
const ERROR_CODE_REGISTER_USER = 13
const ERROR_CODE_USER_QUERY = 14
const ERROR_CODE_USER_BIRD_QUERY = 15
const ERROR_CODE_NO_BIRD = 20
const ERROR_CODE_SELL_BIRD_STATUS = 21
const ERROR_CODE_SELLED_BIRD = 22
const ERROR_CODE_MARKET_BIRDS = 23
const ERROR_CODE_BIRDS_COUNT = 24
const ERROR_CODE_NOT_SELL_BIRD = 25
const ERROR_CODE_NOT_SELL_BIRD_INFO = 26
const ERROR_CODE_BIRD_STATUS = 27
const ERROR_CODE_BIRD_CREATE = 28

const ERROR_CODE_NO_FRUIT = 30
const ERROR_CODE_NO_RECORD = 31
const ERROR_CODE_EAT_FRUIT_TIME = 32

const ERROR_CODE_BUY_SELF_BIRD = 33
const ERROR_CODE_PK_SELF_BIRD = 34
const ERROR_CODE_OP_OTHER_BIRD = 35

const ERROR_CODE_QUERY_CATCH_BIRDS_COUNT = 36
const ERROR_CODE_QUERY_USER_BIRDS_COUNT = 37

const ERROR_CODE_SYSTEM_CRASH = 1000

var ErrorMap = map[int]string{
	ERROR_CODE_SUCCESS:       "success",
	ERROR_CODE_NOTLOGIN:      "用户没有登录",
	ERROR_USER_INVAILD:       "用户不存在",
	ERROR_CODE_PARAM_INVALID: "参数错误",
	ERROR_CODE_OP_INVAILD:    "异常操作",
	ERROR_USER_EXIST:         "地址已经注册",
	ERROR_EMAIL_EXIST:        "邮箱已经注册",
	ERROR_USER_REGISTER:      "用户注册失败",

	ERROR_CODE_NO_BIRD:            "没有发现对应鸟",
	ERROR_CODE_SELL_BIRD_STATUS:   "当前bird不能销售",
	ERROR_CODE_SELLED_BIRD:        "the bird is selling",
	ERROR_CODE_MARKET_BIRDS:       "query market birds failed...",
	ERROR_CODE_BIRDS_COUNT:        "query birds failed...",
	ERROR_CODE_NOT_SELL_BIRD:      "the bird is not selling",
	ERROR_CODE_NOT_SELL_BIRD_INFO: "the bird has no sell info",
	ERROR_CODE_BIRD_STATUS:		   "the bird status is not valid",
	ERROR_CODE_BIRD_CREATE:		   "create bird failed...",
	ERROR_CODE_NO_RECORD:		   "no record",
	ERROR_CODE_EAT_FRUIT_TIME:	   "bird 5分钟内只能吃一次水果",
	ERROR_CODE_BUY_SELF_BIRD:	   "不能购买自己的bird",
	ERROR_CODE_PK_SELF_BIRD:	   "自己的bird不能pk",
	ERROR_CODE_OP_OTHER_BIRD:	   "不能操作别人的bird",
	ERROR_CODE_QUERY_CATCH_BIRDS_COUNT:		"查询用户catch数量失败",
	ERROR_CODE_QUERY_USER_BIRDS_COUNT:		"查询用户bird数量失败",
}

type Nil struct{}

func GenerateResponse(error int, data interface{}, msg string) []byte {
	if data == nil {
		data = Nil{}
	}

	response := make(map[string]interface{})
	response["status"] = error
	response["data"] = data
	response["message"] = msg

	dat, err := json.Marshal(response)
	if err != nil {
		return []byte("{status:" + strconv.Itoa(ERROR_CODE_PARAM_INVALID) + "}")
	}
	log.Info(string(dat))
	return dat
}

func GenerateResponseMap(error int, data *map[string]interface{}, msg string) []byte {
	response := make(map[string]interface{})
	response["status"] = error
	response["data"] = *data
	response["message"] = msg

	dat, err := json.Marshal(response)
	if err != nil {
		return []byte("{status:" + strconv.Itoa(ERROR_CODE_PARAM_INVALID) + "}")
	}
	log.Info(string(dat))
	return dat
}

/**
 ** 请求处理成功
 */
func Success(data interface{}) []byte {
	return GenerateResponse(ERROR_CODE_SUCCESS, data, ErrorMap[ERROR_CODE_SUCCESS])
}

func SuccessWithMap(data *map[string]interface{}) []byte {
	return GenerateResponseMap(ERROR_CODE_SUCCESS, data, ErrorMap[ERROR_CODE_SUCCESS])
}

/**
 ** 异常response
 */
func Failed(error int) []byte {
	msg, err := ErrorMap[error]
	if !err {
		msg = "请求失败"
	}
	return GenerateResponse(error, Nil{}, msg)
}
