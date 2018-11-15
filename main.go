package main

import (
	"flag"
	"github.com/tamura2004/csv2line/interface/controller"
	"github.com/tamura2004/csv2line/interface/presenter"
	"github.com/tamura2004/csv2line/usecase"
	"log"
	"os"
	"path/filepath"
)

type Params struct {
	path      string
	database  string
	format    string
	presenter usecase.Presenter
}

var params Params

func init() {
	params.Init()
}

func (p *Params) Init() {
	flag.StringVar(&p.database, "d", "perf", "Name of database to use")
	flag.StringVar(&p.format, "f", "line", "<line|csv> format of output file")
	flag.Parse()
	p.presenter = presenter.NewLinePresenter(params.database)
	if p.format == "csv" {
		p.presenter = presenter.NewCsvPresenter()
	}
	p.path = flag.Arg(0)
}

func main() {
	// パラメータをファイル名に展開
	inFileNames, err := filepath.Glob(params.path)
	if err != nil {
		log.Fatal(err)
	}

	for _, inFileName := range inFileNames {
		r := NewPerfReader(inFileName)
		w := NewPerfWriter(inFileName)
		w.WriteHeader(w.Writer)
		r.Copy(w)
	}
}

func NewPerfReader(inFileName string) *usecase.PerfReader {
	r, err := os.Open(inFileName)
	if err != nil {
		log.Fatal(err)
	}
	s := controller.NewCsvScanner(r)
	return usecase.NewPerfReader(s)
}

func NewPerfWriter(inFileName string) *usecase.PerfWriter {
	return usecase.NewPerfWriter(inFileName, params.presenter)
}
