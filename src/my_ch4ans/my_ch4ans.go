package main

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

func Flatten(input [][]int) []int {
	flattened := []int{}
	for i := range input {
		for j := range input[i] {
			flattened = append(flattened, input[i][j])
		}
	}
	return flattened
}

func Make2D(input []int, columns int) [][]int {
	inputLength := len(input)
	rows := (inputLength + columns - 1) / columns
	result := make([][]int, 0)
	column := []int{}
	for i := 0; i < rows*columns; i++ {
		if i+1 <= inputLength {
			column = append(column, input[i])
		} else {
			column = append(column, 0)
		}
		if i%columns == columns-1 {
			result = append(result, column)
			column = make([]int, 0)
		}
	}
	return result
}
