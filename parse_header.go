package main

import (
	"fmt"
	"regexp"
)

var word = regexp.MustCompile(`[^\\#()]+`) // "\","#","(",")"以外のすべての文字
var space = regexp.MustCompile(`\s+`)      // 上記から[a-zA-Z0-9_]を除いた文字

func parseHeader(header string) (TagFieldKey, error) {
	var keys TagFieldKey

	header = space.ReplaceAllString(header, `_`)

	xs := word.FindAllString(header, -1)
	if len(xs) < 3 {
		return keys, fmt.Errorf("bad header: %s", header)
	}

	keys.hostname = xs[0]
	keys.category = xs[1]
	switch len(xs) {
	case 3:
		keys.label = xs[2]
	case 4:
		keys.instance = xs[2]
		keys.label = xs[3]
	case 5:
		keys.instance = xs[2]
		keys.index = xs[3]
		keys.label = xs[4]
	default:
		return keys, fmt.Errorf("bad header: %s", header)
	}

	return keys, nil
}
