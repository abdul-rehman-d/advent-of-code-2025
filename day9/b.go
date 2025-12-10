package day9

import (
	"advent-of-code-2025/utils"
	"fmt"
	"strconv"
	"strings"
)

func edgeIntersectsRect(
	x1, y1, x2, y2 int, // edge endpoints
	rx1, ry1, rx2, ry2 int, // rectangle (already sorted: rx1 <= rx2, ry1 <= ry2)
) bool {
	// horizontal edge
	if y1 == y2 {
		if ry1 < y1 && y1 < ry2 { // edge is vertically inside rect
			if max(x1, x2) > rx1 && min(x1, x2) < rx2 {
				// edge's x-range overlaps rect interior
				return true
			}
		}
	} else { // vertical edge (AoC edges are axis-aligned)
		if rx1 < x1 && x1 < rx2 { // edge is horizontally inside rect
			if max(y1, y2) > ry1 && min(y1, y2) < ry2 {
				// edge's y-range overlaps rect interior
				return true
			}
		}
	}
	return false
}

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
	inside := false
	px := float64(p.C)
	py := float64(p.R)

	for i, a := range coords {
		b := coords[(i+1)%len(coords)] // back to zero

		// check if point is ON segment
		if pointOnSegment(p, a, b) {
			return true
		}

		// raycast
		xi := float64(a.C)
		yi := float64(a.R)
		xj := float64(b.C)
		yj := float64(b.R)

		intersects := ((yi > py) != (yj > py)) &&
			(px < (xj-xi)*(py-yi)/(yj-yi)+xi)

		if intersects {
			inside = !inside
		}
	}

	return inside
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
			a := area(ci, cj)
			if a < largest {
				// no point if its not even largest yet
				continue
			}

			valid := true

			// check four corners
			// 9,5 & 2,3
			// 9,5 - 9,3 - 2,3 - 2,5
			corners := []Coord{
				{R: ci.R, C: ci.C},
				{R: cj.R, C: cj.C},
				{R: cj.R, C: ci.C},
				{R: ci.R, C: cj.C},
			}

			for _, p := range corners {
				if !checkInsidePolygon(p) {
					valid = false
					break
				}
			}

			for i, a := range coords {
				b := coords[(i+1)%len(coords)] // wrap around

				ex1, ey1 := a.C, a.R
				ex2, ey2 := b.C, b.R

				if edgeIntersectsRect(ex1, ey1, ex2, ey2, ci.C, cj.C, ci.R, cj.R) {
					valid = false
					break
				}
			}

			if valid {
				largest = a
			}
		}
	}

	return largest
}
