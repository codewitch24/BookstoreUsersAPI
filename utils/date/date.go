package date

import "time"

const (
	apiDateLayout     = "2006-01-02T15:04:05.000Z"
	apiDatabaseLayout = "2006-01-02 15:04:05"
)

func GetNow() time.Time {
	return time.Now().UTC()
}

func GetNowStringDate() string {
	return GetNow().Format(apiDateLayout)
}

func GetNowDatabaseFormat() string {
	return GetNow().Format(apiDatabaseLayout)
}
