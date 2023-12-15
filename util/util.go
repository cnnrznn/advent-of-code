package util

import (
	"bytes"
	"os"
	"strings"
)

func ReadLines(fn string) ([]string, error) {
	bs, err := os.ReadFile(fn)
	if err != nil {
		return nil, err
	}

	bsSplit := bytes.Split(bs, []byte("\n"))

	lines := []string{}

	for _, bs := range bsSplit {
		fields := strings.Fields(string(bs))
		if len(fields) == 0 {
			continue
		}

		lines = append(lines, string(bs))
	}

	return lines, nil
}
