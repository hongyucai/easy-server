package helper

import (
	"fmt"
	"github.com/kataras/golog"
	"strconv"
	"time"
)

var MonthMap map[string]string = map[string]string{
	"January":   "01",
	"February":  "02",
	"March":     "03",
	"April":     "04",
	"May":       "05",
	"June":      "06",
	"July":      "07",
	"August":    "08",
	"September": "09",
	"October":   "10",
	"November":  "11",
	"December":  "12",
}

var MonthIntMap map[string]int = map[string]int{
	"January":   1,
	"February":  2,
	"March":     3,
	"April":     4,
	"May":       5,
	"June":      6,
	"July":      7,
	"August":    8,
	"September": 9,
	"October":   10,
	"November":  11,
	"December":  12,
}

func FailOnError(err error, msg string) {
	if err != nil {
		golog.Errorf("%s: %s", msg, err.Error())
		panic(fmt.Sprintf("%s: %s", msg,err.Error()))
	}
}

// 获取月初时间
func GetEarlyMonthUnix() int64 {
	now := time.Now()
	tm := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
	return tm.Unix()
}

// 获取零时时间
func GetZeroHourUnix() int64 {
	now := time.Now()
	tm := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	return tm.Unix()
}

// 获取当前小时时间
func GetNowHourUnix() int64 {
	now := time.Now()
	tm := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), 0, 0, 0, now.Location())
	return tm.Unix()
}

// 获取当前时间时间戳 秒
func GetNowUnix() int64 {
	return time.Now().Unix()
}
// 获取当前时间时间戳 毫秒
func GetNowUnixHano() int64 {
	return time.Now().UnixNano() / 1e6
}
func GetNowTime() string {
	now := time.Now()
	return fmt.Sprintf("%d-%d-%d %d:%d:%d",now.Year(),now.Month(),now.Day(),now.Hour(),now.Minute(),now.Second())
}

func Timestamp2str(timestamp int64,format string) string {
	return time.Unix(timestamp, 0).Format(format)
}
//日期转化为时间戳
func Datetime2timestamp(datetime string) int64{
	//datetime = "2019-12-06 00:00:00"  //待转化为时间戳的字符串
	tmp, _ := time.ParseInLocation(SysTimeform, datetime, SysTimeLocation)
	return tmp.Unix()
}
//时间戳转化为日期
func timestamp2datetime(timestamp int64) (datetime string) {
	return time.Unix(timestamp, 0).Format(SysTimeform)
}
// 获取年初时间
func GetEarlyYearUnix() int64 {
	now := time.Now()
	tm := time.Date(now.Year(), 1, 1, 0, 0, 0, 0, now.Location())
	return tm.Unix()
}

func GetUnixToFormatString(timestamp int64, f string) string {
	tm := time.Unix(timestamp, 0)
	return tm.Format(f)
}

func GetUnixToString(timestamp int64) string {
	return GetUnixToFormatString(timestamp, "2006-01-02 00:00:00")   //200601021504
}

func GetUnixToHourString(timestamp int64) string {
	return GetUnixToFormatString(timestamp, "15:04")
}

func GetUnixToMonth(timestamp int64) string {
	tm := time.Unix(timestamp, 0)
	return MonthMap[tm.Month().String()]
}

func GetUnixToDay(timestamp int64) int {
	tm := time.Unix(timestamp, 0)
	return tm.Day()
}

func GetUnixToDayTime(timestamp int64) string {
	month := GetUnixToMonth(timestamp)
	day := GetUnixToDay(timestamp)
	d := month + "." + strconv.Itoa(day)
	return d
}

func GetUnixToOldTime(i int) int64 {
	currentMonth := MonthIntMap[time.Now().Month().String()]

	oldMonth := currentMonth - i
	t := time.Date(time.Now().Year(),time.Month(oldMonth),1,0,0,0,0,time.Local)
	return t.Unix()
}

func GetUnixToOldYearTime(i int) int64 {
	currentYear := time.Now().Year()
	oldMonth := currentYear - i

	t := time.Date(oldMonth, 1, 1, 0, 0, 0, 0, time.Local)
	return t.Unix()
}

func GetUnixToOldTimeDay(i int) int64 {
	day := time.Now().Day()
	oldMonth := day - i
	t := time.Date(time.Now().Year(), time.Now().Month(), oldMonth, time.Now().Hour(), time.Now().Minute(), time.Now().Second(), time.Now().Nanosecond(), time.Local)
	return t.Unix()
}