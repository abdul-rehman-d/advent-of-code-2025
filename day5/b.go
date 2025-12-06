package day5

import (
	"cmp"
	"slices"
	"strconv"
	"strings"
)

func merge(ranges []Range) []Range {
	newRanges := []Range{}

	curr := ranges[0]

	for i := 1; i < len(ranges); i++ {
		r := ranges[i]

		// overlapping
		if (r.S >= curr.S && r.S <= curr.E) ||
			(r.E >= curr.S && r.E <= curr.E) {
			n := Range{
				S: min(curr.S, r.S),
				E: max(curr.E, r.E),
			}

			curr = n
			continue
		}

		newRanges = append(newRanges, curr)
		curr = r
	}

	newRanges = append(newRanges, curr)

	return newRanges
}

func PartB(data string) int {
	raw := strings.Split(data, "\n\n")
	if len(raw) < 2 {
		panic("wrong puzzle input")
	}
	rangesRaw := strings.Split(raw[0], "\n")
	ranges := make([]Range, len(rangesRaw))

	for i, r := range rangesRaw {
		raw := strings.Split(r, "-")
		if len(raw) < 2 {
			panic("wrong puzzle input")
		}
		s, _ := strconv.Atoi(raw[0])
		e, _ := strconv.Atoi(raw[1])

		ranges[i] = Range{S: s, E: e}
	}

	slices.SortFunc(ranges, func(a, b Range) int {
		if a.S == b.S {
			return cmp.Compare(a.E, b.E)
		}
		return cmp.Compare(a.S, b.S)
	})

	merged := merge(ranges)

	count := 0
	for _, r := range merged {
		count += (r.E - r.S + 1)
	}
	return count
}
