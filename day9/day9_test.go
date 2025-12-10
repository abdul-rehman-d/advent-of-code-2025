package day9

import (
	"testing"
)

const (
	data = `7,1
11,1
11,7
9,7
9,5
2,5
2,3
7,3`
)

func TestPartA(t *testing.T) {
	expected := 50
	result := PartA(data)

	if expected != result {
		t.Fatalf("\nExpected = %d\nResult = %d\n", expected, result)
	}
}

func TestPartB(t *testing.T) {
	expected := 24
	result := PartB(data)

	if expected != result {
		t.Fatalf("\nExpected = %d\nResult = %d\n", expected, result)
	}
}

func TestCheckInside(t *testing.T) {
	polygon := []Coord{
		{R: 1, C: 7},
		{R: 1, C: 11},
		{R: 3, C: 11},
		{R: 3, C: 7},
	}

	toCheck := []Coord{
		{R: 1, C: 7},
		{R: 1, C: 11},
		{R: 3, C: 11},
		{R: 3, C: 7},
		{R: 2, C: 9},
		{R: 2, C: 7},
		{R: 2, C: 11},
	}

	for _, p := range toCheck {
		result := CheckInside(polygon, p)
		if !result {
			t.Fatalf("\nExpected = true\nResult = false\tp = %+v\n", p)
		}
	}

}
