package main

import (
	"fmt"
	"path/filepath"
	"time"
)

func buildFilename(infile string) (string, error) {
	ext := filepath.Ext(infile)
	if ext != ".csv" {
		return "", fmt.Errorf("infile must be csv: %s", infile)
	}

	basename := infile[:len(infile)-4] // len(".csv") == 4
	t := time.Now().Format("01021504")

	return fmt.Sprintf("%s_%s.line", basename, t), nil
}
