package day7

import "advent-of-code-2025/utils"

type Coord = utils.Coord

func PartB(data string) int {
	grid := utils.NewGrid(data)
	start, _ := grid.Find('S') // can ignore `has`

	possibilities := make(map[Coord]int)

	count := 0

	var recurse func(x Coord) int

	recurse = func(x Coord) int {
		if p, ok := possibilities[x]; ok {
			return p
		}
		out := 0
		if ch, ok := grid.Get(x.R, x.C); ok {
			switch ch {
			case '^':
				left := recurse(x.Add(Coord{R: 0, C: -1}))
				right := recurse(x.Add(Coord{R: 0, C: 1}))
				out = left + right
			case '.':
				out = recurse(x.Add(Coord{R: 1, C: 0}))
			}
		} else {
			out = 1 // end
		}
		possibilities[x] = out
		return out
	}

	count = recurse(start.Add(Coord{R: 1, C: 0}))

	return count
}
