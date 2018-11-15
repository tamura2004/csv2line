package controller

import (
	"encoding/csv"
	"io"
	"log"
)

type CsvScanner struct {
	csv        *csv.Reader
	err        error
	record     []string
	linenumber int
}

func NewCsvScanner(r io.Reader) *CsvScanner {
	csv := csv.NewReader(r)
	return &CsvScanner{csv: csv}
}

func (s *CsvScanner) Scan() bool {
	if s.eof() {
		return false
	}
	s.record, s.err = s.csv.Read()
	s.linenumber++
	return !s.eof()
}

func (s *CsvScanner) eof() bool {
	if s.err == io.EOF {
		return true
	}

	if s.err == nil {
		return false
	} else {
		log.Fatal(s.err)
	}
	return true
}

func (s CsvScanner) Record() []string {
	return s.record
}

func (s CsvScanner) LineNumber() int {
	return s.linenumber
}
