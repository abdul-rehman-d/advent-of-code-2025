package day4

import "advent-of-code-2025/utils"

const MAX_ADJACENT_ALLOWED = 4

func countAdjacent(grid []string, r, c int) int {
	count := 0

	dirs := []struct {
		R int
		C int
	}{
		{-1, 0},  // up
		{1, 0},   // down
		{0, 1},   // right
		{0, -1},  // left
		{-1, -1}, // upper left
		{-1, 1},  // upper right
		{1, -1},  // lower left
		{1, 1},   // lower right
	}

	for _, dir := range dirs {
		targetR := r + dir.R
		targetC := c + dir.C
		if targetR < 0 || targetR >= len(grid) || targetC < 0 || targetC >= len(grid[0]) {
			continue
		}
		if grid[targetR][targetC] == '@' {
			count++
		}
	}

	return count
}

func PartA(data string) int {
	grid := utils.GetLines(data)

	MAX_ROWS := len(grid)
	MAX_COLS := len(grid[0])

	ans := 0

	for r := range MAX_ROWS {
		for c := range MAX_COLS {
			if grid[r][c] != '@' {
				continue
			}
			if countAdjacent(grid, r, c) < MAX_ADJACENT_ALLOWED {
				ans++
			}
		}
	}

	return ans
}
