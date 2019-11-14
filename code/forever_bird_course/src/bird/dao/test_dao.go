package dao

import (
	"strconv"
	"bird/data"
)

func QueryOneTest(id int64) (*data.TestData) {
	sql := `select * from test where id = ` + strconv.FormatInt(id, 10)
	result, err := Query(sql)
	if (err == false) || (len(*result) == 0) {
		return nil
	}

	bird := formatTestData(&(*result)[0])
	return bird
}

func formatTestData(result *interface{}) (*data.TestData) {
	r := (*result).(map[string]string)
	_id, _ := r["id"]
	desc, _ := r["desc"]

	id, _ := strconv.ParseInt(_id, 10, 64)
	testData := data.TestData{
		Id:   id,
		Desc: desc,
	}
	return &testData
}
