package logic

import (
	"bird/dao"
	"bird/data"
	"bird/util"
	"strings"
	"time"
)

func QueryFruitList() (*[]data.Fruit, int) {
	// 查询db获取fruits
	fruits, err := dao.QueryFruits()
	if !err {
		return nil, util.ERROR_CODE_DB_FAILED
	}

	return fruits, util.ERROR_CODE_SUCCESS
}

/**
 * 根据fruit类型获取fruit详情
 */
func QueryFruitByType(t int) (*data.Fruit, int) {
	fruit, err := dao.QueryFruitByType(t)
	if !err {
		return nil, util.ERROR_CODE_DB_FAILED
	}

	return fruit, util.ERROR_CODE_SUCCESS
}

func BuyFruit(address string, tokenId int64, t int64) (*data.BuyFruit, int) {
	//challenger bird存在
	bird := dao.QueryBirdByTokenId(tokenId)
	if bird == nil {
		return nil, util.ERROR_CODE_NO_BIRD
	}

	//非本人操作
	if strings.Compare(bird.Owner, address) != 0 {
		return nil, util.ERROR_CODE_OP_INVAILD
	}

	// bird的状态为normal时才能吃水果
	if bird.Status != data.Bird_Status_Normal {
		return nil, util.ERROR_CODE_BIRD_STATUS
	}

	//检查bird上次吃水果时间
	n := time.Now().Unix()
	if n - bird.EatFruitTime < 24 * 60 * 60 {
		return nil, util.ERROR_CODE_EAT_FRUIT_TIME
	}

	// 吃水果类型是否有效
	fruit, err := dao.QueryFruitByType(int(t))
	if !err {
		return nil, util.ERROR_CODE_DB_FAILED
	}

	// 构造返回结构体
	buyFruit := data.BuyFruit{
		Id:    tokenId,
		Type:  int(t),
		Price: fruit.Price,
		Sign:  util.Sign3_1(tokenId, int8(t), fruit.Price),
	}
	return &buyFruit, util.ERROR_CODE_SUCCESS
}

func BuyFruitHash(address string, tokenId int64, t int64, hash string) (int) {
	bird := dao.QueryBirdByTokenId(tokenId)
	if bird == nil {
		return util.ERROR_CODE_NO_BIRD
	}

	//当前address是否拥有该slot
	if strings.Compare(address, bird.Owner) != 0 {
		return util.ERROR_CODE_OP_INVAILD
	}

	// bird的状态为normal时才能吃水果
	if bird.Status != data.Bird_Status_Normal {
		return util.ERROR_CODE_BIRD_STATUS
	}

	// 记录txn hash，为了task定时任务检测
	err := dao.InsertFruitTxn(tokenId, int(t), address, hash, data.Txn_unconfirmed)
	if !err {
		return util.ERROR_CODE_DB_FAILED
	}

	//更新状态
	dao.UpdateBirdStatus(tokenId, data.Bird_Status_Eat_Unconfirmed)

	return util.ERROR_CODE_SUCCESS
}