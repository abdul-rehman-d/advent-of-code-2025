package day2

import (
	"advent-of-code-2025/utils"
	"fmt"
	"strconv"
	"strings"
)

func isValid(x int) bool {
	if x < 10 {
		return true
	}
	s := fmt.Sprintf("%d", x)
	h := len(s) / 2
	return s[:h] != s[h:]
}

func PartA(data string) int {
	lines := utils.FilterEmptyLines(strings.Split(data, "\n"))

	if len(lines) != 1 {
		panic("wrong puzzle input")
	}

	ans := 0
	pairs := strings.SplitSeq(lines[0], ",")

	for pair := range pairs {
		splitted := strings.Split(pair, "-")
		if len(splitted) != 2 {
			panic("wrong puzzle input")
		}
		startStr := splitted[0]
		endStr := splitted[1]
		start, _ := strconv.Atoi(startStr)
		end, _ := strconv.Atoi(endStr)

		for x := start; x <= end; x++ {
			if !isValid(x) {
				ans += x
			}
		}
	}

	return ans
}
