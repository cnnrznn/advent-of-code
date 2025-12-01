package util

import (
	"bufio"
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

func ReadStdio() ([]string, error) {
	var reader *bufio.Reader

	args := os.Args[1:]
	if len(args) > 0 {
		f, err := os.Open(args[0])
		if err != nil {
			return nil, err
		}
		reader = bufio.NewReader(f)
	} else {
		reader = bufio.NewReader(os.Stdin)
	}

	var lines []string

	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, nil
}
