package dao

import (
	"strconv"
	"bird/data"
)

func QueryFruits() (*[]data.Fruit, bool) {
	sql := `select * from fruits`
	rows, err := Query(sql)
	if err == false {
		return nil, false
	}

	size := len(*rows)
	if size == 0 {
		return nil, true
	}

	var record []data.Fruit
	for _, row := range *rows {
		if row != nil {
			fruit, err := formatFruitInfo(&row)
			if err {
				record = append(record, *fruit)
			}
		}
	}
	return &record, true
}

func QueryFruitByType(t int) (*data.Fruit, bool) {
	sql := "select * from fruits where `type` =" + strconv.Itoa(t)
	rows, err := Query(sql)
	if err == false {
		return nil, false
	}
	size := len(*rows)
	if size == 0 {
		return nil, true
	}

	fruit, err := formatFruitInfo(&(*rows)[0])

	return fruit, true
}

func formatFruitInfo(result *interface{}) (*data.Fruit, bool) {
	r := (*result).(map[string]string)
	id, e0 := r["id"]
	t, e1 := r["type"]
	desc, e2 := r["desc"]
	price, e3 := r["price"]
	s, e4 := r["strength"]

	if e0 && e1 && e2 && e3 && e4 {
		fruitId, _ := strconv.Atoi(id)
		ty, _ := strconv.Atoi(t)
		strength, _ := strconv.Atoi(s)

		fruit := data.Fruit{
			Id:  fruitId,
			Type:     ty,
			Desc:     desc,
			Price:    price,
			Strength: strength,
		}
		return &fruit, true
	}

	return nil, false
}
