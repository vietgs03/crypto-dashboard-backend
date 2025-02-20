package utils

import (
	"time"

	"crypto-dashboard/pkg/constants"
	"crypto-dashboard/pkg/response"

	"github.com/samber/lo"
)

func CurrentDate() string {
	return time.Now().Format(time.DateOnly)
}

func CurrentTimeISO() string {
	return time.Now().Format(time.RFC3339)
}

func PreviousDate() string {
	return time.Now().Add(-24 * time.Hour).Format(time.DateOnly)
}

func TimeToString(t time.Time, frm string) string {
	return t.Format(frm)
}

func StringToTime(dateStr, frm string) (*time.Time, *response.AppError) {
	parsedTime, err := time.Parse(string(frm), dateStr)
	if err != nil {
		return nil, response.NewAppError(constants.TimeInvalid, "invalid time format")
	}

	return &parsedTime, nil
}

func StringToUnix(dateStr, frm string) (int64, *response.AppError) {
	t, err := StringToTime(dateStr, frm)
	if err != nil {
		return 0, err
	}

	return t.Unix(), nil
}

func TimePointerToString(t *time.Time, format string) string {
	if t == nil {
		return ""
	}

	return t.Format(format)
}

func UnixToString(unixTime int64, frm string) string {
	timestamp := time.Unix(unixTime, 0)

	return timestamp.Format(frm)
}

func LastDayOfThisMonth() int {
	now := time.Now()
	year, month := now.Year(), now.Month()

	firstDayOfNextMonth := time.Date(year, month+1, 1, 0, 0, 0, 0, now.Location())

	return firstDayOfNextMonth.Add(-24 * time.Hour).Day()
}

func LastDayOfNextMonth() int {
	nextMonth := time.Now().AddDate(0, 1, 0)
	year, month := nextMonth.Year(), nextMonth.Month()

	firstDayOfNextMonth := time.Date(year, month+1, 1, 0, 0, 0, 0, nextMonth.Location())

	return firstDayOfNextMonth.Add(-24 * time.Hour).Day()
}

func ThisMonth() int {
	return int(time.Now().Month())
}

func NextMonth() int {
	nextMonth := int(time.Now().Month())%12 + 1

	return nextMonth
}

func FormatToEndOfDay(t string) (*string, *response.AppError) {
	d, err := StringToTime(t, time.RFC3339)
	if err != nil {
		return nil, err
	}

	res := d.Add(23*time.Hour + 59*time.Minute).Format("200601021504")

	return &res, nil
}

func StringToDate(str string) string {
	d, err := StringToTime(str, time.RFC3339)
	if err != nil {
		return str
	}

	// Custom date format layout
	const customDateFormat = "2006/01/02"

	res := lo.FromPtr(d).Format(customDateFormat)
	return res
}

func StringToDateWithFormat(str, format string) string {
	d, err := StringToTime(str, time.RFC3339)
	if err != nil {
		return str
	}

	res := lo.FromPtr(d).Format(format)
	return res
}

func FirstDayOfMonth(str string, frm string) (*time.Time, *response.AppError) {
	d, err := StringToTime(str, frm)
	if err != nil {
		return nil, err
	}

	firstDateOfMonth := time.Date(d.Year(), d.Month(), 1, 0, 0, 0, 0, d.Location())

	return &firstDateOfMonth, nil
}

func LastDayOfMonth(str string, frm string) (*time.Time, *response.AppError) {
	d, err := StringToTime(str, frm)
	if err != nil {
		return nil, err
	}

	firstDateOfNextMonth := time.Date(d.Year(), d.Month(), 1, 0, 0, 0, 0, d.Location())
	lastDateOfMonth := firstDateOfNextMonth.AddDate(0, 1, -1)

	return &lastDateOfMonth, nil
}

func RangeDate(dateFrom string, dateTo string, fmt string) (int, *response.AppError) {
	dateFromTime, err := StringToTime(dateFrom, fmt)
	if err != nil {
		return 0, err
	}

	dateToTime, err := StringToTime(dateTo, fmt)
	if err != nil {
		return 0, err
	}

	return int(dateToTime.Sub(*dateFromTime).Hours() / 24), nil
}

func EndOfDay(t time.Time) time.Time {
	// Set the time to 23:59:59.999999999
	return time.Date(
		t.Year(), t.Month(), t.Day(),
		23, 59, 59, 0,
		t.Location(),
	)
}

func BeginningOfDay(t time.Time) time.Time {
	return time.Date(
		t.Year(), t.Month(), t.Day(),
		0, 0, 0, 0,
		t.Location(),
	)
}

func MoreThan24Hours(t time.Time) bool {
	now := time.Now()

	return now.Sub(t).Hours() > 24
}

// CompareDateStrings compares two date strings and returns:
// -2 is error
//
// -1 if date1 is before date2
//	0 if date1 is equal to date2
//	1 if date1 is after date2

func CompareDateStrings(date1, date2, layout string) int {
	parsedDate1, err := time.Parse(layout, date1)
	if err != nil {
		return -2
	}

	parsedDate2, err := time.Parse(layout, date2)
	if err != nil {
		return -2
	}

	switch {
	case parsedDate1.Before(parsedDate2):
		return -1
	case parsedDate1.After(parsedDate2):
		return 1
	default:
		return 0
	}
}
