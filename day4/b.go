package day4

import "advent-of-code-2025/utils"

func PartB(data string) int {
	grid := utils.GetLines(data)

	MAX_ROWS := len(grid)
	MAX_COLS := len(grid[0])

	ans := 0

	count := 1 // dummy
outer:
	for count != 0 {
		count = 0

		newGrid := make([]string, len(grid))
		copy(newGrid, grid)

		for r := range MAX_ROWS {
			for c := range MAX_COLS {
				if grid[r][c] != '@' {
					continue
				}
				if countAdjacent(grid, r, c) < MAX_ADJACENT_ALLOWED {
					count++
					newGrid[r] = utils.ReplaceStringAtIndex(newGrid[r], c, '.')
				}
			}
		}
		if count == 0 {
			continue outer
		}
		ans += count
		grid = newGrid
	}

	return ans
}
