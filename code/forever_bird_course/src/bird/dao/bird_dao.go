package dao

import (
	"bird/data"
	"bird/util"
	"strconv"
)

func QueryBirdByTokenId(tokenId int64) *data.Bird {
	sql := `select * from bird where token_id = ` + strconv.FormatInt(tokenId, 10)
	result, err := Query(sql)
	if (err == false) || (len(*result) == 0) {
		return nil
	}

	bird := formatBirdInfo(&(*result)[0])
	return bird
}

func formatSql(tokenIds *[]int64) string {
	s := ""
	for _, tokenId := range *tokenIds {
		t := strconv.FormatInt(tokenId, 10)
		s += t
		s += ","
	}
	l := len(s)
	return s[0:l-1]
}

/**
 ** 查询birds
 */
func QueryBirdByTokenIds(tokenIds []int64) (*[]data.Bird, bool) {
	sql := `select * from bird where token_id in(` + formatSql(&tokenIds) + `)`

	rows, err := Query(sql)
	if err == false {
		return nil, false
	}

	size := len(*rows)
	if size == 0 {
		return nil, true
	}

	var record []data.Bird
	for _, row := range *rows {
		if row != nil {
			bird := formatBirdInfo(&row)
			record = append(record, *bird)
		}
	}
	return &record, true
}

/**
 ** 查询birds
 */
func QueryBirdsMapByTokenIds(tokenIds []int64) (*map[int64]*data.Bird, bool) {
	sql := `select * from bird where token_id in(` + formatSql(&tokenIds) + `)`
	util.Logger.Info("QueryBirdsMapByTokenIds sql: ", sql)

	rows, err := Query(sql)
	if err == false {
		return nil, false
	}

	size := len(*rows)
	if size == 0 {
		return nil, true
	}

	var record  = make(map[int64]*data.Bird)
	for _, row := range *rows {
		if row != nil {
			bird := formatBirdInfo(&row)
			record[bird.TokenId] = bird
		}
	}
	return &record, true
}

func QueryBirdByAddressAndTokenId(address string, tokenId int64) (*data.Bird) {
	sql := "select * from bird where address = \"" + address + "\" and token_id=" + strconv.FormatInt(tokenId, 10)
	result, err := Query(sql)
	if (err == false) || (len(*result) == 0) {
		return nil
	}

	bird := formatBirdInfo(&(*result)[0])
	return bird
}

func QueryBirdsCount() (int, bool) {
	sql := `select count(1) from bird`
	rows, err := Query(sql)
	if err == false {
		return 0, false
	}

	r := ((*rows)[0]).(map[string]string)
	count, _ := r["count(1)"]
	c, _ := strconv.Atoi(count)

	return c, true
}

func QueryPKListBirdsCount(address string) (int, bool) {
	sql := "select count(1) from bird where `address` <> \"" + address + "\" and `status` = " + strconv.Itoa(data.Bird_Status_Normal)
	rows, err := Query(sql)
	if err == false {
		return 0, false
	}

	r := ((*rows)[0]).(map[string]string)
	count, _ := r["count(1)"]
	c, _ := strconv.Atoi(count)

	return c, true
}

func UpdateBirdStatus(tokenId int64, status int) (bool){
	sql := "update bird set `status`=? where `token_id`=?"
	args := []interface{}{status, tokenId}

	count := UpdateWithArgs(sql, args)
	return count > 0
}

func UpdateBirdGenerateInfo(status int, bird *data.Bird) (bool) {
	sql := "update bird set `status`=?, `name`=?, `rarity`=?, `svg_path`=? where `token_id`=?"
	args := []interface{}{status, bird.Name, bird.Rarity, bird.SvgPath, bird.TokenId}

	count := UpdateWithArgs(sql, args)
	return count > 0
}

func UpdateBirdSvg(tokenId int64, status int, localSvg string, svgPath string) (bool){
	sql := "update bird set `status`=?, `local_svg_path`=?, `svg_path`=? where `token_id`=?"
	args := []interface{}{status, localSvg, svgPath, tokenId}

	count := UpdateWithArgs(sql, args)
	return count > 0
}

func QueryPKListBirds(address string, offset int, count int) (*[]data.Bird, bool) {
	sql := "select * from bird where `address` <> \"" + address + "\" and `status` = " + strconv.Itoa(data.Bird_Status_Normal) + "  order by `weight` asc limit " + strconv.Itoa(offset) + `,` + strconv.Itoa(count)
	rows, err := Query(sql)
	if err == false {
		return nil, false
	}

	size := len(*rows)
	if size == 0 {
		return nil, true
	}

	var record []data.Bird
	for _, row := range *rows {
		if row != nil {
			bird := formatBirdInfo(&row)
			record = append(record, *bird)
		}
	}
	return &record, true
}

func QueryBirdsByRank(offset int, count int) (*[]data.RankBird, bool) {
	sql := "select (@i:=@i+1) as `rowno`,bird.* from bird,(select @i:=" + strconv.Itoa(offset) + ") b order by `weight` desc, `token_id` asc " + " limit " + strconv.Itoa(offset) + `,` + strconv.Itoa(count)
	rows, err := Query(sql)
	if err == false {
		return nil, false
	}

	size := len(*rows)
	if size == 0 {
		return nil, true
	}

	var record []data.RankBird
	for _, row := range *rows {
		if row != nil {
			bird := formatRankBirdInfo(&row)
			record = append(record, *bird)
		}
	}
	return &record, true
}

func QueryBirdsByAddress(address string) (*[]data.Bird, bool) {
	sql := `select * from bird where address = "` + address + `"`
	rows, err := Query(sql)
	if err == false {
		return nil, false
	}

	size := len(*rows)
	if size == 0 {
		return nil, true
	}

	var record []data.Bird
	for _, row := range *rows {
		if row != nil {
			bird := formatBirdInfo(&row)
			record = append(record, *bird)
		}
	}
	return &record, true
}

func QueryBirdsByAddressWithLimit(address string, offset int, count int) (*[]data.Bird, bool) {
	sql := `select * from bird where address = "` + address + `"` + " limit " + strconv.Itoa(offset) + `,` + strconv.Itoa(count)
	rows, err := Query(sql)
	if err == false {
		return nil, false
	}

	size := len(*rows)
	if size == 0 {
		return nil, true
	}

	var record []data.Bird
	for _, row := range *rows {
		if row != nil {
			bird := formatBirdInfo(&row)
			record = append(record, *bird)
		}
	}
	return &record, true
}


func QueryUserNormalBirdsCount(address string) (int, bool) {
	sql := "select count(1) from bird where `status` = 2 and `address` = \"" + address + "\""
	rows, err := Query(sql)
	if err == false {
		return 0, false
	}

	r := ((*rows)[0]).(map[string]string)
	count, _ := r["count(1)"]
	c, _ := strconv.Atoi(count)

	return c, true
}

func QueryNormalBirdsByAddressWithLimit(address string, offset int, count int) (*[]data.Bird, bool) {
	sql :=  "select * from bird where `status` = 2 and `address` = \"" + address + "\" limit " + strconv.Itoa(offset) + "," + strconv.Itoa(count)
	rows, err := Query(sql)
	if err == false {
		return nil, false
	}

	size := len(*rows)
	if size == 0 {
		return nil, true
	}

	var record []data.Bird
	for _, row := range *rows {
		if row != nil {
			bird := formatBirdInfo(&row)
			record = append(record, *bird)
		}
	}
	return &record, true
}

func formatBirdInfo(result *interface{}) (*data.Bird) {
	r := (*result).(map[string]string)
	id, _ := r["token_id"]
	address, _ := r["address"]
	status, _ := r["status"]
	name, _ := r["name"]
	desc, _ := r["desc"]
	birth, _ := r["birthday"]
	rarity, _ := r["rarity"]
	speed, _ := r["speed"]
	power, _ := r["power"]
	level, _ := r["level"]
	weight, _ := r["weight"]
	experience, _ := r["experience"]
	genome, _ := r["genes"]
	svgPath, _ := r["svg_path"]
	eatFruitTime, _ := r["eat_fruit_time"]

	tokenId, _ := strconv.ParseInt(id, 10, 64)
	birthday, _ := strconv.Atoi(birth)
	sp, _ := strconv.Atoi(speed)
	p, _ := strconv.Atoi(power)
	l, _ := strconv.Atoi(level)
	w, _ := strconv.ParseInt(weight,10, 64)
	exp, _ := strconv.ParseInt(experience,10, 64)
	s, _ := strconv.Atoi(status)
	eat, _ := strconv.ParseInt(eatFruitTime, 10, 64)

	bird := data.Bird{
		TokenId:tokenId,
		Name:name,
		Desc:desc,
		Birthday:uint(birthday),
		Rarity:rarity,
		Speed:sp,
		Power:p,
		Level:l,
		Weight:w,
		Experience:exp,
		Owner:address,
		Status:s,
		Genes:genome,
		SvgPath:svgPath,
		EatFruitTime:eat,
	}
	return &bird
}

func formatRankBirdInfo(result *interface{}) (*data.RankBird) {
	r := (*result).(map[string]string)
	rowno, _ := r["rowno"]
	id, _ := r["token_id"]
	address, _ := r["address"]
	status, _ := r["status"]
	name, _ := r["name"]
	desc, _ := r["desc"]
	birth, _ := r["birthday"]
	rarity, _ := r["rarity"]
	speed, _ := r["speed"]
	power, _ := r["power"]
	level, _ := r["level"]
	weight, _ := r["weight"]
	experience, _ := r["experience"]
	genome, _ := r["genes"]
	svgPath, _ := r["svg_path"]

	row_no, _ := strconv.Atoi(rowno)
	tokenId, _ := strconv.ParseInt(id, 10, 64)
	birthday, _ := strconv.Atoi(birth)
	sp, _ := strconv.Atoi(speed)
	p, _ := strconv.Atoi(power)
	l, _ := strconv.Atoi(level)
	w, _ := strconv.ParseInt(weight,10, 64)
	exp, _ := strconv.ParseInt(experience,10, 64)
	s, _ := strconv.Atoi(status)

	bird := data.RankBird{
		RowNo:row_no,
		TokenId:tokenId,
		Name:name,
		Desc:desc,
		Birthday:uint(birthday),
		Rarity:rarity,
		Speed:sp,
		Power:p,
		Level:l,
		Weight:w,
		Experience:exp,
		Owner:address,
		Status:s,
		Genes:genome,
		SvgPath:svgPath,
	}
	return &bird
}