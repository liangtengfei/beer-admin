package datetime

import "time"

const (
	YYYYMMDDHHMMSS  = "2006-01-02 15:04:05"
	YYYYMMDDHHMMSS4 = "20060102150405"
	YYYYMMDD        = "2006-01-02"
	YYYYMMDD4       = "20060102"
	HHMMSS          = "15:04:05"
	HHMMSS4         = "150405"
)

// Full eg: 2006-01-02 15:04:05
func Full(t ...time.Time) string {
	if len(t) > 0 {
		return t[0].Format(YYYYMMDDHHMMSS)
	}
	return time.Now().Format(YYYYMMDDHHMMSS)
}
func Full4Path(t ...time.Time) string {
	if len(t) > 0 {
		return t[0].Format(YYYYMMDDHHMMSS4)
	}
	return time.Now().Format(YYYYMMDDHHMMSS4)
}

func FullDate(t ...time.Time) string {
	if len(t) > 0 {
		return t[0].Format(YYYYMMDD)
	}
	return time.Now().Format(YYYYMMDD)
}
func FullDate4Path(t ...time.Time) string {
	if len(t) > 0 {
		return t[0].Format(YYYYMMDD4)
	}
	return time.Now().Format(YYYYMMDD4)
}

func FullTime(t ...time.Time) string {
	if len(t) > 0 {
		return t[0].Format(HHMMSS)
	}
	return time.Now().Format(HHMMSS)
}
func FullTime4Path(t ...time.Time) string {
	if len(t) > 0 {
		return t[0].Format(HHMMSS4)
	}
	return time.Now().Format(HHMMSS4)
}
