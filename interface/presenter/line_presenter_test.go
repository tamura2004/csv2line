package presenter

import (
	"github.com/tamura2004/csv2line/domain"
	"os"
	"testing"
)

func TestLinePresenterExtName(t *testing.T) {
	p := NewLinePresenter("perf")
	have := p.ExtName()
	want := "line"
	if want != have {
		t.Fatalf("bad hostname want = %q, have = %q", want, have)
	}
}

func ExampleLinePresenterWriteCol() {
	p := NewLinePresenter("perf")
	h := domain.NewHeader(`\\hostname\category(instance#index)\label`)
	p.WriteCol(os.Stdout, h, "value", 0)
	// Output:
	// category,host=hostname,instance=instance,index=index label=value 0
}

func ExampleLinePresenterWriteHeader() {
	p := NewLinePresenter("perf")
	p.WriteHeader(os.Stdout)
	// Output:
	// # DDL
	// CREATE DATABASE perf
	//
	// # DML
	// # CONTEXT-DATABASE: perf
}
