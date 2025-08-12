package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func day11Calculate() {
	dat, err := os.ReadFile("inputs/day11.txt")
	if err != nil {
		log.Fatal(err)
	}
	beginingState := strings.Fields(string(dat))
	var currentState []string
	for i := 0; i < 75; i++ {
		for _, stone := range beginingState {
			currentState = append(currentState, calculateStone(stone)...)
		}
		beginingState = currentState
		currentState = make([]string, 0)
		print(i)
	}
	fmt.Println(len(beginingState))
}

func calculateStone(stone string) []string {
	out := make([]string, 0)
	if len(stone)%2 == 0 {
		right, _ := strconv.Atoi(stone[:len(stone)/2])
		left, _ := strconv.Atoi(stone[len(stone)/2:])
		out = append(out, strconv.Itoa(right))
		out = append(out, strconv.Itoa(left))

	} else if stone == "0" {
		out = append(out, "1")
	} else {
		stoneInt, _ := strconv.Atoi(stone)
		out = append(out, strconv.Itoa(stoneInt*2024))
	}
	return out
}
