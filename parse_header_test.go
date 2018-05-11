package main

import (
	"testing"
)

func TestParseHeader(t *testing.T) {
	cases := []string{
		`\\hostname\category(instance#index)\label`,
		`\\hostname\category(instance)\label`,
		`\\hostname\category\label`,
	}

	for _, header := range cases {
		keys, err := parseHeader(header)
		if err != nil {
			t.Error(err)
		}

		if keys.label != "label" {
			t.Errorf("bad header: %q", keys)
		}
	}
}
