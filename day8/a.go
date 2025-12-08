package day8

import (
	"advent-of-code-2025/utils"
	"cmp"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Coord struct {
	X float64
	Y float64
	Z float64
}

func (c Coord) Dist(a Coord) float64 {
	x := math.Pow(a.X-c.X, 2)
	y := math.Pow(a.Y-c.Y, 2)
	z := math.Pow(a.Z-c.Z, 2)
	return math.Sqrt(x + y + z)
}

func NewCoord(str string) Coord {
	c := strings.Split(str, ",")
	if len(c) < 3 {
		panic("less than 3 coords")
	}
	x, _ := strconv.Atoi(c[0])
	y, _ := strconv.Atoi(c[1])
	z, _ := strconv.Atoi(c[2])

	return Coord{
		X: float64(x),
		Y: float64(y),
		Z: float64(z),
	}
}

type Meh struct {
	D float64
	I int
}

func parseDistanceMap(coords []Coord) [][]Meh {
	distMap := make([][]Meh, len(coords)) // 0(n^2)

	// 0(n)
	for i := range len(coords) {
		dists := make([]Meh, 0, len(coords))
		for j := range len(coords) {
			if i == j {
				continue
			}
			a := coords[i]
			b := coords[j]
			d := a.Dist(b)
			dists = append(dists, Meh{d, j})
		}
		slices.SortFunc(dists, func(a, b Meh) int {
			return cmp.Compare(a.D, b.D)
		})
		distMap[i] = dists
	}
	return distMap
}

func getShortest(distMap [][]Meh) (int, int) {
	shortest := math.MaxFloat64
	i, j := 0, 0
	for idx, d := range distMap {
		if d[0].D < shortest {
			shortest = d[0].D
			j = d[0].I
			i = idx
		}
	}

	return i, j
}

type Circuit struct {
	Id  int
	Len int
}

const MAX = 10

func PartA(data string) int {
	lines := utils.GetLines(data)

	max_iters := MAX
	if m := os.Getenv("MAX"); m != "" {
		mx, err := strconv.Atoi(m)
		if err == nil {
			max_iters = mx
		}
	}

	coords := make([]Coord, len(lines))

	for i, line := range lines {
		coords[i] = NewCoord(line)
	}

	circuits := make([]*Circuit, len(lines))
	for i := range len(lines) {
		circuits[i] = &Circuit{Id: i, Len: 1}
	}

	distMap := parseDistanceMap(coords)

	for range max_iters {
		i, j := getShortest(distMap)

		if circuits[i].Id != circuits[j].Id {
			l := circuits[i].Len + circuits[j].Len
			oldId := -1
			newC := &Circuit{}
			if circuits[j].Len < circuits[i].Len {
				// go to i
				oldId = circuits[j].Id
				newC = circuits[i]
			} else {
				// go to j
				oldId = circuits[i].Id
				newC = circuits[j]
			}
			// update all old ids
			for k := range len(circuits) {
				if circuits[k].Id == oldId {
					circuits[k] = newC
				}
			}
			circuits[i].Len = l
		}
		// else: stay where you are
		// pop the shortest one (both sides)
		distMap[i] = distMap[i][1:]
		distMap[j] = distMap[j][1:]
	}

	slices.SortFunc(circuits, func(i, j *Circuit) int {
		return cmp.Compare(j.Len, i.Len)
	})

	product := 1
	count := 0
	lastIdx := -1
	for _, c := range circuits {
		if c.Id == lastIdx {
			continue
		}
		lastIdx = c.Id
		product *= int(c.Len)
		count++
		if count == 3 {
			break
		}
	}

	return product
}
