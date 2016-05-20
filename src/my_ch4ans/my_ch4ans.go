package main

// import (
// 	"fmt"
// )

func UniqueInts(input []int) []int {
	uniques := []int{}
	counts := make(map[int]int)
	for i := range input {
		if _, found := counts[input[i]]; !found {
			counts[input[i]] = 1
			uniques = append(uniques, input[i])
		}
	}
	return uniques
}
