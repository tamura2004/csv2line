package main

import (
	"fmt"
	"io"
	"log"
)

var count int

func writeCol(w io.Writer, key TagFieldKey, col string, timestamp int64) {
	tags := fmt.Sprintf("host=%s", key.hostname)

	if key.instance != "" {
		tags += fmt.Sprintf(",instance=%s", key.instance)
	}

	if key.index != "" {
		tags += fmt.Sprintf(",index=%s", key.index)
	}

	fmt.Fprintf(
		w,
		"%s,%s %s=%s %d\n",
		key.category,
		tags,
		key.label,
		col,
		timestamp,
	)

	count++
	if count%100000 == 0 {
		log.Printf("Processed %d columns.\n", count)
	}
}
