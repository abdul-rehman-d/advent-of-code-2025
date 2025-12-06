package day1

import (
	"testing"
)

const (
	data = `L68
L30
R48
L5
R60
L55
L1
L99
R14
L82`
)

func TestPartA(t *testing.T) {
	expected := 3
	result := PartA(data)

	if expected != result {
		t.Fatalf("\nExpected = %d\nResult = %d\n", expected, result)
	}
}

func TestPartB(t *testing.T) {
	expected := 6
	result := PartB(data)

	if expected != result {
		t.Fatalf("\nExpected = %d\nResult = %d\n", expected, result)
	}
}
