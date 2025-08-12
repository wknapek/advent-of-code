package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type CoordinateDay14 struct {
	x, y int
}

type Robot struct {
	pos CoordinateDay14 // position
	vel CoordinateDay14 // velocity
}

func day14() {
	dat, err := os.ReadFile("inputs/day14.txt")
	if err != nil {
		log.Fatal(err)
	}
	robots := parseInputDay14(string(dat))

	fmt.Println(simulateAndCalculate(robots))

	fmt.Println(findEasterEgg(robots))
}

func parseInputDay14(input string) []Robot {
	lines := strings.Split(input, "\n")
	robots := make([]Robot, len(lines))

	for i, line := range lines {
		parts := strings.Split(line, " ")
		// please dont judge me for the following three lines... it works...
		robots[i] = Robot{
			pos: CoordinateDay14{AtoiNoErr(strings.Split(strings.Split(parts[0], ",")[0], "=")[1]), AtoiNoErr(strings.Split(parts[0], ",")[1])},
			vel: CoordinateDay14{AtoiNoErr(strings.Split(strings.Split(parts[1], ",")[0], "=")[1]), AtoiNoErr(strings.Split(parts[1], ",")[1])},
		}
	}

	return robots
}

func simulateAndCalculate(robots []Robot) int {
	// simulate for 100 seconds
	for s := 0; s < 100; s++ {
		moveRobots(robots)
	}

	// count robots in each quadrant
	q1, q2, q3, q4 := 0, 0, 0, 0
	midX, midY := 101/2, 103/2

	for _, r := range robots {
		// skip robots exactly on the middle lines
		if r.pos.x == midX && r.pos.y == midY {
			continue
		}

		// count robots in each quad
		if r.pos.x < midX && r.pos.y < midY {
			q1++ // Top-left
		} else if r.pos.x > midX && r.pos.y < midY {
			q2++ // Top-right
		} else if r.pos.x < midX && r.pos.y > midY {
			q3++ // Bottom-left
		} else if r.pos.x > midX && r.pos.y > midY {
			q4++ // Bottom-right
		}
	}

	return q1 * q2 * q3 * q4
}

func moveRobots(robots []Robot) {
	for i := range robots {
		// Update position
		robots[i].pos.x += robots[i].vel.x
		robots[i].pos.y += robots[i].vel.y

		// Handle wrapping
		robots[i].pos.x = ((robots[i].pos.x % 101) + 101) % 101
		robots[i].pos.y = ((robots[i].pos.y % 103) + 103) % 103
	}
}

func findEasterEgg(robots []Robot) int {
	// Try each second until we find the pattern
	for t := 1; t < 11000; t++ {
		moveRobots(robots)

		// Check for pattern
		if hasLongHorizontalLine(robots) {
			return t
		}
	}
	return -1
}

func hasLongHorizontalLine(robots []Robot) bool {
	// Count robots in each row
	rowCounts := make(map[int]map[int]bool)
	for i := 0; i < 103; i++ {
		rowCounts[i] = make(map[int]bool)
	}

	// Record robot positions by row
	for _, robot := range robots {
		rowCounts[robot.pos.y][robot.pos.x] = true
	}

	// Look for a row with many robots and long consecutive sequences
	for y := 0; y < 103; y++ {
		if len(rowCounts[y]) > 30 {
			consecutive := 0
			for x := 0; x < 100; x++ {
				if rowCounts[y][x] && rowCounts[y][x+1] {
					consecutive++
					if consecutive > 25 {
						return true
					}
				} else {
					consecutive = 0
				}
			}
		}
	}

	return false
}

func AtoiNoErr(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
