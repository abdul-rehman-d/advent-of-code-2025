package day11

import (
	"advent-of-code-2025/utils"
	"fmt"
)

type Args struct {
	K       string
	Visited string
}

func PartB(data string) int {
	lines := utils.GetLines(data)

	mapp := make(map[string][]string, len(lines))
	for _, line := range lines {
		k, v := parse(line)
		mapp[k] = v
	}

	memo := make(map[Args]int)

	var recurse func(k string, visited string) int

	recurse = func(k string, visited string) int {
		visited2 := "" + visited

		switch k {
		case "fft":
			visited2 = "#" + visited2[1:]
		case "dac":
			visited2 = visited2[:1] + "#"
		}

		a := Args{K: k, Visited: visited2}

		if x, ok := memo[a]; ok {
			return x
		}

		if k == "out" {
			if visited2 == "##" {
				memo[a] = 1
			} else {
				memo[a] = 0
			}
			return memo[a]
		}

		v, ok := mapp[k]
		if !ok {
			panic(fmt.Sprintf("why doesnt map have %s", k))
		}

		out := 0
		for _, k := range v {
			out += recurse(k, visited2)
		}

		memo[a] = out
		return out
	}

	allNexts, ok := mapp["svr"]
	if !ok {
		panic("no why svr?")
	}

	ans := 0
	for _, p := range allNexts {
		x := recurse(p, "..")
		ans += x
	}

	return ans
}
