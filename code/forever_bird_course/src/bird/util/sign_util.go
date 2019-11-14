package util

import (
	"bytes"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
)

const Sign_contracts_salt = "forever_bird"

/**
 * int256转为对应16进制
 */
func int256toHex(number string) string {
	value := big.NewInt(256)
	value, _ = value.SetString(number, 10)
	return fmt.Sprintf("%064x", value)
}

/**
 * int64转为对应16进制
 */
func int64toHex(number int64) string {
	return fmt.Sprintf("%016x", number)
}

/**
 * int8转为对应16进制
 */
func int8toHex(number int8) string {
	return fmt.Sprintf("%02x", number)
}

/**
 * int64、string类型计算签名
 */
func Sign(p1 int64, p2 string) string {
	// 先转为16进制字符串
	p1Hex := int64toHex(p1)
	p2Hex := int256toHex(p2)

	// 参与hash的message放入byte数组
	b := bytes.Buffer{}
	b.Write(common.Hex2Bytes(p1Hex))
	b.Write(common.Hex2Bytes(p2Hex))
	b.Write([]byte(Sign_contracts_salt))

	// 计算hash
	h := crypto.Keccak256Hash([]byte(b.String()))
	return h.String()
}

/**
 * int64、int64类型计算签名
 */
func Sign2Int64(p1 int64, p2 int64) string {
	// 先转为16进制字符串
	p1Hex := int64toHex(p1)
	p2Hex := int64toHex(p2)

	// 参与hash的message放入byte数组
	b := bytes.Buffer{}
	b.Write(common.Hex2Bytes(p1Hex))
	b.Write(common.Hex2Bytes(p2Hex))
	b.Write([]byte(Sign_contracts_salt))

	// 计算hash
	h := crypto.Keccak256Hash([]byte(b.String()))
	return h.String()
}

/**
 * int64、string、string类型计算签名
 */
func Sign2(p1 int64, p2 string, p3 string) string {
	// 先转为16进制字符串
	p1Hex := int64toHex(p1)
	p2Hex := int256toHex(p2)
	p3Hex := int256toHex(p3)

	// 参与hash的message放入byte数组
	b := bytes.Buffer{}
	b.Write(common.Hex2Bytes(p1Hex))
	b.Write(common.Hex2Bytes(p2Hex))
	b.Write(common.Hex2Bytes(p3Hex))
	b.Write([]byte(Sign_contracts_salt))

	// 计算hash
	h := crypto.Keccak256Hash([]byte(b.String()))
	return h.String()
}

/**
 * int64、int8、string类型计算签名
 */
func Sign3_1(p1 int64, p2 int8, p3 string) string {
	// 先转为16进制字符串
	p1Hex := int64toHex(p1)
	p2Hex := int8toHex(p2)
	p3Hex := int256toHex(p3)

	// 参与hash的message放入byte数组
	b := bytes.Buffer{}
	b.Write(common.Hex2Bytes(p1Hex))
	b.Write(common.Hex2Bytes(p2Hex))
	b.Write(common.Hex2Bytes(p3Hex))
	b.Write([]byte(Sign_contracts_salt))

	// 计算hash
	h := crypto.Keccak256Hash([]byte(b.String()))
	return h.String()
}

/**
 * 通过web端hash还原地址
 */
func RecoverTypedSignature(msg string, sig string) (string, bool) {
	//  计算`message`对应的Keccak256 Hash值
	message := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(msg), msg)
	data := crypto.Keccak256Hash([]byte(message))
	sign := hexutil.MustDecode(sig)

	if sign[64] != 27 && sign[64] != 28 {
		return "error", false
	}
	sign[64] -= 27

	// 恢复出数字签名所用的公钥
	recoveredPub, err := crypto.Ecrecover(data.Bytes(), sign)
	if err != nil {
		return "error", false
	}

	// 转换为`ecdsa.PublicKey{}`类型中的点坐标X,Y
	pubKey := crypto.ToECDSAPub(recoveredPub)

	// 获取地址
	recoveredAddr := crypto.PubkeyToAddress(*pubKey)

	//地址转为十六进制字符串
	return hexutil.Encode(recoveredAddr.Bytes()), true
}