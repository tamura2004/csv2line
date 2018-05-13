package main

import (
	// "log"
	// "os"
	"path/filepath"
	"flag"
	// "github.com/k0kubun/pp"
	// "gopkg.in/alecthomas/kingpin.v2"
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
	database string
)



func init() {
	flag.StringVar(&InFileName, "in", "*.csv", "File name of input csv")
	flag.StringVar(&database, "d", "perf", "Name of database to use")
	flag.Parse()
}

func main() {
	// パラメータをファイル名に展開
	inFileNames, err := filepath.Glob(InFileName)
	handleError(err)

	for _, inFileName := range inFileNames {
		handleCSV(inFileName)
	}
}
