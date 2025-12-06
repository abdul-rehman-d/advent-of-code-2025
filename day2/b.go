package day2

import (
	"advent-of-code-2025/utils"
	"fmt"
	"strconv"
	"strings"
)

func isValid2(x int) bool {
	if x < 10 {
		return true
	}
	xStr := fmt.Sprintf("%d", x)
	halfIdx := len(xStr) / 2
outer:
	for a := halfIdx; a > 0; a-- {
		if len(xStr)%a != 0 {
			continue
		}
		target := xStr[:a]
		// x=12341234; halfIdx=4; a=4; target=1234
		// loop limt=8/4=2
		// i=1; xStr[4:8]=1234
		for i := 1; i < len(xStr)/a; i++ {
			curr := xStr[i*a : (i+1)*a]
			if curr != target {
				continue outer
			}
		}
		return false
	}
	return true
}

func PartB(data string) int {
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
			if !isValid2(x) {
				ans += x
			}
		}
	}

	return ans
}
