package data

const (
	Txn_unconfirmed = 0 //未确定订单是否存在
	Txn_mining      = 1 //订单存在，但是还没有挖矿打包
	Txn_success     = 2 //智能合约调用成功
	Txn_failed      = 3 //智能合约调用失败
	Txn_invalid     = 4 //无效订单
)