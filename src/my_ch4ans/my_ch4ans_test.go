package main

import (
	"log"
	"testing"
)

func TestUniqueInts(t *testing.T) {
	log.SetFlags(0)
	log.Println("TEST UniqueInts()")
	input := []int{9, 1, 9, 5, 4, 4, 2, 1, 5, 4, 8, 8, 4, 3, 6, 9, 5, 7, 5}
	expectedOutput := []int{9, 1, 5, 4, 2, 8, 3, 6, 7}
	realOutput := UniqueInts(input)
	if !equalIntSlices(realOutput, expectedOutput) {
		t.Fatalf("UniqueInts() failed to produce expected output")
	}

}

func TestFlatten(t *testing.T) {
	log.SetFlags(0)
	log.Println("TEST Flatten()")
	irregularMatrix := [][]int{{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11},
		{12, 13, 14, 15},
		{16, 17, 18, 19, 20}}

	expectedOutput := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
		11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	realOutput := Flatten(irregularMatrix)
	if !equalIntSlices(realOutput, expectedOutput) {
		t.Fatalf("Flatten() failed to produce expected output")
	}
}

func equalIntSlices(slice1 []int, slice2 []int) bool {
	if len(slice1) == len(slice2) {
		for i := range slice1 {
			if slice1[i] != slice2[i] {
				return false
			}
		}
		return true
	}
	return false
}
