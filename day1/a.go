package day1

import (
	"strconv"
	"strings"
)

func PartA(data string) int {
	lines := strings.Split(data, "\n")

	x := 50
	ans := 0

	for _, instruction := range lines {
		if len(instruction) == 0 {
			continue // empty line at end probs
		}
		dir := instruction[:1]
		a, _ := strconv.Atoi(instruction[1:])
		switch dir {
		case "L":
			x -= a
		case "R":
			x += a
		}
		if x < 0 {
			for x < 0 {
				x += 100
			}
		} else if x >= 100 {
			for x >= 100 {
				x -= 100
			}
		}

		if x == 0 {
			ans++
		}
	}

	return ans
}
