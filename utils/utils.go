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

func ReplaceStringAtIndex(s string, idx int, ch byte) string {
	out := ""
	for i := 0; i < len(s); i++ {
		if i == idx {
			out += string(ch)
		} else {
			out += string(s[i])
		}
	}
	return out
}

func CompareBoolArr(a1, a2 []bool) bool {
	if len(a1) != len(a2) {
		return false
	}

	for i, a := range a1 {
		if a != a2[i] {
			return false
		}
	}

	return true
}
