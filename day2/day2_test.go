package day2

import (
	"testing"
)

const (
	data = "11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124"
)

func TestPartA(t *testing.T) {
	expected := 1227775554
	result := PartA(data)

	if expected != result {
		t.Fatalf("\nExpected = %d\nResult = %d\n", expected, result)
	}
}

func TestPartB(t *testing.T) {
	expected := 4174379265
	result := PartB(data)

	if expected != result {
		t.Fatalf("\nExpected = %d\nResult = %d\n", expected, result)
	}
}

func TestPartBValidater(t *testing.T) {
	shouldFailTestCases := []int{
		11,
		22,
		99,
		111,
		999,
		1010,
		1188511885,
		222222,
		446446,
		38593859,
		565656,
		824824824,
		2121212121,
	}
	shouldPassTestCases := []int{
		2121212118,
		2121212119,
		2121212120,
		2121212122,
		2121212123,
		2121212124,
	}

	for _, input := range shouldFailTestCases {
		if isValid2(input) {
			t.Fatal("should return false", input)
		}
	}
	for _, input := range shouldPassTestCases {
		if !isValid2(input) {
			t.Fatal("should return true", input)
		}
	}
}
