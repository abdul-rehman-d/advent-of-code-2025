package day11

import (
	"advent-of-code-2025/utils"
	"fmt"
	"strings"
)

func parse(line string) (string, []string) {
	raw := strings.Split(line, ": ")
	if len(raw) < 2 {
		panic("invalid input")
	}
	key := raw[0]
	vals := strings.Split(raw[1], " ")
	if len(vals) == 0 {
		panic("invalid input")
	}

	return key, vals
}

func PartA(data string) int {
	lines := utils.GetLines(data)

	mapp := make(map[string][]string, len(lines))
	for _, line := range lines {
		k, v := parse(line)
		mapp[k] = v
	}

	waysMap := make(map[string]int)

	var recurse func(k string) int

	recurse = func(k string) int {
		if k == "out" {
			return 1
		}

		if x, ok := waysMap[k]; ok {
			return x
		}

		out := 0
		v, ok := mapp[k]
		if !ok {
			panic(fmt.Sprintf("why doesnt map have %s", k))
		}

		for _, k := range v {
			out += recurse(k)
		}

		waysMap[k] = out
		return out
	}

	ans := recurse("you")

	return ans
}
