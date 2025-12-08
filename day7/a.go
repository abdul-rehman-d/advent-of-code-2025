package day7

import (
	"advent-of-code-2025/utils"
)

func PartA(data string) int {
	grid := utils.NewGrid(data)
	start, _ := grid.Find('S') // can ignore `has`

	set := make(map[utils.Coord]bool)

	count := 0

	var recurse func(r, c int)

	recurse = func(r, c int) {
		curr := utils.Coord{R: r, C: c}
		if set[curr] {
			return
		}
		// if is_split {
		// 	count++
		// }
		set[curr] = true
		if ch, ok := grid.Get(r, c); ok {
			switch ch {
			case '^':
				count++
				recurse(r, c-1) // left
				recurse(r, c+1) // right
				return
			case '.':
				recurse(r+1, c)
				return
			default:
				return

			}
		} else {
			return
		}
	}

	recurse(start.R+1, start.C)

	return count
}
