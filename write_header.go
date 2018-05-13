package main

import (
	"fmt"
	"io"
)

func writeHeader(w io.Writer) {
	fmt.Fprint(w, "# DDL\n")
	fmt.Fprintf(w, "CREATE DATABASE %s\n\n", database)
	fmt.Fprint(w, "# DML\n")
	fmt.Fprintf(w, "# CONTEXT-DATABASE: %s\n\n", database)
}
