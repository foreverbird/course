package util

import (
	"strings"
	"math/big"
	"math"
	"strconv"
)

/**
 * email 格式检查
 */
func EmailDesens(email string) (result string) {
	result = "***"

	if strings.Contains(email,"@") {
		res := strings.Split(email,"@")
		nameLen := len(res[0])
		if nameLen == 1 {
			result = res[0] + "***@" + res[1]
		} else if nameLen == 2 {
			prex := Substr(email,0,1)
			result = prex + "***@" + res[1]
		} else if nameLen > 2 {
			prex := Substr(email,0,2)
			result = prex + "***@" + res[1]
		}
	}
	return
}

/**
 * 获取子字符串
 */
func Substr(str string, start int, end int) string {
	rs := []rune(str)
	return string(rs[start:end])
}

/**
 * 以太币转为wei
 */
func Eth2Wei(s string) string {
	//这是处理位数的代码段 number.Big()
	v, _ := strconv.ParseFloat(s, 10)
	tenDecimal := big.NewFloat(math.Pow(10, float64(18)))
	value := big.NewFloat(v)
	convertAmount, _ := new(big.Float).Mul(tenDecimal, value).Int(&big.Int{})

	return convertAmount.String()
}

func Wei2Eth(s string) string {
	//这是处理位数的代码段 number.Big()
	v, _ := strconv.ParseFloat(s, 10)
	tenDecimal := big.NewFloat(math.Pow(10, float64(18)))
	value := big.NewFloat(v)
	convertAmount := new(big.Float).Quo(value, tenDecimal)

	return convertAmount.String()
}
