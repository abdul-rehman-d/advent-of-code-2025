package day10

import (
	"advent-of-code-2025/utils"
	"fmt"
	"strconv"
	"strings"
)

type A struct {
	ShouldBe            string
	Instructions        [][]int
	JoltageRequirements []int
}

func NewA(s string) A {
	i := 0
	a := A{
		ShouldBe:            "",
		Instructions:        [][]int{},
		JoltageRequirements: []int{},
	}
	for i < len(s) {
		ch := s[i]

		if ch == '[' {
			i++
			ch = s[i]
			for ch != ']' {
				a.ShouldBe += string(ch)
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

func PartA(data string) int {
	lines := utils.GetLines(data)

	as := make([]A, len(lines))

	for i, line := range lines {
		as[i] = NewA(line)
		fmt.Printf("%+v\n", as[i])
	}

	return 0
}
