package main

import (
	"flag"
	"path/filepath"
)

type TagFieldKey struct {
	hostname string
	category string
	instance string
	index    string
	label    string
}

var (
	InFileName string
	database   string
)

func init() {
	flag.StringVar(&database, "d", "perf", "Name of database to use")
	flag.Parse()
	InFileName = flag.Arg(0)
}

func main() {
	// パラメータをファイル名に展開
	inFileNames, err := filepath.Glob(InFileName)
	handleError(err)

	for _, inFileName := range inFileNames {
		handleCSV(inFileName)
	}
}
