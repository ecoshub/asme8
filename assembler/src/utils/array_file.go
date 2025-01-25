package utils

import (
	"os"
	"strings"
)

func ReadArrayFile(path string) ([]string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	file := string(data)
	var lines []string
	if strings.Contains(file, "\r\n") {
		lines = strings.Split(file, "\r\n")
	} else {
		lines = strings.Split(file, "\n")
	}
	clean := make([]string, 0, len(lines))
	for _, l := range lines {
		c := strings.TrimSpace(l)
		if c == "" {
			continue
		}
		clean = append(clean, c)
	}
	return clean, nil
}
