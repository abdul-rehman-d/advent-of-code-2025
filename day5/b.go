package day5

import (
	"cmp"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func merge(ranges []Range) []Range {
	newRanges := []*Range{}
	for i := 0; i < len(ranges)-1; i++ {
		curr := ranges[i]
		next := ranges[i+1]

		fmt.Println(curr, next, i)

		// overlapping
		if (next.S >= curr.S && next.S <= curr.E) ||
			(next.E >= curr.S && next.E <= curr.E) {
			n := Range{
				S: min(curr.S, next.S),
				E: max(curr.E, next.E),
			}

			newRanges = append(newRanges, nil)
			ranges[i+1] = n
		} else {
			newRanges = append(newRanges, &curr)
		}

		fmt.Println(i, len(ranges)-1)
	}

	newRanges = append(newRanges, &ranges[len(ranges)-1])

	out := []Range{}
	for _, r := range newRanges {
		if r == nil {
			continue
		}
		out = append(out, *r)
	}

	return out
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

	for _, r := range ranges {
		fmt.Println(r)
	}

	merged := merge(ranges)

	fmt.Println("")
	for _, r := range merged {
		fmt.Println(r)
	}

	count := 0
	for _, r := range merged {
		count += (r.E - r.S + 1)
	}
	return count
}
