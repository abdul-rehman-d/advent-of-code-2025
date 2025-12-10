package day9

import (
	"advent-of-code-2025/utils"
	"fmt"
	"strconv"
	"strings"
)

type Coord = utils.Coord

func area(a, b Coord) int {
	l := max(a.R, b.R) - min(a.R, b.R) + 1
	w := max(a.C, b.C) - min(a.C, b.C) + 1

	return int(l * w)
}

func PartA(data string) int {
	lines := utils.GetLines(data)
	coords := make([]Coord, len(lines))

	for i, line := range lines {
		a := strings.Split(line, ",")
		if len(a) < 2 {
			panic(fmt.Sprintf("less than 2 coordinates at ln %d", i))
		}
		x, _ := strconv.Atoi(a[0])
		y, _ := strconv.Atoi(a[1])

		coords[i] = Coord{R: x, C: y}
	}

	largest := 0

	for i, ci := range coords {
		for j, cj := range coords {
			if i == j {
				continue
			}
			a := area(ci, cj)
			if largest < a {
				largest = a
			}
		}
	}

	return largest
}
