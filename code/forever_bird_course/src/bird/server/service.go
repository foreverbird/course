package server

import (
	"bird/logic"
	"net/http"
	"os"
	"fmt"
	"strings"
	"runtime"
	"bird/log"
	"bird/util"
)

/**
 * 获取静态目录和端口号
 */
func getFileServerPath() (string, string) {
	os := runtime.GOOS
	fmt.Println("os:", os)

	if strings.Compare(os, "windows") == 0 {
		return "D:/project/private/golang/bird_forever_fe/fe", ":3097"
	}
	return "/home/work/bird/deploy/webroot", ":3097"
}

/**
  ** 初始化路由规则
 */
func initRouter() {
	http.Handle("/api/test", http.HandlerFunc(TestHandler))

	//用户注册
	http.Handle("/api/register", http.HandlerFunc(RegisterHandler))
	http.Handle("/api/login", http.HandlerFunc(LoginHandler))
	http.Handle("/api/logout", http.HandlerFunc(LogoutHandler))

	http.Handle("/api/bird/info", http.HandlerFunc(BirdInfoHandler))
	http.Handle("/api/fruit/list", http.HandlerFunc(FruitListHandler))
	http.Handle("/api/fruit/info", http.HandlerFunc(FruitInfoHandler))
	http.Handle("/api/buy/fruit", http.HandlerFunc(BuyFruitHandler))
	http.Handle("/api/buy/fruit/hash", http.HandlerFunc(BuyFruitHashHandler))

	http.Handle("/api/user/exist", http.HandlerFunc(UserExistHandler))
	http.Handle("/api/me", http.HandlerFunc(SelfInfoHandler))
	http.Handle("/api/user", http.HandlerFunc(UserInfoHandler))
	http.Handle("/api/mybirds", http.HandlerFunc(MyBirdsHandler))
	http.Handle("/api/mybirds/pk", http.HandlerFunc(MyBirdsPKListHandler))
	http.Handle("/api/user/birdlist", http.HandlerFunc(UserBirdListHandler))
	http.Handle("/api/market/birds", http.HandlerFunc(MarketBirdsHandler))
	http.Handle("/api/rank", http.HandlerFunc(RankHandler))

	http.Handle("/api/catch/hash", http.HandlerFunc(CatchHashHandler))
	http.Handle("/api/pk", http.HandlerFunc(PKHandler))
	http.Handle("/api/pk/hash", http.HandlerFunc(PKHashHandler))
	http.Handle("/api/pk/list", http.HandlerFunc(PKListHandler))

	http.Handle("/api/buy/bird", http.HandlerFunc(BuyBirdHandler))
	http.Handle("/api/buy/bird/hash", http.HandlerFunc(BuyBirdHashHandler))
	http.Handle("/api/sell/bird", http.HandlerFunc(SellBirdHandler))
	http.Handle("/api/sell/off", http.HandlerFunc(SellOffHandler))

	http.Handle("/api/event/generate/bird", http.HandlerFunc(GenerateBirdHandler))

	http.Handle("/api/record/catch/bird", http.HandlerFunc(CatchBirdRecordHandler))
	http.Handle("/api/record/buy/bird", http.HandlerFunc(BuyBirdRecordHandler))
	http.Handle("/api/record/buy/fruit", http.HandlerFunc(BuyFruitRecordHandler))

	http.Handle("/api/record/tx/statistics", http.HandlerFunc(TxStatisticsHandler))

	http.Handle("/api/record/catch/bird/info", http.HandlerFunc(CatchBirdRecordInfoHandler))
	http.Handle("/api/record/buy/bird/info", http.HandlerFunc(BuyBirdRecordInfoHandler))
	http.Handle("/api/record/buy/fruit/info", http.HandlerFunc(BuyFruitRecordInfoHandler))
	http.Handle("/api/record/pk", http.HandlerFunc(PKRecordHandler))
	http.Handle("/api/record/pk/info", http.HandlerFunc(PKInfoRecordHandler))

	http.Handle("/api/record/profit/statistics", http.HandlerFunc(ProfitStatisticsRecordHandler))
	http.Handle("/api/record/profit/list", http.HandlerFunc(ProfitListRecordHandler))
	http.Handle("/api/record/profit", http.HandlerFunc(ProfitRecordHandler))

	http.Handle("/api/sign", http.HandlerFunc(SignHandler))
}

var GSession  *logic.SessionMgr = nil //session管理器

/**
 * 开启服务
 */
func StartServer() {
	// 设置日志
	log.LogTo("stdout", "INFO")

	util.Logger.Info("start server")

	fileServerPath, port:= getFileServerPath()

	// 初始化session管理
	GSession = logic.NewSessionMgr("token", 30 * 60)

	// 设置静态目录
	fsh := http.FileServer(http.Dir(fileServerPath))
	http.Handle("/", http.StripPrefix("/", fsh))

	//初始化路由表
	initRouter()

	// 初始化bird创建类对象，加载bird配置
	logic.BirdGeneratorInstance()
	// listener服务
	err := http.ListenAndServe(port, nil)
	if err != nil {
		os.Exit(1)
	}
}

