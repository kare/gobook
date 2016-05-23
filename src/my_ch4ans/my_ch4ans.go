package main

import (
	"sort"
	"strings"
)

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

func ParseIni(input []string) map[string]map[string]string {
	result := make(map[string]map[string]string)
	currentMap := make(map[string]string)
	for i := len(input) - 1; i >= 0; i-- {
		if input[i] == "" || input[i][0] == ';' {
			continue
		}
		if header, ok := getHeader(input[i]); ok {
			// Store previously built up map using header as key
			if len(currentMap) > 0 {
				result[header] = currentMap
				currentMap = make(map[string]string)
			}
		} else {
			if k, v, ok := getKeyValuePair(input[i]); ok {
				// Add KV pair to current map
				currentMap[k] = v
			}
		}
	}
	return result
}

func getHeader(line string) (string, bool) {
	if line[0] == '[' && line[len(line)-1] == ']' {
		return line[1 : len(line)-1], true
	} else {
		return "", false
	}
}

func getKeyValuePair(line string) (string, string, bool) {
	if separated := strings.Split(line, "="); len(separated) == 2 {
		return separated[0], separated[1], true
	}
	return "", "", false
}

func PrintIni(input map[string]map[string]string) string {
	var result string
	topLevelKeys := make([]string, 0, len(input))
	for topKey := range input {
		topLevelKeys = append(topLevelKeys, topKey)
	}
	sort.Strings(topLevelKeys)
	for _, topKey := range topLevelKeys {
		result += "[" + topKey + "]\n"
		secondLevelKeys := make([]string, 0, len(input[topKey]))
		for secondLevelKey := range input[topKey] {
			secondLevelKeys = append(secondLevelKeys, secondLevelKey)
		}
		sort.Strings(secondLevelKeys)
		for _, secondLevelKey := range secondLevelKeys {
			result += secondLevelKey + "=" + input[topKey][secondLevelKey] + "\n"
		}
		result += "\n"
	}
	return result
}
