package main

import (
	"fmt"
	"io"
)

func writeHeader(w io.Writer) {
	fmt.Fprint(w, "# DDL\n")
	fmt.Fprint(w, "CREATE DATABASE perf\n\n")
	fmt.Fprint(w, "# DML\n")
	fmt.Fprint(w, "# CONTEXT-DATABASE: perf\n\n")
}
