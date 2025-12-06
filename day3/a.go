package day3

import (
	"advent-of-code-2025/utils"
)

func PartA(data string) int {
	lines := utils.GetLines(data)

	ans := 0

	for _, line := range lines {
		arr := make([]int, len(line))
		for i, ch := range line {
			arr[i] = int(ch - '0')
		}

		largestFirstBatteryIdx := 0
		for i := 1; i < len(arr)-1; i++ {
			if arr[i] > arr[largestFirstBatteryIdx] {
				largestFirstBatteryIdx = i
			}
		}

		largestSecondBattery := 0
		for i := largestFirstBatteryIdx + 1; i < len(arr); i++ {
			largestSecondBattery = max(largestSecondBattery, arr[i])
		}

		ans += (arr[largestFirstBatteryIdx] * 10) + largestSecondBattery
	}

	return ans
}
