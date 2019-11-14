package server

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"bird/logic"
	"bird/util"
)

/**
 * 解析用户请求参数
 */
func parseBody(r *http.Request) (*map[string]interface{}, int) {
	//提取参数
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, util.ERROR_CODE_PARAM_INVALID
	}
	var params map[string]interface{}
	err = json.Unmarshal(body, &params)
	if err != nil {
		return nil, util.ERROR_CODE_PARAM_INVALID
	}
	return &params, 0
}

/**
 * 配置应答包http头
 */
func handleResponseHeader(w http.ResponseWriter) {
	w.Header().Set("content-type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
}

func TestHandler(w http.ResponseWriter, r *http.Request) {
	/*defer func() {
		if err := recover(); err != nil {
			w.Write(util.Failed(util.ERROR_CODE_SYSTEM_CRASH))
		}
	}()*/

	handleResponseHeader(w)

	params, err := parseBody(r)
	if err != 0 {
		w.Write(util.Failed(err))
		return
	}

	//判断是否参数完整
	id, err0 := (*params)["id"].(float64)
	if err0  {
		testData, e := logic.DealTest((int64)(id))
		if e != 0 {
			w.Write(util.Failed(e))
		} else {
			w.Write(util.Success(testData))
		}
		return
	}
	w.Write(util.Failed(util.ERROR_CODE_PARAM_INVALID))
}

/**
 * 用户是否登录
 */
func isUserLogin(w http.ResponseWriter, r *http.Request) (string, bool) {
	//用户是否登录
	sessionId := GSession.CheckCookieValid(w, r)
	if len(sessionId) == 0 {
		return "", false
	}

	//获取用户地址
	address, e := GSession.GetSessionVal(sessionId, "user")
	if !e {
		return "", false
	}

	return address.(string), true
}

/**
 * 记录用户session
 */
func recordUserSession(w http.ResponseWriter, r *http.Request, address string) {
	GSession.EndSession(w, r)
	sessionId := GSession.StartSession(w, r)
	GSession.SetSessionVal(sessionId, "user", address)
}

/**
  ** 新用户注册
 */
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	/*defer func() {
			if err := recover(); err != nil {
				panicHandler(err)
				w.Write(util.Failed(util.ERROR_CODE_SYSTEM_CRASH))
			}
		}()*/

	handleResponseHeader(w)
	user, err := parseBody(r)
	if err != 0 {
		w.Write(util.Failed(err))
		return
	}

	//判断是否参数完整
	address, err0 := (*user)["address"].(string)
	email, err1 := (*user)["email"].(string)
	nick, err2 := (*user)["nick"].(string)
	sign, err3 := (*user)["sign"].(string)

	if err0 && err1 && err2 && err3 {
		//判断用户签名是否正确
		isUser := logic.IsUserValid(address, sign)
		if isUser {
			e := logic.RegisterUser(address, email, nick)
			if e != 0 {
				w.Write(util.Failed(e))
			} else {
				//登录成功
				recordUserSession(w, r, address)
				w.Write(util.Success(nil))
			}
		} else {
			w.Write(util.Failed(util.ERROR_CODE_OP_INVAILD))
		}

		return
	}
	w.Write(util.Failed(util.ERROR_CODE_PARAM_INVALID))
}

/**
 ** 用户登录
 */
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	/*defer func() {
		if err := recover(); err != nil {
			panicHandler(err)
			w.Write(util.Failed(util.ERROR_CODE_SYSTEM_CRASH))
		}
	}()*/

	handleResponseHeader(w)

	params, err := parseBody(r)
	if err != 0 {
		w.Write(util.Failed(err))
		return
	}

	//判断是否参数完整
	address, e0 := (*params)["address"].(string)
	sign, e1 := (*params)["sign"].(string)
	if e0 && e1 {
		//判断用户签名是否正确
		isUser := logic.IsUserValid(address, sign)
		if isUser {
			user, e := logic.QueryUserInfo(address)
			if e != 0 {
				w.Write(util.Failed(e))
			} else {
				//登录成功
				recordUserSession(w, r, address)
				w.Write(util.Success(user))
			}
		} else {
			w.Write(util.Failed(util.ERROR_CODE_OP_INVAILD))
		}

		return
	}

	w.Write(util.Failed(util.ERROR_CODE_PARAM_INVALID))
}

/**
 * 用户退出登陆
 */
func LogoutHandler(w http.ResponseWriter, r *http.Request) {

}

/**
 ** 查询用户信息
 */
func SelfInfoHandler(w http.ResponseWriter, r *http.Request) {
	/*defer func() {
		if err := recover(); err != nil {
			panicHandler(err)
			w.Write(util.Failed(util.ERROR_CODE_SYSTEM_CRASH))
		}
	}()*/

	handleResponseHeader(w)
	//用户是否登录
	address, e := isUserLogin(w, r)
	if !e {
		w.Write(util.Failed(util.ERROR_CODE_NOTLOGIN))
		return
	}

	user, ret := logic.QueryUserInfo(address)
	if ret != 0 {
		w.Write(util.Failed(ret))
	} else {
		w.Write(util.Success(user))
	}
}

/**
 * 查询市场bird列表
 */
func MarketBirdsHandler(w http.ResponseWriter, r *http.Request) {
	handleResponseHeader(w)

	params, err := parseBody(r)
	if err != 0 {
		w.Write(util.Failed(err))
		return
	}

	/*defer func() {
		if err := recover(); err != nil {
			panicHandler(err)
			w.Write(util.Failed(util.ERROR_CODE_SYSTEM_CRASH))
		}
	}()*/

	//判断是否参数完整
	startPage, e0 := (*params)["start_page"]
	pageSize, e1 := (*params)["page_size"]
	orderType, e2 := (*params)["order_type"]
	sortType, e3 := (*params)["sort_type"]

	util.Logger.Info("MarketBirdsHandler:", startPage, pageSize, orderType, sortType)

	if e0 && e1 && e2 && e3 {
		bird, e := logic.GetMarketBirds((int)(startPage.(float64)), (int)(pageSize.(float64)), (int)(orderType.(float64)), (int)(sortType.(float64)))
		if e != 0 {
			w.Write(util.Failed(e))
		} else {
			w.Write(util.Success(bird))
		}
		return
	}

	w.Write(util.Failed(util.ERROR_CODE_PARAM_INVALID))
}

/**
 ** 查询用户信息
 */
func UserInfoHandler(w http.ResponseWriter, r *http.Request) {
	/*defer func() {
		if err := recover(); err != nil {
			panicHandler(err)
			w.Write(util.Failed(util.ERROR_CODE_SYSTEM_CRASH))
		}
	}()*/

	handleResponseHeader(w)
	//用户是否登录
	_, e := isUserLogin(w, r)
	if !e {
		w.Write(util.Failed(util.ERROR_CODE_NOTLOGIN))
		return
	}

	user, err := parseBody(r)
	if err != 0 {
		w.Write(util.Failed(err))
		return
	}

	//判断是否参数完整
	address, err0 := (*user)["address"].(string)
	if err0 {
		user, e := logic.QueryUserInfo(address)
		if e != 0 {
			w.Write(util.Failed(e))
		} else {
			w.Write(util.Success(user))
		}
		return
	}
	w.Write(util.Failed(util.ERROR_CODE_PARAM_INVALID))
}

/**
 * 判断用户是否存在
 */
func UserExistHandler(w http.ResponseWriter, r *http.Request) {
	/*defer func() {
		if err := recover(); err != nil {
			panicHandler(err)
			w.Write(util.Failed(util.ERROR_CODE_SYSTEM_CRASH))
		}
	}()*/

	handleResponseHeader(w)
	user, err := parseBody(r)
	if err != 0 {
		w.Write(util.Failed(err))
		return
	}

	//判断是否参数完整
	address, err0 := (*user)["address"].(string)
	if err0 {
		user, e := logic.QueryUserInfo(address)
		if e != 0 {
			w.Write(util.Failed(e))
		} else {
			w.Write(util.Success(user))
		}
		return
	}
	w.Write(util.Failed(util.ERROR_CODE_PARAM_INVALID))
}

/**
 * 用户购买bird请求，返回调用智能合约的参数
 */
func BuyBirdHandler(w http.ResponseWriter, r *http.Request) {
	/*defer func() {
		if err := recover(); err != nil {
			panicHandler(err)
			w.Write(util.Failed(util.ERROR_CODE_SYSTEM_CRASH))
		}
	}()*/

	handleResponseHeader(w)

	//用户是否登录
	address, e := isUserLogin(w, r)
	if !e {
		w.Write(util.Failed(util.ERROR_CODE_NOTLOGIN))
		return
	}

	params, err := parseBody(r)
	if err != 0 {
		w.Write(util.Failed(err))
		return
	}

	//判断是否参数完整
	tokenId, e0 := (*params)["token_id"].(float64)
	if e0 {
		buyBird, e := logic.BuyBird(address, int64(tokenId))
		if e != 0 {
			w.Write(util.Failed(e))
		} else {
			w.Write(util.Success(buyBird))
		}
		return
	}

	w.Write(util.Failed(util.ERROR_CODE_PARAM_INVALID))
}

/**
 * 购买bird之后，用户记录txn hash
 */
func BuyBirdHashHandler(w http.ResponseWriter, r *http.Request) {
	/*defer func() {
		if err := recover(); err != nil {
			panicHandler(err)
			w.Write(util.Failed(util.ERROR_CODE_SYSTEM_CRASH))
		}
	}()*/

	handleResponseHeader(w)
	//用户是否登录
	address, e := isUserLogin(w, r)
	if !e {
		w.Write(util.Failed(util.ERROR_CODE_NOTLOGIN))
		return
	}

	params, err := parseBody(r)
	if err != 0 {
		w.Write(util.Failed(err))
		return
	}

	//判断是否参数完整
	tokenId, e0 := (*params)["token_id"].(float64)
	hash, e1 := (*params)["hash"].(string)
	if e0 && e1 {
		e := logic.BuyBirdHash(address, int64(tokenId), hash)
		if e != 0 {
			w.Write(util.Failed(e))
		} else {
			w.Write(util.Success(nil))
		}
		return
	}

	w.Write(util.Failed(util.ERROR_CODE_PARAM_INVALID))
}

/**
 * 销售bird
 */
func SellBirdHandler(w http.ResponseWriter, r *http.Request) {
	/*defer func() {
			if err := recover(); err != nil {
				panicHandler(err)
				w.Write(util.Failed(util.ERROR_CODE_SYSTEM_CRASH))
			}
		}()*/

	handleResponseHeader(w)

	params, err := parseBody(r)
	if err != 0 {
		w.Write(util.Failed(err))
		return
	}

	//用户是否登录
	address, e := isUserLogin(w, r)
	if !e {
		w.Write(util.Failed(util.ERROR_CODE_NOTLOGIN))
		return
	}

	//判断是否参数完整
	tokenId, e0 := (*params)["token_id"].(float64)
	price, e1 := (*params)["price"].(string)
	if e0 && e1 {
		e := logic.SellBird((int64)(tokenId), address, price)
		if e != 0 {
			w.Write(util.Failed(e))
		} else {
			w.Write(util.Success(""))
		}
		return
	}

	w.Write(util.Failed(util.ERROR_CODE_PARAM_INVALID))
}

/**
 * 获取bird详细信息
 */
func SellOffHandler(w http.ResponseWriter, r *http.Request) {
	/*defer func() {
			if err := recover(); err != nil {
				panicHandler(err)
				w.Write(util.Failed(util.ERROR_CODE_SYSTEM_CRASH))
			}
		}()*/

	handleResponseHeader(w)
	params, err := parseBody(r)
	if err != 0 {
		w.Write(util.Failed(err))
		return
	}

	//用户是否登录
	address, e := isUserLogin(w, r)
	if !e {
		w.Write(util.Failed(util.ERROR_CODE_NOTLOGIN))
		return
	}

	//判断是否参数完整
	tokenId, e0 := (*params)["token_id"].(float64)
	if e0 {
		e := logic.SellOff((int64)(tokenId), address)
		if e != 0 {
			w.Write(util.Failed(e))
		} else {
			w.Write(util.Success(nil))
		}
		return
	}

	w.Write(util.Failed(util.ERROR_CODE_PARAM_INVALID))
}

/**
 * 购买bird之后，用户记录txn hash
 */
func BirdInfoHandler(w http.ResponseWriter, r *http.Request) {
	/*defer func() {
		if err := recover(); err != nil {
			panicHandler(err)
			w.Write(util.Failed(util.ERROR_CODE_SYSTEM_CRASH))
		}
	}()*/

	handleResponseHeader(w)
	//用户是否登录
	_, e := isUserLogin(w, r)
	if !e {
		w.Write(util.Failed(util.ERROR_CODE_NOTLOGIN))
		return
	}

	params, err := parseBody(r)
	if err != 0 {
		w.Write(util.Failed(err))
		return
	}

	//判断是否参数完整
	tokenId, err0 := (*params)["token_id"].(float64)
	if err0 {
		bird, e := logic.QueryBirdById((int64)(tokenId))
		if e != 0 {
			w.Write(util.Failed(e))
		} else {
			w.Write(util.Success(bird))
		}
		return
	}

	w.Write(util.Failed(util.ERROR_CODE_PARAM_INVALID))
}

/**
 * 获取水果列表
 */
func FruitListHandler(w http.ResponseWriter, r *http.Request) {
	/*defer func() {
		if err := recover(); err != nil {
			panicHandler(err)
			w.Write(util.Failed(util.ERROR_CODE_SYSTEM_CRASH))
		}
	}()*/

	handleResponseHeader(w)
	//用户是否登录
	_, err := isUserLogin(w, r)
	if !err {
		w.Write(util.Failed(util.ERROR_CODE_NOTLOGIN))
		return
	}

	fruits, e := logic.QueryFruitList()
	if e != 0 {
		w.Write(util.Failed(e))
	} else {
		w.Write(util.Success(fruits))
	}
}

/**
 * 通过fruit类型获取水果信息
 */
func FruitInfoHandler(w http.ResponseWriter, r *http.Request) {
	/*defer func() {
		if err := recover(); err != nil {
			panicHandler(err)
			w.Write(util.Failed(util.ERROR_CODE_SYSTEM_CRASH))
		}
	}()*/

	handleResponseHeader(w)
	//用户是否登录
	_, e := isUserLogin(w, r)
	if !e {
		w.Write(util.Failed(util.ERROR_CODE_NOTLOGIN))
		return
	}

	params, err := parseBody(r)
	if err != 0 {
		w.Write(util.Failed(err))
		return
	}

	//判断是否参数完整
	t, e0 := (*params)["type"].(float64)
	if e0 {
		fruit, e := logic.QueryFruitByType((int)(t))
		if e != 0 {
			w.Write(util.Failed(e))
		} else {
			w.Write(util.Success(fruit))
		}
		return
	}

	w.Write(util.Failed(util.ERROR_CODE_PARAM_INVALID))
}

/**
 * 购买水果接口，返回调用智能合约的参数
 */
func BuyFruitHandler(w http.ResponseWriter, r *http.Request) {
	/*defer func() {
		if err := recover(); err != nil {
			panicHandler(err)
			w.Write(util.Failed(util.ERROR_CODE_SYSTEM_CRASH))
		}
	}()*/

	handleResponseHeader(w)
	//用户是否登录
	address, e := isUserLogin(w, r)
	if !e {
		w.Write(util.Failed(util.ERROR_CODE_NOTLOGIN))
		return
	}

	params, err := parseBody(r)
	if err != 0 {
		w.Write(util.Failed(err))
		return
	}

	//判断是否参数完整
	tokenId, e0 := (*params)["token_id"].(float64)
	t, e1 := (*params)["type"].(float64)
	if e0 && e1 {
		fruit, e := logic.BuyFruit(address, int64(tokenId), int64(t))
		if e != 0 {
			w.Write(util.Failed(e))
		} else {
			w.Write(util.Success(fruit))
		}
		return
	}

	w.Write(util.Failed(util.ERROR_CODE_PARAM_INVALID))
}

/**
 * 购买fruit之后，用户记录txn hash
 */
func BuyFruitHashHandler(w http.ResponseWriter, r *http.Request) {
	/*defer func() {
		if err := recover(); err != nil {
			panicHandler(err)
			w.Write(util.Failed(util.ERROR_CODE_SYSTEM_CRASH))
		}
	}()*/

	handleResponseHeader(w)
	//用户是否登录
	address, e := isUserLogin(w, r)
	if !e {
		w.Write(util.Failed(util.ERROR_CODE_NOTLOGIN))
		return
	}

	params, err := parseBody(r)
	if err != 0 {
		w.Write(util.Failed(err))
		return
	}

	//判断是否参数完整
	tokenId, e0 := (*params)["token_id"].(float64)
	t, e1 := (*params)["type"].(float64)
	hash, e2 := (*params)["hash"].(string)

	if e0 && e1 && e2 {
		e := logic.BuyFruitHash(address, int64(tokenId), int64(t), hash)
		if e != 0 {
			w.Write(util.Failed(e))
		} else {
			w.Write(util.Success(nil))
		}
		return
	}

	w.Write(util.Failed(util.ERROR_CODE_PARAM_INVALID))
}

/**
 * 获取用户自己的bird列表
 */
func MyBirdsHandler(w http.ResponseWriter, r *http.Request) {
	/*defer func() {
		if err := recover(); err != nil {
			panicHandler(err)
			w.Write(util.Failed(util.ERROR_CODE_SYSTEM_CRASH))
		}
	}()*/

	handleResponseHeader(w)

	//用户是否登录
	address, e := isUserLogin(w, r)
	if !e {
		w.Write(util.Failed(util.ERROR_CODE_NOTLOGIN))
		return
	}
	//address := "0x89b45023ea3d9a8f343b794fe2d2ce18d28ddf47"

	params, err := parseBody(r)
	if err != 0 {
		w.Write(util.Failed(err))
		return
	}

	startPage, err0 := (*params)["start_page"].(float64)
	pageSize, err1 := (*params)["page_size"].(float64)
	if err0 && err1 {
		birds, e0 := logic.QueryBirdsByAddress(address, int(startPage), int(pageSize))
		if e0 != 0 {
			w.Write(util.Failed(e0))
		} else {
			w.Write(util.Success(birds))
		}
		return
	}
	w.Write(util.Failed(util.ERROR_CODE_PARAM_INVALID))
}

/**
 * 获取用户可以参与PK的bird列表
 */
func MyBirdsPKListHandler(w http.ResponseWriter, r *http.Request) {
	/*defer func() {
		if err := recover(); err != nil {
			panicHandler(err)
			w.Write(util.Failed(util.ERROR_CODE_SYSTEM_CRASH))
		}
	}()*/

	handleResponseHeader(w)

	//用户是否登录
	address, e := isUserLogin(w, r)
	if !e {
		w.Write(util.Failed(util.ERROR_CODE_NOTLOGIN))
		return
	}
	//address := "0x715afba02987ed2f053cb7d071b8039b02fcbe0e"

	params, err := parseBody(r)
	if err != 0 {
		w.Write(util.Failed(err))
		return
	}

	startPage, err0 := (*params)["start_page"].(float64)
	pageSize, err1 := (*params)["page_size"].(float64)
	if err0 && err1 {
		birds, e0 := logic.QueryUserNormalBirds(address, int(startPage), int(pageSize))
		if e0 != 0 {
			w.Write(util.Failed(e0))
		} else {
			w.Write(util.Success(birds))
		}
		return
	}
	w.Write(util.Failed(util.ERROR_CODE_PARAM_INVALID))
}

/**
 * 获取其他用户bird列表
 */
func UserBirdListHandler(w http.ResponseWriter, r *http.Request) {
	/*defer func() {
		if err := recover(); err != nil {
			panicHandler(err)
			w.Write(util.Failed(util.ERROR_CODE_SYSTEM_CRASH))
		}
	}()*/

	handleResponseHeader(w)

	//用户是否登录
	_, e := isUserLogin(w, r)
	if !e {
		w.Write(util.Failed(util.ERROR_CODE_NOTLOGIN))
		return
	}

	params, err := parseBody(r)
	if err != 0 {
		w.Write(util.Failed(err))
		return
	}

	//判断是否参数完整
	address, e0 := (*params)["address"].(string)
	startPage, e1 := (*params)["start_page"].(float64)
	pageSize, e2 := (*params)["page_size"].(float64)
	if e0 && e1 && e2 {
		birds, ret := logic.QueryBirdsByAddress(address, int(startPage), int(pageSize))
		if ret != 0 {
			w.Write(util.Failed(ret))
		} else {
			w.Write(util.Success(birds))
		}
		return
	}

	w.Write(util.Failed(util.ERROR_CODE_PARAM_INVALID))
}

/**
 * bird pk接口，返回调用智能合约的参数
 */
func PKHandler(w http.ResponseWriter, r *http.Request) {
	/*defer func() {
		if err := recover(); err != nil {
			panicHandler(err)
			w.Write(util.Failed(util.ERROR_CODE_SYSTEM_CRASH))
		}
	}()*/

	handleResponseHeader(w)
	//用户是否登录
	address, e := isUserLogin(w, r)
	if !e {
		w.Write(util.Failed(util.ERROR_CODE_NOTLOGIN))
		return
	}

	user, err := parseBody(r)
	if err != 0 {
		w.Write(util.Failed(err))
		return
	}

	//判断是否参数完整
	challengerId, e0 := (*user)["challengerId"].(float64)
	resisterId, e1 := (*user)["resisterId"].(float64)
	if e0 && e1 {
		birdPK, e := logic.PKBird(address, int64(challengerId), int64(resisterId))
		if e != 0 {
			w.Write(util.Failed(e))
		} else {
			w.Write(util.Success(birdPK))
		}
		return
	}
	w.Write(util.Failed(util.ERROR_CODE_PARAM_INVALID))
}

/**
 * bird pk智能合约txn id记录接口
 */
func PKHashHandler(w http.ResponseWriter, r *http.Request) {
	/*defer func() {
			if err := recover(); err != nil {
				panicHandler(err)
				w.Write(util.Failed(util.ERROR_CODE_SYSTEM_CRASH))
			}
		}()*/

	handleResponseHeader(w)
	//用户是否登录
	address, e := isUserLogin(w, r)
	if !e {
		w.Write(util.Failed(util.ERROR_CODE_NOTLOGIN))
		return
	}

	user, err := parseBody(r)
	if err != 0 {
		w.Write(util.Failed(err))
		return
	}

	//判断是否参数完整
	challengerId, e0 := (*user)["challengerId"].(float64)
	resisterId, e1 := (*user)["resisterId"].(float64)
	hash, e2 := (*user)["hash"].(string)
	if e0 && e1 && e2 {
		e := logic.PKBirdHash(address, int64(challengerId), int64(resisterId), hash)
		if e != 0 {
			w.Write(util.Failed(e))
		} else {
			w.Write(util.Success(nil))
		}
		return
	}
	w.Write(util.Failed(util.ERROR_CODE_PARAM_INVALID))
}

/**
 * 返回除用户之外的能够pk的bird列表
 */
func PKListHandler(w http.ResponseWriter, r *http.Request) {
	/*defer func() {
		if err := recover(); err != nil {
			panicHandler(err)
			w.Write(util.Failed(util.ERROR_CODE_SYSTEM_CRASH))
		}
	}()*/

	handleResponseHeader(w)//用户是否登录
	address, e := isUserLogin(w, r)
	if !e {
		w.Write(util.Failed(util.ERROR_CODE_NOTLOGIN))
		return
	}

	params, err := parseBody(r)
	if err != 0 {
		w.Write(util.Failed(err))
		return
	}

	//判断是否参数完整
	startPage, e0 := (*params)["start_page"]
	pageSize, e1 := (*params)["page_size"]
	//orderType, e2 := (*params)["order_type"]
	//sortType, e3 := (*params)["sort_type"]

	util.Logger.Info("PKListHandler:", startPage, pageSize)

	if e0 && e1 {
		bird, e := logic.GetPKListBirds(address, (int)(startPage.(float64)), (int)(pageSize.(float64)), 1, 1)
		if e != 0 {
			w.Write(util.Failed(e))
		} else {
			w.Write(util.Success(bird))
		}
		return
	}

	w.Write(util.Failed(util.ERROR_CODE_PARAM_INVALID))
}

/**
 * 按照体重排序的bird列表
 */
func RankHandler(w http.ResponseWriter, r *http.Request) {
	/*defer func() {
		if err := recover(); err != nil {
			panicHandler(err)
			w.Write(util.Failed(util.ERROR_CODE_SYSTEM_CRASH))
		}
	}()*/

	handleResponseHeader(w)
	//用户是否登录
	_, e := isUserLogin(w, r)
	if !e {
		w.Write(util.Failed(util.ERROR_CODE_NOTLOGIN))
		return
	}

	params, err := parseBody(r)
	if err != 0 {
		w.Write(util.Failed(err))
		return
	}

	//判断是否参数完整
	startPage, e0 := (*params)["start_page"].(float64)
	pageSize, e1 := (*params)["page_size"].(float64)
	if e0 && e1 {
		birdPK, e := logic.RankBird(int(startPage), int(pageSize))
		if e != 0 {
			w.Write(util.Failed(e))
		} else {
			w.Write(util.Success(birdPK))
		}
		return
	}
	w.Write(util.Failed(util.ERROR_CODE_PARAM_INVALID))
}

/**
 ** 抓鸟txn hash记录接口
 */
func CatchHashHandler(w http.ResponseWriter, r *http.Request) {
	/*defer func() {
		if err := recover(); err != nil {
			panicHandler(err)
			w.Write(util.Failed(util.ERROR_CODE_SYSTEM_CRASH))
		}
	}()*/

	handleResponseHeader(w)

	//用户是否登录
	address, e := isUserLogin(w, r)
	if !e {
		w.Write(util.Failed(util.ERROR_CODE_NOTLOGIN))
		return
	}

	user, err := parseBody(r)
	if err != 0 {
		w.Write(util.Failed(err))
		return
	}

	//判断是否参数完整
	hash, e0 := (*user)["hash"].(string)
	if e0 {
		e := logic.CatchBirdHash(address, hash)
		if e != 0 {
			w.Write(util.Failed(e))
		} else {
			w.Write(util.Success(nil))
		}
		return
	}
	w.Write(util.Failed(util.ERROR_CODE_PARAM_INVALID))
}

/**
 * bird创建接口，根据基因生成图片等
 */
func GenerateBirdHandler(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			w.Write(util.Failed(util.ERROR_CODE_SYSTEM_CRASH))
		}
	}()

	handleResponseHeader(w)
	params, err := parseBody(r)
	if err != 0 {
		w.Write(util.Failed(err))
		return
	}

	//判断是否参数完整
	id, _ := (*params)["token_id"].(float64)

	_, e := logic.GenerateBird((int64)(id))
	if e != 0 {
		w.Write(util.Failed(e))
	} else {
		w.Write(util.Success(nil))
	}
}

/**
 * 获取catch记录
 */
func CatchBirdRecordHandler(w http.ResponseWriter, r *http.Request) {
	/*defer func() {
		if err := recover(); err != nil {
			panicHandler(err)
			w.Write(util.Failed(util.ERROR_CODE_SYSTEM_CRASH))
		}
	}()*/

	handleResponseHeader(w)
	//用户是否登录
	_, e := isUserLogin(w, r)
	if !e {
		w.Write(util.Failed(util.ERROR_CODE_NOTLOGIN))
		return
	}

	params, err := parseBody(r)
	if err != 0 {
		w.Write(util.Failed(err))
		return
	}

	//判断是否参数完整
	startPage, e0 := (*params)["start_page"].(float64)
	pageSize, e1 := (*params)["page_size"].(float64)
	if e0 && e1 {
		birdPK, e := logic.QueryQueryCatchRecordInWeek(int(startPage), int(pageSize))
		if e != 0 {
			w.Write(util.Failed(e))
		} else {
			w.Write(util.Success(birdPK))
		}
		return
	}
	w.Write(util.Failed(util.ERROR_CODE_PARAM_INVALID))
}

/**
 * 获取buy bird记录
 */
func BuyBirdRecordHandler(w http.ResponseWriter, r *http.Request) {
	/*defer func() {
		if err := recover(); err != nil {
			panicHandler(err)
			w.Write(util.Failed(util.ERROR_CODE_SYSTEM_CRASH))
		}
	}()*/

	handleResponseHeader(w)
	//用户是否登录
	_, e := isUserLogin(w, r)
	if !e {
		w.Write(util.Failed(util.ERROR_CODE_NOTLOGIN))
		return
	}

	params, err := parseBody(r)
	if err != 0 {
		w.Write(util.Failed(err))
		return
	}

	//判断是否参数完整
	startPage, e0 := (*params)["start_page"].(float64)
	pageSize, e1 := (*params)["page_size"].(float64)
	if e0 && e1 {
		birdPK, e := logic.QueryBuyBirdRecordInWeek(int(startPage), int(pageSize))
		if e != 0 {
			w.Write(util.Failed(e))
		} else {
			w.Write(util.Success(birdPK))
		}
		return
	}
	w.Write(util.Failed(util.ERROR_CODE_PARAM_INVALID))
}

/**
 * 获取buy fruit记录
 */
func BuyFruitRecordHandler(w http.ResponseWriter, r *http.Request) {
	/*defer func() {
		if err := recover(); err != nil {
			panicHandler(err)
			w.Write(util.Failed(util.ERROR_CODE_SYSTEM_CRASH))
		}
	}()*/

	handleResponseHeader(w)
	//用户是否登录
	_, e := isUserLogin(w, r)
	if !e {
		w.Write(util.Failed(util.ERROR_CODE_NOTLOGIN))
		return
	}

	params, err := parseBody(r)
	if err != 0 {
		w.Write(util.Failed(err))
		return
	}

	//判断是否参数完整
	startPage, e0 := (*params)["start_page"].(float64)
	pageSize, e1 := (*params)["page_size"].(float64)
	if e0 && e1 {
		birdPK, e := logic.QueryBuyFruitRecordInWeek(int(startPage), int(pageSize))
		if e != 0 {
			w.Write(util.Failed(e))
		} else {
			w.Write(util.Success(birdPK))
		}
		return
	}
	w.Write(util.Failed(util.ERROR_CODE_PARAM_INVALID))
}

/**
 * 交易统计接口
 */
func TxStatisticsHandler(w http.ResponseWriter, r *http.Request) {
	/*defer func() {
		if err := recover(); err != nil {
			panicHandler(err)
			w.Write(util.Failed(util.ERROR_CODE_SYSTEM_CRASH))
		}
	}()*/

	handleResponseHeader(w)
	//用户是否登录
	_, e := isUserLogin(w, r)
	if !e {
		w.Write(util.Failed(util.ERROR_CODE_NOTLOGIN))
		return
	}

	c, p, err := logic.QueryTxStatistics()
	if err != 0 {
		w.Write(util.Failed(err))
	} else {
		var profit = make(map[string]interface{})
		profit["tx_count"] = c
		profit["total_price"] = p

		w.Write(util.Success(profit))
	}
}

/**
 * 交易统计接口
 */
func CatchBirdRecordInfoHandler(w http.ResponseWriter, r *http.Request) {
	/*defer func() {
		if err := recover(); err != nil {
			panicHandler(err)
			w.Write(util.Failed(util.ERROR_CODE_SYSTEM_CRASH))
		}
	}()*/

	handleResponseHeader(w)
	//用户是否登录
	address, e := isUserLogin(w, r)
	if !e {
		w.Write(util.Failed(util.ERROR_CODE_NOTLOGIN))
		return
	}

	params, err := parseBody(r)
	if err != 0 {
		w.Write(util.Failed(err))
		return
	}

	//判断是否参数完整
	startPage, e0 := (*params)["start_page"].(float64)
	pageSize, e1 := (*params)["page_size"].(float64)
	if e0 && e1 {
		birdPK, e := logic.QueryCatchRecordByAddress(address, int(startPage), int(pageSize))
		if e != 0 {
			w.Write(util.Failed(e))
		} else {
			w.Write(util.Success(birdPK))
		}
		return
	}
	w.Write(util.Failed(util.ERROR_CODE_PARAM_INVALID))
}

func BuyBirdRecordInfoHandler(w http.ResponseWriter, r *http.Request) {
	/*defer func() {
		if err := recover(); err != nil {
			panicHandler(err)
			w.Write(util.Failed(util.ERROR_CODE_SYSTEM_CRASH))
		}
	}()*/

	handleResponseHeader(w)
	//用户是否登录
	_, e := isUserLogin(w, r)
	if !e {
		w.Write(util.Failed(util.ERROR_CODE_NOTLOGIN))
		return
	}

	params, err := parseBody(r)
	if err != 0 {
		w.Write(util.Failed(err))
		return
	}

	//判断是否参数完整
	tokenId, e0 := (*params)["token_id"].(float64)
	startPage, e0 := (*params)["start_page"].(float64)
	pageSize, e1 := (*params)["page_size"].(float64)
	if e0 && e1 {
		birdPK, e := logic.QueryBuyBirdRecordByTokenId(int64(tokenId), int(startPage), int(pageSize))
		if e != 0 {
			w.Write(util.Failed(e))
		} else {
			w.Write(util.Success(birdPK))
		}
		return
	}
	w.Write(util.Failed(util.ERROR_CODE_PARAM_INVALID))
}

func BuyFruitRecordInfoHandler(w http.ResponseWriter, r *http.Request) {
	/*defer func() {
		if err := recover(); err != nil {
			panicHandler(err)
			w.Write(util.Failed(util.ERROR_CODE_SYSTEM_CRASH))
		}
	}()*/

	handleResponseHeader(w)
	//用户是否登录
	_, e := isUserLogin(w, r)
	if !e {
		w.Write(util.Failed(util.ERROR_CODE_NOTLOGIN))
		return
	}

	params, err := parseBody(r)
	if err != 0 {
		w.Write(util.Failed(err))
		return
	}

	//判断是否参数完整
	tokenId, e0 := (*params)["token_id"].(float64)
	startPage, e0 := (*params)["start_page"].(float64)
	pageSize, e1 := (*params)["page_size"].(float64)
	if e0 && e1 {
		birdPK, e := logic.QueryBuyFruitRecordByTokenId(int64(tokenId), int(startPage), int(pageSize))
		if e != 0 {
			w.Write(util.Failed(e))
		} else {
			w.Write(util.Success(birdPK))
		}
		return
	}
	w.Write(util.Failed(util.ERROR_CODE_PARAM_INVALID))
}

func PKRecordHandler(w http.ResponseWriter, r *http.Request) {
	/*defer func() {
		if err := recover(); err != nil {
			panicHandler(err)
			w.Write(util.Failed(util.ERROR_CODE_SYSTEM_CRASH))
		}
	}()*/

	handleResponseHeader(w)
	//用户是否登录
	_, e := isUserLogin(w, r)
	if !e {
		w.Write(util.Failed(util.ERROR_CODE_NOTLOGIN))
		return
	}

	params, err := parseBody(r)
	if err != 0 {
		w.Write(util.Failed(err))
		return
	}

	//判断是否参数完整
	tokenId, e0 := (*params)["token_id"].(float64)
	startPage, e0 := (*params)["start_page"].(float64)
	pageSize, e1 := (*params)["page_size"].(float64)
	if e0 && e1 {
		birdPK, e := logic.QueryPKRecordByTokenId(int64(tokenId), int(startPage), int(pageSize))
		if e != 0 {
			w.Write(util.Failed(e))
		} else {
			w.Write(util.Success(birdPK))
		}
		return
	}
	w.Write(util.Failed(util.ERROR_CODE_PARAM_INVALID))
}

func PKInfoRecordHandler(w http.ResponseWriter, r *http.Request) {
	/*defer func() {
		if err := recover(); err != nil {
			panicHandler(err)
			w.Write(util.Failed(util.ERROR_CODE_SYSTEM_CRASH))
		}
	}()*/

	handleResponseHeader(w)
	//用户是否登录
	_, e := isUserLogin(w, r)
	if !e {
		w.Write(util.Failed(util.ERROR_CODE_NOTLOGIN))
		return
	}

	params, err := parseBody(r)
	if err != 0 {
		w.Write(util.Failed(err))
		return
	}

	//判断是否参数完整
	hash, e0 := (*params)["hash"].(string)
	if e0 {
		birdPK, e := logic.QueryPKInfoRecordByHash(hash)
		if e != 0 {
			w.Write(util.Failed(e))
		} else {
			w.Write(util.Success(birdPK))
		}
		return
	}
	w.Write(util.Failed(util.ERROR_CODE_PARAM_INVALID))
}

/**
 * 分润统计
 */
func ProfitStatisticsRecordHandler(w http.ResponseWriter, r *http.Request) {
	/*defer func() {
		if err := recover(); err != nil {
			panicHandler(err)
			w.Write(util.Failed(util.ERROR_CODE_SYSTEM_CRASH))
		}
	}()*/

	handleResponseHeader(w)
	//用户是否登录
	_, e := isUserLogin(w, r)
	if !e {
		w.Write(util.Failed(util.ERROR_CODE_NOTLOGIN))
		return
	}

	params, err := parseBody(r)
	if err != 0 {
		w.Write(util.Failed(err))
		return
	}

	//判断是否参数完整
	week, e0 := (*params)["week"].(float64)
	if !e0 {
		week = 0
	}
	birdPK, err := logic.QueryProfitStatisticsRecord(int(week))
	if err != 0 {
		w.Write(util.Failed(err))
	} else {
		w.Write(util.Success(birdPK))
	}
}

/**
 * 分润列表
 */
func ProfitListRecordHandler(w http.ResponseWriter, r *http.Request) {
	/*defer func() {
		if err := recover(); err != nil {
			panicHandler(err)
			w.Write(util.Failed(util.ERROR_CODE_SYSTEM_CRASH))
		}
	}()*/

	handleResponseHeader(w)
	//用户是否登录
	_, e := isUserLogin(w, r)
	if !e {
		w.Write(util.Failed(util.ERROR_CODE_NOTLOGIN))
		return
	}

	params, err := parseBody(r)
	if err != 0 {
		w.Write(util.Failed(err))
		return
	}

	//判断是否参数完整
	week, e0 := (*params)["week"].(float64)
	if !e0 {
		week = 0
	}
	birdPK, err := logic.QueryProfitListRecord(int(week))
	if err != 0 {
		w.Write(util.Failed(err))
	} else {
		w.Write(util.Success(birdPK))
	}
}

/**
 * 分润记录
 */
func ProfitRecordHandler(w http.ResponseWriter, r *http.Request) {
	/*defer func() {
		if err := recover(); err != nil {
			panicHandler(err)
			w.Write(util.Failed(util.ERROR_CODE_SYSTEM_CRASH))
		}
	}()*/

	handleResponseHeader(w)
	//用户是否登录
	_, e := isUserLogin(w, r)
	if !e {
		w.Write(util.Failed(util.ERROR_CODE_NOTLOGIN))
		return
	}

	params, err := parseBody(r)
	if err != 0 {
		w.Write(util.Failed(err))
		return
	}

	//判断是否参数完整
	week, e0 := (*params)["week"].(float64)
	if !e0 {
		week = 0
	}
	birdPK, err := logic.QueryProfitRecord(int(week))
	if err != 0 {
		w.Write(util.Failed(err))
	} else {
		w.Write(util.Success(birdPK))
	}
}