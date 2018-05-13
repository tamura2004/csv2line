package main

import (
	"time"
)

func buildTimestamp(s string) int64 {
	// タイムスタンプをナノ秒で取得
	t, err := time.Parse("01/02/2006 15:04:05.000 -0700", s+" +0900")
	handleError(err)

	return t.UnixNano()
}