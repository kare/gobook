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
	if len(expectedOutput) == len(realOutput) {
		for i := range expectedOutput {
			if expectedOutput[i] != realOutput[i] {
				break
			}
		}
		return
	}
	t.Fatalf("UniqueInts() failed to produce expected output")

}
