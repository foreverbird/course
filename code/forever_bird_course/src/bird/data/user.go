package data

/**
 ** 用户基本信息
 */
type User struct {
	//db中标识
	Id int64
	//eth 地址
	Address string `json:"address"`
	Email string `json:"email"`
	//昵称
	Nick string `json:"nick"`
	//头像标识
	Image string `json:"image"`
}
