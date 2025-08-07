package main

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func FiveOne() {
	pairs, updates := parseInput("inputs/day5.txt")
	verifiedUpdates, _ := verifyUpdates(updates, pairs)
	middleNums := addMiddleNums(verifiedUpdates)
	fmt.Println(middleNums)
}

func FiveTwo() {
	pairs, updates := parseInput("inputs/day5.txt")
	_, unverifiedUpdates := verifyUpdates(updates, pairs)
	reordered := reorderUnverified(unverifiedUpdates, pairs)
	middleNums := addMiddleNums(reordered)
	fmt.Println(middleNums)
}

func addMiddleNums(updates [][]int) int {
	var total int
	for _, update := range updates {
		updateLength := len(update)
		midPoint := int(updateLength / 2)
		middleItem := update[midPoint]
		total += middleItem
	}
	return total
}

func reorderUnverified(updates [][]int, pairs [][]int) [][]int {
	var reordered [][]int
	for _, update := range updates {
	INNERLOOP:
		problemFound := false
		for index := 1; index < len(update); index++ {
			toCheckFirst := []int{update[index-1], update[index]}
			if !ArrayContainsDeep(pairs, toCheckFirst) {
				reflect.Swapper(update)(index-1, index)
				problemFound = true
			}
		}
		if problemFound {
			goto INNERLOOP
		}
		reordered = append(reordered, update)
	}
	return reordered
}

func verifyUpdates(updates, pairs [][]int) ([][]int, [][]int) {
	var verifiedUpdates [][]int
	var unverifiedUpdates [][]int
	for _, update := range updates {
		verified := true
		for index := range update {
			if index > len(update)-2 {
				break
			}
			toCheck := []int{update[index], update[index+1]}
			if !ArrayContainsDeep(pairs, toCheck) {
				verified = false
				unverifiedUpdates = append(unverifiedUpdates, update)
				break
			}
		}
		if verified {
			verifiedUpdates = append(verifiedUpdates, update)
		}
	}
	return verifiedUpdates, unverifiedUpdates
}

func parseInput(path string) ([][]int, [][]int) {
	scanner, file := GetFileScanner(path)
	defer file.Close()
	var pairs [][]int
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		pair := strings.Split(line, "|")
		firstNum, _ := strconv.Atoi(pair[0])
		secondNum, _ := strconv.Atoi(pair[1])
		pairs = append(pairs, []int{firstNum, secondNum})
	}
	var updates [][]int
	for scanner.Scan() {
		line := scanner.Text()
		update := strings.Split(line, ",")
		updateInts := ArrayStringToInt(update)
		updates = append(updates, updateInts)
	}
	return pairs, updates
}

func ArrayStringToInt(input []string) []int {
	var output []int
	for _, val := range input {
		intVal, _ := strconv.Atoi(val)
		output = append(output, intVal)
	}
	return output
}

func ArrayContainsDeep(nums [][]int, target []int) bool {
	for _, num := range nums {
		if reflect.DeepEqual(num, target) {
			return true
		}
	}
	return false
}
