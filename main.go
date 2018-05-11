package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"
)

type Param struct {
	database string
}

var param Param

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
	flag.StringVar(&param.database, "d", "perf", "Database to connect to the server.")
}

func main() {
	// パラメータをファイル名に展開
	inFileNames, err := filepath.Glob(os.Args[1])
	handleError(err)

	for _, inFileName := range inFileNames {
		handleCSV(inFileName)
	}
}
