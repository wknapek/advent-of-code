package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func day3() {
	file, err := os.Open("inputs/day3.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		err = file.Close()
		if err != nil {

		}
	}(file)

	resA := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		resA += multiplicate(line)
		multiplicate2(line)
	}
	fmt.Println(fmt.Sprintf("Part One: %d", resA))
}

func multiplicate(line string) int {
	res := 0
	re := regexp.MustCompile("mul\\(\\d*,\\d*\\)")
	matches := re.FindAllStringSubmatch(line, -1)
	for _, match := range matches {
		numbers := match[0][4 : len(match[0])-1]
		digits := strings.Split(numbers, ",")
		mul1, _ := strconv.Atoi(digits[0])
		mul2, _ := strconv.Atoi(digits[1])
		fmt.Println(numbers)
		res += mul1 * mul2
	}
	return res
}

func multiplicate2(line string) {
	input, _ := os.ReadFile("inputs/day3.txt")
	inputStr := string(input)

	reSections := regexp.MustCompile(`do\(\).*?|don't\(\).*?`)

	sections := reSections.FindAllStringIndex(inputStr, -1)
	nextStartIndex := 0
	enabledRead := true
	var result int
	for _, section := range sections {
		start, end := section[0], section[1]
		if enabledRead {
			result += multiplicate(inputStr[nextStartIndex:start])
		}
		nextStartIndex = end
		action := inputStr[start:end]
		switch action {
		case "do()":
			if enabledRead == false {
				enabledRead = true
			}
		case "don't()":
			if enabledRead == true {
				enabledRead = false
			}
		}
	}
	if enabledRead {
		result += multiplicate(inputStr[nextStartIndex:])
	}
	fmt.Println(result)
}
