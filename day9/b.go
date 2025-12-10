package day9

import (
	"advent-of-code-2025/utils"
	"fmt"
	"strconv"
	"strings"
)

func pointOnSegment(p, a, b Coord) bool {
	// Treat C as x, R as y
	ax, ay := a.C, a.R
	bx, by := b.C, b.R
	px, py := p.C, p.R

	// Cross product must be zero for collinearity
	cross := (bx-ax)*(py-ay) - (by-ay)*(px-ax)
	if cross != 0 {
		return false
	}

	// And p must be within the bounding box of a and b
	if px < min(ax, bx) || px > max(ax, bx) {
		return false
	}
	if py < min(ay, by) || py > max(ay, by) {
		return false
	}

	return true
}

func CheckInside(coords []Coord, p Coord) bool {
	count := 0

	for i, a := range coords {
		b := coords[(i+1)%len(coords)] // back to zero

		// check if point is ON segment
		if pointOnSegment(p, a, b) {
			return true
		}

		// raycast
		ay, ax := float64(a.R), float64(a.C)
		by, bx := float64(b.R), float64(b.C)
		py, px := float64(p.R), float64(p.C)

		if ((py < ay) != (py < by)) &&
			(px < ax+(py-ay)*(bx-ax)/(by-ay)) {
			count++
		}
	}

	return count%2 == 1 // odd=inside
}

func PartB(data string) int {
	lines := utils.GetLines(data)
	coords := make([]Coord, len(lines))

	maxR := 0
	maxC := 0

	for i, line := range lines {
		a := strings.Split(line, ",")
		if len(a) < 2 {
			panic(fmt.Sprintf("less than 2 coordinates at ln %d", i))
		}
		x, _ := strconv.Atoi(a[0])
		y, _ := strconv.Atoi(a[1])

		if y > maxR {
			maxR = y
		}
		if x > maxC {
			maxC = x
		}

		coords[i] = Coord{C: x, R: y}
	}

	maxR++
	maxC++

	fmt.Println(maxR, maxC)

	set := make(map[Coord]bool)

	var checkInsidePolygon func(p Coord) bool

	checkInsidePolygon = func(p Coord) bool {
		if isInside, ok := set[p]; ok {
			return isInside
		}

		out := CheckInside(coords, p)
		set[p] = out
		return out
	}

	largest := 0

	for i, ci := range coords {
		for j, cj := range coords {
			if i == j {
				continue
			}
			valid := true

			// check four corners
			// 9,5 & 2,3
			// 9,5 - 9,3 - 2,3 - 2,5
			toCheck := []Coord{
				{R: ci.R, C: ci.C},
				{R: cj.R, C: cj.C},
				{R: cj.R, C: ci.C},
				{R: ci.R, C: cj.C},
			}

			for _, p := range toCheck {
				aa := checkInsidePolygon(p)
				if !aa {
					valid = false
				}
			}

			if valid {
				a := area(ci, cj)
				if largest < a {
					largest = a
				}
			}
		}
	}

	return largest
}
