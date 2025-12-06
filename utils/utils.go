package utils

import "strings"

func GetLines(data string) []string {
	return FilterEmptyLines(strings.Split(data, "\n"))
}

func FilterEmptyLines(data []string) []string {
	out := make([]string, 0, len(data))
	for _, line := range data {
		if line != "" {
			out = append(out, line)
		}
	}
	return out
}
