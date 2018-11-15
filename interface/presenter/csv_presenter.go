package presenter

import (
	"fmt"
	"github.com/tamura2004/csv2line/domain"
	"github.com/tamura2004/csv2line/usecase"
	"io"
	"time"
)

type CsvPresenter struct {
}

func NewCsvPresenter() usecase.Presenter {
	return &CsvPresenter{}
}

func (c *CsvPresenter) ExtName() string {
	return "csv"
}

func (c *CsvPresenter) WriteHeader(w io.Writer) {
	fmt.Fprint(w, "date,time,host,category,instance,index,label,value\n")
}

func (c *CsvPresenter) WriteCol(w io.Writer, key domain.Header, col string, timestamp int64) {
	t := time.Unix(0, timestamp/10000000000*10000000000) // round up 10s
	fmt.Fprintf(
		w,
		"%s,%s,%s,%s,%s,%s,%s,%s\n",
		t.Format("2006-01-02"),
		t.Format("15:04:05"),
		key.Hostname,
		key.Category,
		key.Instance,
		key.Index,
		key.Label,
		col,
	)
}
