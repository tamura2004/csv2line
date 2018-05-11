package main

import (
	"io"
)

func writeRow(w io.Writer, keys []TagFieldKey, row []string) {

	t := toTimestamp(row[0])

	// データを出力
	for j, col := range row[1:] {
		if col == " " || col == "0" {
			continue
		}
		writeCol(w, keys[j], col, t)
	}
}
