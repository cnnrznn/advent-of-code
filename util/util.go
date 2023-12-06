package util

import (
	"bytes"
	"os"
)

func ReadLines(fn string) ([]string, error) {
	bs, err := os.ReadFile(fn)
	if err != nil {
		return nil, err
	}

	bsSplit := bytes.Split(bs, []byte("\n"))

	lines := []string{}

	for _, bs := range bsSplit {
		lines = append(lines, string(bs))
	}

	return lines, nil
}
