package day8

import (
	"advent-of-code-2025/utils"
	"fmt"
	"math"
)

func getShortest2(distMap [][]Meh) (int, int, bool) {
	shortest := math.MaxFloat64
	i, j := -1, -1
	for idx, d := range distMap {
		if len(d) == 0 {
			continue
		}
		if d[0].D < shortest {
			shortest = d[0].D
			j = d[0].I
			i = idx
		}
	}

	if i == -1 && j == -1 {
		return i, j, false
	}

	return i, j, true
}

func PartB(data string) int {
	lines := utils.GetLines(data)

	total_junctions := len(lines)

	coords := make([]Coord, total_junctions)

	for i, line := range lines {
		coords[i] = NewCoord(line)
	}

	circuits := make([]*Circuit, total_junctions)
	for i := range total_junctions {
		circuits[i] = &Circuit{Id: i, Len: 1}
	}

	distMap := parseDistanceMap(coords)

	i, j := -1, -1

	for {
		has := true
		i, j, has = getShortest2(distMap)
		if !has {
			break
		}

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
			if l == total_junctions {
				break
			}
		}
		// else: stay where you are
		// pop the shortest one (both sides)
		distMap[i] = distMap[i][1:]
		distMap[j] = distMap[j][1:]
	}

	product := 1

	fmt.Printf("%v - %v\n", coords[i], coords[j])

	product = int(coords[i].X) * int(coords[j].X)

	return product
}
