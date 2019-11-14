package util

import (
	"math/big"
)

const Trade_transaction_Fee_rate_n = "3"
const Trade_transaction_Fee_rate_d = "100"


func CountTransactionFee(price string) string {
	value := big.NewInt(256)
	value, _ = value.SetString(price, 10)

	feeN := big.NewInt(256)
	feeN, _ = feeN.SetString(Trade_transaction_Fee_rate_n, 10)

	feeD := big.NewInt(256)
	feeD, _ = feeD.SetString(Trade_transaction_Fee_rate_d, 10)

	value = value.Mul(value, feeN)
	value = value.Div(value, feeD)

	return value.String()
}