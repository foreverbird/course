package util

import (
	"time"
	"fmt"
)

/**
 * 获取当前星期的起始时间
 */
func CountCurWeekStartingTime() (int) {
	// 当前时间
	now := time.Now().Local()
	s := now.Format("2006-01-02 15:04:05")

	fmt.Println("now time:", s)

	// 去掉时间
	s = s[0:10]
	s = s + " 00:00:00"
	p, _ := time.ParseInLocation("2006-01-02 15:04:05", s, time.Local)

	// 时间间隔
	var deltaDay int
	weekDay := now.Weekday()
	switch weekDay {
	case time.Sunday:
		deltaDay = 6
	case time.Monday:
		deltaDay = 0
	case time.Tuesday:
		deltaDay = 1
	case time.Wednesday:
		deltaDay = 2
	case time.Thursday:
		deltaDay = 3
	case time.Friday:
		deltaDay = 4
	case time.Saturday:
		deltaDay = 5
	}

	fmt.Println("curweek starting time:", int(p.Unix()) - deltaDay * 24 * 3600)

	return int(p.Unix()) - deltaDay * 24 * 3600 + 1
}
