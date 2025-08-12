package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func toCode(input string) []int {
	chars := strings.Split(input, "")
	return StringsToNumbers(chars)
}

func day9Part1(input string) int {
	code := toCode(input)

	blocksLen := 0
	for i, val := range code {
		if i%2 == 0 {
			blocksLen += val
		}
	}

	blocks := make([]int, blocksLen)
	i := 0
	id := len(code) - 1

blocksLoop:
	for j := 0; j < len(code); j++ {
		if j%2 == 0 {
			for k := code[j]; k > 0; k-- {
				blocks[i] = j / 2
				i++
			}
		} else {
			emptyFields := code[j]
			for k := 0; k < emptyFields; k++ {
				if code[id] == 0 {
					id -= 2
				}
				if i >= blocksLen {
					break blocksLoop
				}
				blocks[i] = id / 2
				code[id]--
				i++
			}
		}
	}

	checksum := 0
	for i, val := range blocks {
		checksum += i * val
	}

	return checksum
}

func day9Part2(input string) int {
	code := toCode(input)

	blocksLen := 0
	for _, val := range code {
		blocksLen += val
	}

	blocks := make([]int, blocksLen)
	recs := make(map[int]int)
	i := 0

	for j := 0; j < len(code); j++ {
		recs[j] = i
		if j%2 == 0 {
			for k := code[j]; k > 0; k-- {
				blocks[i] = j / 2
				i++
			}
		} else {
			i += code[j]
		}
	}

	id := 1
	for j := len(code) - 1; j >= 0; j -= 2 {
		for k := id; k <= j; k += 2 {
			if code[k] >= code[j] {
				i := recs[k]
				for l := 0; l < code[j]; l++ {
					blocks[i] = j / 2
					i++
				}

				i = recs[j]
				for l := 0; l < code[j]; l++ {
					blocks[i] = 0
					i++
				}

				code[k] -= code[j]
				recs[k] += code[j]
				code[j] = 0

				if code[id] == 0 {
					id += 2
				}
				break
			}
		}
	}

	checksum := 0
	for i, val := range blocks {
		checksum += i * val
	}

	return checksum
}

func StringsToNumbers(strings []string) []int {
	nums := make([]int, len(strings))
	for i, str := range strings {
		num, _ := strconv.Atoi(str)
		nums[i] = num
	}
	return nums
}

func day9Calculate() {
	data, err := os.ReadFile("inputs/day9.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(day9Part1(string(data)))
	fmt.Println(day9Part2(string(data)))
}
