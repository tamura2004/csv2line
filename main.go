package main

import (
	"log"
	"os"
	"path/filepath"
)

type TagFieldKey struct {
	hostname string
	category string
	instance string
	index    string
	label    string
}

func init() {
	if len(os.Args) < 2 {
		log.Fatal("No file name")
	}
}

func main() {
	// パラメータをファイル名に展開
	inFileNames, err := filepath.Glob(os.Args[1])
	handleError(err)

	for _, inFileName := range inFileNames {
		handleCSV(inFileName)
	}
}
