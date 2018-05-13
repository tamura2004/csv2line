package main

import (
	"os"
)

func handleCSV(inFileName string) {
	outFileName, err := buildFilename(inFileName)
	handleError(err)

	// 出力ファイルをオープン
	outFile, err := os.Create(outFileName)
	handleError(err)
	defer outFile.Close()

	// lineファイルヘッダを出力
	writeHeader(outFile)

	// 入力ファイルをオープン
	inFile, err := os.Open(inFileName)
	handleError(err)
	defer inFile.Close()

	var tagFieldKeys []TagFieldKey
	csvReadWithIndex(inFile, func(row []string, i int) {
		if i == 0 {
			// カラムヘッダをパース
			tagFieldKeys = parseHeaderRow(row)
		} else {
			// 行を出力
			writeRow(outFile, tagFieldKeys, row)
		}
	})

}
