package main

import (
	"encoding/csv"
	"io"
)

func csvReadWithIndex(r io.Reader, yield func(row []string, i int)) {
	c := csv.NewReader(r)
	i := 0
	for {
		row, err := c.Read()
		if err == io.EOF {
			return
		} else if err != nil {
			handleError(err)
		} else {
			yield(row, i)
			i++
		}
	}
}
