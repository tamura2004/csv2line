package usecase

import (
	"fmt"
	"github.com/tamura2004/csv2line/domain"
	"io"
	"log"
	"os"
	"time"
)

type Presenter interface {
	WriteHeader(io.Writer)
	WriteCol(w io.Writer, key domain.Header, col string, timestamp int64)
	ExtName() string
}

type HeaderParser interface {
	ParseHeader([]string)
	Header(int) domain.Header
}

type PerfWriter struct {
	Count int
	Presenter
	HeaderParser
	io.Writer
}

func NewPerfWriter(inFile string, p Presenter) *PerfWriter {
	fileName := outFileName(inFile, p.ExtName())
	w, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}

	return &PerfWriter{
		Presenter:    p,
		HeaderParser: &PerfHeaderParser{},
		Writer:       w,
	}
}

func outFileName(inFile, extName string) string {
	basename := inFile[:len(inFile)-4] // len(".csv") == 4
	t := time.Now().Format("01021504")
	return fmt.Sprintf("%s_%s.%s", basename, t, extName)
}

func (w *PerfWriter) Write(record []string, lineNumber int) {
	if lineNumber == 1 {
		w.ParseHeader(record)
	} else {
		w.WriteRow(record)
	}
}

func (w *PerfWriter) WriteRow(row []string) {
	t := unixNano(row[0])

	// データを出力
	for j, col := range row[1:] {
		if col == " " || col == "0" {
			continue
		}
		w.WriteCol(w.Writer, w.Header(j), col, t)
		w.CountUp()
	}
}

func unixNano(s string) int64 {
	// タイムスタンプをナノ秒で取得
	t, err := time.Parse("01/02/2006 15:04:05.000 -0700", s+" +0900")
	if err != nil {
		log.Fatal(err)
	}

	return t.UnixNano()
}

func (w *PerfWriter) CountUp() {
	w.Count++
	if w.Count%100000 == 0 {
		log.Printf("Processed %d columns.\n", w.Count)
	}
}
