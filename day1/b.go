package day1

import (
	"strconv"
	"strings"
)

func PartB(data string) int {
	lines := strings.Split(data, "\n")

	x := 50
	ans := 0

	for _, instruction := range lines {
		if len(instruction) == 0 {
			continue // empty line at end probs
		}
		dir := instruction[:1]
		a, _ := strconv.Atoi(instruction[1:])
		for range a {
			switch dir {
			case "L":
				x--
			case "R":
				x++
			}
			switch x {
			case -1:
				x = 99
			case 100:
				x = 0
				ans++
			case 0:
				ans++
			}
		}
	}

	return ans
}
