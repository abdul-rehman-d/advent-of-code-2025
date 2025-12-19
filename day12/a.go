package day12

import (
	"fmt"
	"strconv"
	"strings"
)

type Shape struct {
	shape string
}

type Region struct {
	Rows       int
	Cols       int
	Quantities []int
}

func PartA(data string) int {
	splitted := strings.Split(data, "\n\n")
	if len(splitted) < 2 {
		panic("invalid input")
	}

	regions := parseRegions(splitted[len(splitted)-1])

	count := 0

	for _, region := range regions {
		can := (region.Cols / 3) * (region.Rows / 3)
		fmt.Println(can, sum(region.Quantities))
		if can >= sum(region.Quantities) {
			count++
		}
	}

	return count
}

func sum(a []int) int {
	b := 0
	for _, a := range a {
		b += a
	}
	return b
}

func parseRegions(str string) []Region {
	splitted := strings.Split(str, "\n")

	regions := make([]Region, len(splitted)-1)

	for i, s := range splitted[:len(splitted)-1] {
		fmt.Println(s)
		inner := strings.Split(s, " ")

		x, y := parseLengthAndWidth(inner[0])

		regions[i] = Region{
			Rows:       y,
			Cols:       x,
			Quantities: parseInts(inner[1:]),
		}
	}

	return regions
}

func parseInts(strs []string) []int {
	ints := make([]int, len(strs))
	for i, s := range strs {
		x, _ := strconv.Atoi(s)
		ints[i] = x
	}
	return ints
}

func parseLengthAndWidth(s string) (int, int) {
	a := strings.TrimSuffix(s, ":")
	b := strings.Split(a, "x")
	if len(b) != 2 {
		panic(fmt.Sprintf("invalid input in regions %s - %s - %v", s, a, b))
	}

	x, _ := strconv.Atoi(b[0])
	y, _ := strconv.Atoi(b[1])

	return x, y
}
