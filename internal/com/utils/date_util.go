package utils
import "time"

const (
	TIME_FORMAT = "2006-01-02 15:04:05"
	TIME_FORMAT_YYYYMMDD = "20060102"
)

/**
 * 将日期(字符串)转为时间戳
 */
func DateToTimestamp(dateStr string) int64 {
	loc, _ := time.LoadLocation("Local")
	theTime, _ := time.ParseInLocation(TIME_FORMAT, dateStr, loc)
	return theTime.Unix()
}

func CurDateTime() string {
	return time.Now().Format(TIME_FORMAT)
}