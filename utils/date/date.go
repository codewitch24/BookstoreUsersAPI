package date

import "time"

const (
	ApiDateLayout = "2006-01-02T15:04:05.000Z"
)

func GetNow() time.Time {
	return time.Now().UTC()
}

func GetNowStringDate() string {
	return GetNow().Format(ApiDateLayout)
}
