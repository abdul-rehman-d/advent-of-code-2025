package day6

import (
	"advent-of-code-2025/utils"
)

func blah(nums []int, multiply bool) int {
	a := 0
	if multiply {
		a = 1
	}
	for _, num := range nums {
		if multiply {
			a *= num
		} else {
			a += num
		}
	}
	return a
}

func PartB(data string) int {
	lines := utils.GetLines(data)

	max_rows := len(lines)
	max_cols := len(lines[0])

	ans := 0

	nums := []int{}
	multiply := false
	for c := range max_cols {
		x := 0
		for r := range max_rows {
			ch := lines[r][c]
			if ch == ' ' {
				continue
			} else if ch == '+' {
				multiply = false
			} else if ch == '*' {
				multiply = true
			} else {
				x = (x * 10) + int(ch-'0')
			}
		}
		if x == 0 {
			// empty
			ans += blah(nums, multiply)
			nums = []int{}
		} else {
			nums = append(nums, x)
		}
	}

	ans += blah(nums, multiply)

	return ans
}
