package utils

import (
	"fmt"
	"strconv"
	"time"
)

func GetCurrentTimestampInt() int64 {
	return time.Now().Unix()
}

func GetCurrentTimestampString() string {
	seconds := time.Now().Unix()
	return strconv.FormatInt(seconds, 10)
}

type TimeDetail struct {
	Year   string `json:"year"`
	Month  string `json:"month"`
	Date   string `json:"date"`
	Hour   string `json:"hour"`
	Minute string `json:"minute"`
	Second string `json:"second"`
}

const (
	noYear = "01-02"
)

// Time2Date 2006-01-02
func Time2Date(timeStr time.Time) string {
	return timeStr.Format(time.DateOnly)
}

// Time2FullDateTime 2006-01-02 15:04:05
func Time2FullDateTime(timeStr time.Time) string {
	return timeStr.Format(time.DateTime)
}

func Time2DateMonthOnly(timeStr time.Time) string {
	return timeStr.Format(noYear)
}

func GetMessageId() string {
	return fmt.Sprintf("%d", time.Now().Nanosecond())
}

func GetDayTime() int64 {
	t := time.Now()
	newTime := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	return newTime.Unix()
}

func GetDate() string {
	strDate := time.Unix(time.Now().Unix(), 0).Format(time.DateOnly)
	return strDate
}

func GetTomorrowDate() string {
	tomorrow := time.Now().AddDate(0, 0, 1)
	return tomorrow.Format(time.DateOnly)
}

func GetDateOnly() string {
	return time.Now().Format(time.DateOnly)
}

func GetDateAndHour() string {
	strDate := time.Unix(time.Now().Unix(), 0).Format("2006010215")
	return strDate
}

func GetDateByTime(someTime int64) string {
	strDate := time.Unix(someTime, 0).Format("20060102")
	return strDate
}

func GetDateDetailString() string {
	current := time.Now()
	date := current.Format("2006-01-02-15-04-05")
	return date
}

func Date2Time(stringTime string) (time.Time, error) {
	return time.ParseInLocation(time.DateOnly, stringTime, time.Local)
}

func GetHourByString(stringTime string) (error, time.Time) {
	stamp, err := time.ParseInLocation("2006010215", stringTime, time.Local)
	return err, stamp
}

func GetTimeDetail(someTime int64) (detail TimeDetail) {
	strYear := time.Unix(someTime, 0).Format("2006")
	strMonth := time.Unix(someTime, 0).Format("200601")
	strDate := time.Unix(someTime, 0).Format("20060102")
	strHour := time.Unix(someTime, 0).Format("2006010215")
	strMinute := time.Unix(someTime, 0).Format("200601021504")
	strSecond := time.Unix(someTime, 0).Format("20060102150405")

	detail = TimeDetail{
		Year:   strYear,
		Month:  strMonth,
		Date:   strDate,
		Hour:   strHour,
		Minute: strMinute,
		Second: strSecond,
	}
	return
}

// GetFirstDateOfWeek 计算本周周一的日期
func GetFirstDateOfWeek(nowTime time.Time) time.Time {
	now := nowTime

	offset := int(time.Monday - now.Weekday())
	if offset > 0 {
		offset = -6
	}
	weekStartDate := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset)
	return weekStartDate
}

func TimeCompare(timeBefore, timeAfter string) bool {
	//time1 := "2015-03-20 08:50:29"
	//time2 := "2015-03-21 09:04:25"
	//先把时间字符串格式化成相同的时间类型
	t1, err := time.Parse(time.DateOnly, timeBefore)
	t2, err := time.Parse(time.DateOnly, timeAfter)
	if err == nil && t1.Before(t2) {
		return true
	}
	return false
}

// CalcAge Calculate birthday based on date, date format: yyyy MM DD
func CalcAge(birthdayStr string) (int32, error) {
	birthDate, err := time.Parse(time.DateOnly, birthdayStr)
	if err != nil {
		return 0, err
	}
	currentTime := time.Now()
	age := int32(currentTime.Year() - birthDate.Year())
	if currentTime.Month() < birthDate.Month() || (currentTime.Month() == birthDate.Month() && currentTime.Day() < birthDate.Day()) {
		age--
	}
	return age, nil
}

// CalcAgeWithoutErr Calculate birthday based on date, date format: yyyy-MM-DD, return without err
func CalcAgeWithoutErr(birthdayStr string) int32 {
	birthDate, err := time.Parse(time.DateOnly, birthdayStr)
	if err != nil {
		return -1
	}
	currentTime := time.Now()
	age := int32(currentTime.Year() - birthDate.Year())
	if currentTime.Month() < birthDate.Month() || (currentTime.Month() == birthDate.Month() && currentTime.Day() < birthDate.Day()) {
		age--
	}
	return age
}

// GetCurrentDateTime Get the 'yyyy-MM-DD HH:MM:SS' format of the current time.
func GetCurrentDateTime() string {
	return time.Now().Format(time.DateTime)
}

// GetMinDateTime Get the minimum date formats for the current time's day
func GetMinDateTime() string {
	now := time.Now()
	startOfDay := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	// 将time.Time类型转换为指定格式的string类型
	return startOfDay.Format(time.DateTime)
}

// GetMaxDateTime Retrieve the maximum date formats for the current time's day
func GetMaxDateTime() string {
	now := time.Now()
	startOfDay := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, now.Location())
	// 将time.Time类型转换为指定格式的string类型
	return startOfDay.Format(time.DateTime)
}

// GetDayNumber get the specified date as the day of the current year
func GetDayNumber(t time.Time) int {
	return t.YearDay()
}

// GetWeekNumber get the specified week of the current year for the specified time
func GetWeekNumber(t time.Time) int {
	_, week := t.ISOWeek()
	return week
}

func CalcDaysBetweenDate(s, e string) (int, error) {
	t1, err := time.Parse(time.DateOnly, s)
	if err != nil {
		return 0, err
	}
	t2, err := time.Parse(time.DateOnly, e)
	if err != nil {
		return 0, err
	}
	// Calculate the difference in days between two dates
	duration := t2.Sub(t1)

	return int(duration.Hours() / 24), nil
}

func CalcMillisecondsBetween(t1, t2 time.Time) int64 {
	return t2.Sub(t1).Milliseconds()
}
