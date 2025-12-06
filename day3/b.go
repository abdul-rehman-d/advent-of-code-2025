package day3

import (
	"advent-of-code-2025/utils"
)

const (
	NUM = 12
)

// 9876 54321111111
// 9 8765 4321111111
// 98 7654 321111111

// 2342 34234234278
// 234 23 4234234278

func PartB(data string) int {
	lines := utils.GetLines(data)

	ans := 0
	for _, line := range lines {
		start := 0
		jolt := 0
		for i := range NUM {
			largest := 0

			limit := len(line) - NUM + i + 1

			for k := start; k < limit; k++ {
				curr := int(line[k] - '0')
				if curr > largest {
					largest = curr
					start = k + 1
				}
			}

			jolt = (jolt * 10) + largest
		}
		ans += jolt
	}

	return ans
}
