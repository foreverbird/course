package logic

import (
	"bird/dao"
	"bird/util"
	"bird/data"
)

func DealTest(id int64) (*data.TestData, int) {
	testData := dao.QueryOneTest(id)
	return testData, util.ERROR_CODE_SUCCESS
}
