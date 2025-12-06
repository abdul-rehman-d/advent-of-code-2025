package day6

import (
	"advent-of-code-2025/utils"
)

func PartA(data string) int {
	lines := utils.GetLines(data)

	nums := [][]int{}
	exps := []bool{} // multiply=true

	for i, s := range lines {
		if i == len(lines)-1 {
			for _, ch := range s {
				switch ch {
				case '+':
					exps = append(exps, false)
				case '*':
					exps = append(exps, true)
				}
			}
			continue
		}
		row := []int{}

		x := 0

		for i, ch := range s {
			if ch == ' ' {
				if x != 0 {
					row = append(row, x)
					x = 0
				}
				continue
			}

			x = (x * 10) + int(ch-'0')

			if i == len(s)-1 {
				row = append(row, x)
			}
		}

		nums = append(nums, row)
	}

	limit := len(nums[0])
	ans := 0

	for i := range limit {
		a := 0
		if exps[i] {
			a = 1
		}
		for _, num := range nums {
			if exps[i] {
				a *= num[i]
			} else {
				a += num[i]
			}
		}

		ans += a
	}

	return ans
}
