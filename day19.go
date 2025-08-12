package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func day19Calculate() {
	file, err := os.Open("inputs/day19.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	counter := 0
	readAvailablePatterns := false
	availablePatterns := make([]string, 0)
	pattternsToProduce := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if !readAvailablePatterns {
			availablePatterns = append(availablePatterns, strings.Split(line, ", ")...)
			readAvailablePatterns = true
			continue
		}
		if line == "" {
			continue
		}
		pattternsToProduce = append(pattternsToProduce, line)
		fmt.Println(line)
	}
	sort.Slice(availablePatterns, func(s1, s2 int) bool {
		l1 := len(availablePatterns[s1])
		l2 := len(availablePatterns[s2])
		if l1 == l2 {
			return availablePatterns[s1] > availablePatterns[s2]
		}
		return l1 > l2
	})
	for _, pattern := range pattternsToProduce {
		if isPatternPossible(pattern, availablePatterns) {
			counter++
		}
	}
	fmt.Println(counter)
}

func isPatternPossible(pattern string, patterns []string) bool {
	for {
		lenBefore := len(pattern)
		for _, p := range patterns {
			idx := strings.Index(pattern, p)
			if idx != -1 {
				pattern = pattern[:idx] + pattern[idx+len(p):]
				continue
			}
		}
		lenAfter := len(pattern)
		if lenAfter == lenBefore && lenAfter > 0 {
			return false
		}
		if lenAfter == 0 {
			return true
		}
	}
}
