package day10

import (
	"advent-of-code-2025/utils"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Machine struct {
	ShouldBe            []bool
	Instructions        [][]int
	JoltageRequirements []int
}

func PartA(data string) int {
	lines := utils.GetLines(data)

	machines := make([]Machine, len(lines))

	for i, line := range lines {
		machines[i] = NewMachine(line)
	}

	ans := 0

	for _, machine := range machines {
		possibilities := getPossibilities(len(machine.Instructions))
		a := getShortest(machine, possibilities)
		if a == -1 {
			panic(fmt.Sprintf("issue at machine %+v", machine))
		}
		ans += a
	}

	return ans
}

func getShortest(m Machine, possibilities [][]bool) int {
	shortest := math.MaxInt
	for _, p := range possibilities {
		count := 0
		state := make([]bool, len(m.ShouldBe))
		for a, b := range p {
			if !b {
				continue
			}
			count++
			for _, i := range m.Instructions[a] {
				state[i] = !state[i]
			}
		}

		if count < shortest && utils.CompareBoolArr(state, m.ShouldBe) {
			shortest = count
		}
	}
	return shortest
}

func getPossibilities(n int) [][]bool {
	out := [][]bool{}
	for mask := 1; mask < (1 << n); mask++ {
		vec := make([]bool, n)

		for i := 0; i < n; i++ {
			vec[i] = mask&(1<<i) != 0
		}

		out = append(out, vec)
	}
	return out
}

func NewMachine(s string) Machine {
	i := 0
	a := Machine{
		ShouldBe:            []bool{},
		Instructions:        [][]int{},
		JoltageRequirements: []int{},
	}
	for i < len(s) {
		ch := s[i]

		if ch == '[' {
			i++
			ch = s[i]
			for ch != ']' {
				a.ShouldBe = append(a.ShouldBe, ch == '#')
				i++
				ch = s[i]
			}
		} else if ch == '(' {
			pointsRaw := ""
			i++
			ch = s[i]
			for ch != ')' {
				pointsRaw += string(ch)
				i++
				ch = s[i]
			}
			raw := strings.Split(pointsRaw, ",")
			ins := make([]int, len(raw))
			for i, xr := range raw {
				x, _ := strconv.Atoi(xr)
				ins[i] = x
			}
			a.Instructions = append(a.Instructions, ins)
		} else if ch == '{' {
			joltageStr := ""
			i++
			ch = s[i]
			for ch != '}' {
				joltageStr += string(ch)
				i++
				ch = s[i]
			}
			raw := strings.Split(joltageStr, ",")
			joltage := make([]int, len(raw))
			for i, xr := range raw {
				x, _ := strconv.Atoi(xr)
				joltage[i] = x
			}
			a.JoltageRequirements = joltage
			return a
		}
		i++
	}
	return a
}
