package day5

import (
	"testing"
)

const (
	data = `3-5
10-14
16-20
12-18

1
5
8
11
17
32`
)

func TestPartA(t *testing.T) {
	expected := 3
	result := PartA(data)

	if expected != result {
		t.Fatalf("\nExpected = %d\nResult = %d\n", expected, result)
	}
}

func TestPartB(t *testing.T) {
	expected := 14
	result := PartB(data)

	if expected != result {
		t.Fatalf("\nExpected = %d\nResult = %d\n", expected, result)
	}
}
