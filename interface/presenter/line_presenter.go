package presenter

import (
	"fmt"
	"github.com/tamura2004/csv2line/domain"
	"github.com/tamura2004/csv2line/usecase"
	"io"
)

type LinePresenter struct {
	database string
}

func NewLinePresenter(database string) usecase.Presenter {
	return &LinePresenter{database}
}

func (f *LinePresenter) ExtName() string {
	return "line"
}

func (f *LinePresenter) WriteHeader(w io.Writer) {
	fmt.Fprint(w, "# DDL\n")
	fmt.Fprintf(w, "CREATE DATABASE %s\n\n", f.database)
	fmt.Fprint(w, "# DML\n")
	fmt.Fprintf(w, "# CONTEXT-DATABASE: %s\n\n", f.database)
}

func (f *LinePresenter) WriteCol(w io.Writer, key domain.Header, col string, timestamp int64) {
	tags := fmt.Sprintf("host=%s", key.Hostname)

	if key.Instance != "" {
		tags += fmt.Sprintf(",instance=%s", key.Instance)
	}

	if key.Index != "" {
		tags += fmt.Sprintf(",index=%s", key.Index)
	}

	fmt.Fprintf(
		w,
		"%s,%s %s=%s %d\n",
		key.Category,
		tags,
		key.Label,
		col,
		timestamp,
	)
}
