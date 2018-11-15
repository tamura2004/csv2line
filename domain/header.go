package domain

import (
	"log"
	"regexp"
)

// Field Header: \\hostname\category(instance#index)\label
// example: \\T2CH8300\Process(svchost#5)\Pool Paged Bytes
type Header struct {
	Hostname string
	Category string
	Instance string
	Index    string
	Label    string
}

const WORD = `([^\\#()]*)`

var pattern = regexp.MustCompile(
	`^` + `\\\\` + WORD + //hostname
		`\\` + WORD + //category
		`\(?` + WORD + //instance
		`#?` + WORD + //index
		`\)?` +
		`\\` + WORD + //label
		`$`,
)
var space = regexp.MustCompile(`\s+`)

func NewHeader(text string) Header {
	text = space.ReplaceAllString(text, ``)
	result := pattern.FindAllStringSubmatch(text, -1)

	if len(result) == 0 {
		log.Fatalf("bad header: %s", text)
	}

	h := result[0][1:]

	return Header{
		Hostname: h[0],
		Category: h[1],
		Instance: h[2],
		Index:    h[3],
		Label:    h[4],
	}
}
