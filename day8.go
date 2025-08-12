package main

import (
	"bufio"
	"fmt"
	"io"
	"maps"
	"os"
	"strings"
)

func day8Part1(inputPath string) {
	antennaGrid := ReadFileTo2D(inputPath, "")
	signalGrid := make(map[Position]struct{})

	iMax := len(antennaGrid) - 1
	jMax := len(antennaGrid[0]) - 1

	// Find tower
	for i, row := range antennaGrid {
		for j, antenna := range row {
			if antenna == "." {
				continue
			}

			// Find matching towers
			for mi, mrow := range antennaGrid {
				for mj, mantenna := range mrow {
					if antenna == "." || antenna != mantenna || (mi == i && mj == j) {
						continue
					}

					di, dj := mi-i, mj-j
					signalPos := Position{i - di, j - dj}

					if !IsOutOfBounds2D(signalPos.i, signalPos.j, iMax, jMax) {
						signalGrid[signalPos] = struct{}{}
					}

				}
			}
		}
	}

	signalPositions := maps.Keys(signalGrid)
	count := IterLength(signalPositions)
	fmt.Printf("Part 01: %v (number of antinodes)\n", count)
}

func day8Part02(inputPath string) {
	antennaGrid := ReadFileTo2D(inputPath, "")
	signalGrid := make(map[Position]struct{})

	iMax := len(antennaGrid) - 1
	jMax := len(antennaGrid[0]) - 1

	// Find tower
	for i, row := range antennaGrid {
		for j, antenna := range row {
			if antenna == "." {
				continue
			}

			// Find matching towers
			for mi, mrow := range antennaGrid {
				for mj, mantenna := range mrow {
					if antenna == "." || antenna != mantenna || (mi == i && mj == j) {
						continue
					}

					signalPos := Position{mi, mj} // The first signal is on the antenna itself
					di, dj := mi-i, mj-j
					n := 0
					for !IsOutOfBounds2D(signalPos.i, signalPos.j, iMax, jMax) {
						signalGrid[signalPos] = struct{}{}
						signalPos.i = i - n*di
						signalPos.j = j - n*dj
						n++
					}
				}
			}
		}
	}

	signalPositions := maps.Keys(signalGrid)
	count := IterLength(signalPositions)
	fmt.Printf("Part 02: %v (number of antinodes considering harmonics)\n", count)
}

func readTo2D(reader io.Reader, delimiter string) ([][]string, error) {
	grid := [][]string{}
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		tokens := strings.Split(scanner.Text(), delimiter)
		grid = append(grid, tokens)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return grid, nil
}

func ReadFileTo2D(inputPath, delimiter string) [][]string {
	file, err := os.Open(inputPath)
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err = file.Close()
		if err != nil {

		}
	}(file)

	content, err := readTo2D(file, delimiter)
	if err != nil {
		panic(err)
	}
	return content
}

func IsOutOfBounds2D(i, j, iMax, jMax int) bool {
	return i < 0 || i > iMax || j < 0 || j > jMax
}
