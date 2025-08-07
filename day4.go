package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func day4One() {
	lines := getLines("inputs/day4.txt")
	word := "XMAS"
	found := 0
	for x := 0; x < len(lines); x++ {
		for y := 0; y < len(lines[x]); y++ {
			found += checkEast(lines, x, y, word)
			found += checkWest(lines, x, y, word)
			found += checkNorth(lines, x, y, word)
			found += checkSouth(lines, x, y, word)
			found += checkNorthEast(lines, x, y, word)
			found += checkNorthWest(lines, x, y, word)
			found += checkSouthEast(lines, x, y, word)
			found += checkSouthWest(lines, x, y, word)
		}
	}
	fmt.Printf("Found %d occurrences\n", found)
}

func day4Two() {
	lines := getLines("inputs/day4.txt")
	word1 := "MAS"
	word2 := "SAM"
	found := 0
	for x := 0; x < len(lines); x++ {
		for y := 0; y < len(lines[x]); y++ {
			if checkSouthEast(lines, x, y, word1) == 1 {
				if checkSouthWest(lines, x, y+2, word1) == 1 || checkNorthEast(lines, x+2, y, word1) == 1 {
					found++
				}
			}
			if checkSouthEast(lines, x, y, word2) == 1 {
				if checkSouthWest(lines, x, y+2, word2) == 1 || checkNorthEast(lines, x+2, y, word2) == 1 {
					found++
				}
			}
		}
	}
	fmt.Printf("Found %d occurrences\n", found)
}

func getLines(path string) [][]string {
	scanner, file := GetFileScanner(path)
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)
	var lines [][]string
	for scanner.Scan() {
		line := scanner.Text()
		letters := strings.Split(line, "")
		lines = append(lines, letters)
	}
	return lines
}

func checkEast(lines [][]string, x, y int, word string) int {
	if lines[x][y] == string(word[0]) {
		if y+(len(word)-1) < len(lines[x]) {
			for i := 0; i < len(word); i++ {
				if lines[x][y+i] != string(word[i]) {
					return 0
				}
			}
			return 1
		}
	}
	return 0
}

func checkWest(lines [][]string, x, y int, word string) int {
	if lines[x][y] == string(word[0]) {
		if y-(len(word)-1) >= 0 {
			for i := 0; i < len(word); i++ {
				if lines[x][y-i] != string(word[i]) {
					return 0
				}
			}
			return 1
		}
	}
	return 0
}

func checkNorth(lines [][]string, x, y int, word string) int {
	if lines[x][y] == string(word[0]) {
		if x-(len(word)-1) >= 0 {
			for i := 0; i < len(word); i++ {
				if lines[x-i][y] != string(word[i]) {
					return 0
				}
			}
			return 1
		}
	}
	return 0
}

func checkSouth(lines [][]string, x, y int, word string) int {
	if lines[x][y] == string(word[0]) {
		if x+(len(word)-1) < len(lines) {
			for i := 0; i < len(word); i++ {
				if lines[x+i][y] != string(word[i]) {
					return 0
				}
			}
			return 1
		}
	}
	return 0
}

func checkNorthEast(lines [][]string, x, y int, word string) int {
	if lines[x][y] == string(word[0]) {
		if x-(len(word)-1) >= 0 && y+(len(word)-1) < len(lines[x]) {
			for i := 0; i < len(word); i++ {
				if lines[x-i][y+i] != string(word[i]) {
					return 0
				}
			}
			return 1
		}
	}
	return 0
}

func checkNorthWest(lines [][]string, x, y int, word string) int {
	if lines[x][y] == string(word[0]) {
		if x-(len(word)-1) >= 0 && y-(len(word)-1) >= 0 {
			for i := 0; i < len(word); i++ {
				if lines[x-i][y-i] != string(word[i]) {
					return 0
				}
			}
			return 1
		}
	}
	return 0
}

func checkSouthEast(lines [][]string, x, y int, word string) int {
	if lines[x][y] == string(word[0]) {
		if x+(len(word)-1) < len(lines) && y+(len(word)-1) < len(lines[x]) {
			for i := 0; i < len(word); i++ {
				if lines[x+i][y+i] != string(word[i]) {
					return 0
				}
			}
			return 1
		}
	}
	return 0
}

func checkSouthWest(lines [][]string, x, y int, word string) int {
	if lines[x][y] == string(word[0]) {
		if x+(len(word)-1) < len(lines) && y-(len(word)-1) >= 0 {
			for i := 0; i < len(word); i++ {
				if lines[x+i][y-i] != string(word[i]) {
					return 0
				}
			}
			return 1
		}
	}
	return 0
}

func GetFileScanner(path string) (*bufio.Scanner, *os.File) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("file open error:", err)
	}
	scanner := bufio.NewScanner(file)
	return scanner, file
}
