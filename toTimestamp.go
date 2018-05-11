package main

import (
	"time"
)

func toTimestamp(s string) int64 {
	// タイムスタンプをナノ秒で取得
	t, err := time.Parse("01/02/2006 15:04:05.000", s)
	handleError(err)
	return t.UnixNano()
}
