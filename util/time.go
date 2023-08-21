package util

import "time"

// DayToMilli expiration days to expiration time (13-digit timestamp)
func DayToMilli(day uint) uint {
	return NowMilli() + day*86400*1000
}

// NowMilli Get the 13-digit timestamp of the current time
func NowMilli() uint {
	return uint(time.Now().UnixMilli())
}
