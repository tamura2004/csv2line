package main

import (
	"errors"
	"regexp"
	"testing"
)

func TestCreateOutFile(t *testing.T) {

	cases := []struct {
		infile string
		outReg string
		err    error
	}{
		{infile: "test.csv", outReg: `^test_\d{8}.line$`, err: nil},
		{infile: "test.line", outReg: `.`, err: errors.New("infile must be csv: test.line")},
	}

	for _, c := range cases {
		outfile, err := createOutFile(c.infile)

		if err != nil {
			if err.Error() != c.err.Error() {
				t.Errorf("Bad error want=%q got=%q", c.err, err)
			}
		} else {
			re, err := regexp.Compile(c.outReg)

			if err != nil {
				t.Errorf("Bad regexp: %v", c.outReg)
			}

			if !re.MatchString(outfile) {
				t.Errorf("Bad outfile want=%v got=%v", c.outReg, outfile)
			}
		}
	}
}
