package utils

import (
	"fmt"
	"time"
)

func GetDayTimeNow() string {
	loc, _ := time.LoadLocation("Asia/Bangkok") // GMT+7
	now := time.Now().In(loc)
	return fmt.Sprintf("%s (%02d/%02d/%04d) %02d:%02d:%02d",
		now.Weekday().String(),
		now.Day(), now.Month(), now.Year(),
		now.Hour(), now.Minute(), now.Second())
}

func GetHourNow() string {
	loc, _ := time.LoadLocation("Asia/Bangkok") // GMT+7
	now := time.Now().In(loc)
	return fmt.Sprintf("%d", now.Hour())
}
