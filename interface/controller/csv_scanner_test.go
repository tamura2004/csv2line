package controller

import (
	"reflect"
	"strings"
	"testing"
)

func TestCsvScannerScanLen(t *testing.T) {
	csv := strings.NewReader("1,2,3,4,5\n1,2,3,4,5\n1,2,3,4,5")
	s := NewCsvScanner(csv)

	for s.Scan() {
		have := len(s.Record())
		want := 5
		if have != want {
			t.Fatalf("bad field number, want = %q have = %q", want, have)
		}
	}
}

func TestCsvScannerScanRecord(t *testing.T) {
	csv := strings.NewReader("a,b\nc,d")
	s := NewCsvScanner(csv)
	testcase := [][]string{
		{"a", "b"},
		{"c", "d"},
	}

	for s.Scan() {
		have := s.Record()
		want := testcase[s.LineNumber()-1]

		if !reflect.DeepEqual(have, want) {
			t.Fatalf("bad field, want = %q have = %q", want, have)
		}
	}
}
