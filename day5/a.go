package day5

import (
	"strconv"
	"strings"
)

type Range struct {
	S int
	E int
}

func PartA(data string) int {
	raw := strings.Split(data, "\n\n")
	if len(raw) < 2 {
		panic("wrong puzzle input")
	}
	rangesRaw := strings.Split(raw[0], "\n")
	ids := strings.Split(raw[1], "\n")

	ranges := make([]Range, len(rangesRaw))

	for i, r := range rangesRaw {
		raw := strings.Split(r, "-")
		if len(raw) < 2 {
			panic("wrong puzzle input")
		}
		s, _ := strconv.Atoi(raw[0])
		e, _ := strconv.Atoi(raw[1])

		ranges[i] = Range{
			S: s,
			E: e,
		}
	}

	count := 0

	for _, idStr := range ids {
		id, _ := strconv.Atoi(idStr)

		for _, r := range ranges {
			if id >= r.S && id <= r.E {
				count++
				break
			}
		}
	}

	return count
}
