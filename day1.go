package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func prepareInput() ([]int, []int) {
	left := make([]int, 0)
	right := make([]int, 0)
	file, err := os.Open("inputs/day1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		splitedLine := strings.Fields(line)
		leftLine, _ := strconv.Atoi(splitedLine[0])
		left = append(left, leftLine)
		rightLine, _ := strconv.Atoi(splitedLine[1])
		right = append(right, rightLine)
	}
	return left, right
}

func day1Calculate() int {
	left, right := prepareInput()
	sort.Ints(left)
	sort.Ints(right)
	sum := 0
	for i := 0; i < len(left); i++ {
		sum += int(math.Abs(float64(left[i] - right[i])))
	}
	return sum
}

func day1CalculateB() int {
	left, right := prepareInput()
	sum := 0
	for i := 0; i < len(left); i++ {
		multiplicator := 0
		for j := 0; j < len(right); j++ {
			if left[i] == right[j] {
				multiplicator++
			}
		}
		sum += left[i] * multiplicator
	}
	return sum
}
