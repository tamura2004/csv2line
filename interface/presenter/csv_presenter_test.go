package presenter

import (
	"github.com/tamura2004/csv2line/domain"
	"os"
	"testing"
)

func TestCsvPresenterExtName(t *testing.T) {
	p := NewCsvPresenter()
	have := p.ExtName()
	want := "csv"
	if want != have {
		t.Fatalf("bad hostname want = %q, have = %q", want, have)
	}
}

func ExampleCsvPresenterWriteCol() {
	p := NewCsvPresenter()
	h := domain.NewHeader(`\\hostname\category(instance#index)\label`)
	p.WriteCol(os.Stdout, h, "value", 0)
	// Output:
	// 1970-01-01,09:00:00,hostname,category,instance,index,label,value
}

func ExampleCsvPresenterWriteHeader() {
	p := NewCsvPresenter()
	p.WriteHeader(os.Stdout)
	// Output:
	// date,time,host,category,instance,index,label,value
}
