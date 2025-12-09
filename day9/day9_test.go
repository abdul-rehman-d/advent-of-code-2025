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
	expected := 0
	result := PartB(data)

	if expected != result {
		t.Fatalf("\nExpected = %d\nResult = %d\n", expected, result)
	}
}
