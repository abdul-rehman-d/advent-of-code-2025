package day3

import (
	"testing"
)

const (
	data = `987654321111111
811111111111119
234234234234278
818181911112111`
)

func TestPartA(t *testing.T) {
	expected := 357
	result := PartA(data)

	if expected != result {
		t.Fatalf("\nExpected = %d\nResult = %d\n", expected, result)
	}
}

func TestPartB(t *testing.T) {
	expected := 3121910778619
	result := PartB(data)

	if expected != result {
		t.Fatalf("\nExpected = %d\nResult = %d\n", expected, result)
	}
}
