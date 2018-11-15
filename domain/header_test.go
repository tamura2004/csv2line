package domain

import (
	"testing"
)

func TestParseHostname(t *testing.T) {
	text := `\\host name\cate gory(in.stance#index)\label`

	want := Header{
		Hostname: "hostname",
		Category: "category",
		Instance: "in.stance",
		Index:    "index",
		Label:    "label",
	}

	have := NewHeader(text)
	if want != have {
		t.Fatalf("bad hostname want = %q, have = %q", want, have)
	}
}

func TestParseHostnameNoIndex(t *testing.T) {
	text := `\\host name\cate gory(in.stance)\label`

	want := Header{
		Hostname: "hostname",
		Category: "category",
		Instance: "in.stance",
		Index:    "",
		Label:    "label",
	}

	have := NewHeader(text)
	if want != have {
		t.Fatalf("bad hostname want = %q, have = %q", want, have)
	}
}

func TestParseHostnameNoInstance(t *testing.T) {
	text := `\\host name\cate gory\label`

	want := Header{
		Hostname: "hostname",
		Category: "category",
		Instance: "",
		Index:    "",
		Label:    "label",
	}

	have := NewHeader(text)
	if want != have {
		t.Fatalf("bad hostname want = %q, have = %q", want, have)
	}
}
