package method

import (
	"fmt"
	"time"
)

// time.DateTime = "2006-01-02 15:04:05"
// time.DateOnly = "2006-01-02"
// time.TimeOnly = "15:04:05"
type serverTimeDate struct {
	Year      int
	Month     int
	Day       int
	Hour      int
	Minute    int
	Second    int
	Timestamp int64
}
type serverTime struct {
	Now     string
	NowDay  string
	NowTime string
	Date    serverTimeDate
}

func Time() *serverTime {
	now := time.Now()
	year := now.Year()
	month := int(now.Month())
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()
	timestamp := now.Unix()
	nowDay := fmt.Sprintf("%d-%d-%d", year, month, day)
	nowTime := fmt.Sprintf("%d:%d:%d", hour, minute, second)
	
	return &serverTime{
		Now:     nowDay + " " + nowTime,
		NowDay:  nowDay,
		NowTime: nowTime,
		Date: serverTimeDate{
			Year:      year,
			Month:     month,
			Day:       day,
			Hour:      hour,
			Minute:    minute,
			Second:    second,
			Timestamp: timestamp,
		},
	}
}
