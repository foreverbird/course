package dao

import (
	"database/sql"
	"log"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"bird/util"
	"time"
)

var gDb *sql.DB

func init() {
	var err error
	dataSource := "bird:b3069797e91a31f7@tcp(127.0.0.1:3306)/bird?charset=utf8"
	gDb, err = sql.Open("mysql", dataSource)
	if err != nil {
		util.Logger.Error("open db failed: " + err.Error())
		log.Fatal("open db failed: " + err.Error())
	}

	// 设置最大打开连接数
	gDb.SetMaxOpenConns(2000)
	// 设置闲置连接数
	gDb.SetMaxIdleConns(1000)
	// 设置连接池的最长连接时间
	gDb.SetConnMaxLifetime(2 * time.Hour)

	gDb.Ping()
}

/**
 ** 数据库查询接口
 */
func Query(sql string) (*[]interface{}, bool) {
	fmt.Println(sql)
	rows, err := gDb.Query(sql)
	defer rows.Close()

	if err != nil {
		//查询数据库失败
		util.Logger.Error("query failed: ", err)
		return nil, false
	}

	columns, _ := rows.Columns()
	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))
	for j := range values {
		scanArgs[j] = &values[j]
	}

	var record []interface{}
	for rows.Next() {
		//将行数据保存到record字典
		err = rows.Scan(scanArgs...)
		row := make(map[string]string)
		for i, col := range values {
			if col != nil {
				row[columns[i]] = string(col.([]byte))
			}
		}
		record = append(record, row)
	}

	return &record, true
}

/**
 ** 插入数据
 */
func Insert(sql string) (id int64) {
	result, err := gDb.Exec(sql)
	if err != nil {
		return 0
	}

	// 获取插入的id
	id, err = result.LastInsertId()
	if err != nil {
		return 0
	}
	return id
}


/**
 ** 插入数据
 */
func InsertWithArgs(sql string, args []interface{}) (id int64) {
	s, err := gDb.Prepare(sql)
	if err != nil {
		fmt.Println(sql)
		fmt.Println(err)
		return -1
	}

	result, err := s.Exec(args...)
	if err != nil {
		util.Logger.Error("insert db failed", sql, err)
		return -1
	}

	id, err = result.LastInsertId()
	if err != nil {
		return 0
	}
	return id
}


func exec(sql string) (count int64) {
	result, err := gDb.Exec(sql)
	if err != nil {
		return 0
	}

	count, err = result.RowsAffected()
	if (err != nil) || (count == 0) {
		return 0
	}
	return count
}

func execWithArgs(sql string, args []interface{}) (count int64) {
	s, err := gDb.Prepare(sql)
	if err != nil {
		return 0
	}

	result, err := s.Exec(args...)
	if err != nil {
		return 0
	}

	count, err = result.RowsAffected()
	if (err != nil) || (count == 0) {
		return 0
	}
	return count
}

/**
 ** 更新数据
 */
func Update(sql string) (count int64) {
	return exec(sql)
}

/**
 ** 更新数据
 */
func UpdateWithArgs(sql string, args []interface{}) (count int64) {
	return execWithArgs(sql, args)
}

/**
 ** 删除数据
 */
func Delete(sql string) (count int64) {
	return exec(sql)
}
