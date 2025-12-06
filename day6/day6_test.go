package day6

import (
	"testing"
)

const (
	data = `123 328  51 64 
 45 64  387 23 
  6 98  215 314
*   +   *   +  `
)

func TestPartA(t *testing.T) {
	expected := 4277556
	result := PartA(data)

	if expected != result {
		t.Fatalf("\nExpected = %d\nResult = %d\n", expected, result)
	}
}

func TestPartB(t *testing.T) {
	expected := 3263827
	result := PartB(data)

	if expected != result {
		t.Fatalf("\nExpected = %d\nResult = %d\n", expected, result)
	}
}
