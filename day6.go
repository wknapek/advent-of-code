package main

import (
	"bufio"
	"fmt"
	"io"
	"iter"
	"maps"
	"os"
	"slices"
	"strings"
	"sync"
	"sync/atomic"
)

const OBSTACLE, VISITED, LEFT, UP, RIGHT, DOWN = "#", "X", "<", "^", ">", "v"

type GuardPath [][]string

type Pose struct {
	i, j int
	dir  string
}

type Position struct {
	i, j int
}

func Solve(inputPath string) {
	gp := GuardPath(ReadFileToGrid(inputPath, ""))
	startPose := FindStartPose(&gp)
	prevPose := startPose

	var wg sync.WaitGroup
	var obstacleCount int32

	visitedPosition := make(map[Position]struct{})

	for pose := range gp.Steps(startPose) {
		pos := Position{pose.i, pose.j}

		// Part 02 - Do not put an obstacle in the same position twice
		if _, visited := visitedPosition[pos]; visited {
			continue
		}

		visitedPosition[pos] = struct{}{}

		// Part 02 - Do not put an obstacle on the start position
		if pose == startPose {
			continue
		}

		blockedPath := gp.DeepCopy()
		blockedPath[pose.i][pose.j] = OBSTACLE

		wg.Add(1)
		go func(from Pose) {
			defer wg.Done()
			if IsLoop(&blockedPath, from) {
				atomic.AddInt32(&obstacleCount, 1)
			}
		}(prevPose)

		prevPose = pose
	}
	count := IterLength(maps.Keys(visitedPosition))
	fmt.Printf("Part 01: %v (unique positions visited)\n", count)

	wg.Wait()
	fmt.Printf("Part 02: %v (number of obstructions that create a loop)\n", obstacleCount)
}

func IsLoop(gp *GuardPath, startPose Pose) bool {
	visitedPose := make(map[Pose]struct{})
	for pose := range gp.Steps(startPose) {
		if _, visited := visitedPose[pose]; visited {
			return true
		}
		visitedPose[pose] = struct{}{}
	}
	return false
}

func (gp *GuardPath) DeepCopy() GuardPath {
	cpy := make(GuardPath, len(*gp))
	for i, row := range *gp {
		cpy[i] = make([]string, len(row))
		copy(cpy[i], (*gp)[i])
	}
	return cpy
}

func (gp *GuardPath) Steps(startPose Pose) iter.Seq[Pose] {
	iMax := len(*gp) - 1
	jMax := len((*gp)[0]) - 1

	pose := startPose

	return func(yield func(Pose) bool) {
		di, dj := 0, 0
		for {
			if !yield(pose) {
				return
			}
			for {
				di, dj = ComputeNextStep(pose.dir)
				if IsOutOfBounds(pose.i+di, pose.j+dj, iMax, jMax) {
					return
				}
				if (*gp)[pose.i+di][pose.j+dj] != OBSTACLE {
					break
				}
				pose.dir = GetNextDirection(pose.dir)
			}
			pose.i += di
			pose.j += dj
		}
	}
}

func (gp GuardPath) String() string {
	s := ""
	for _, row := range gp {
		s += strings.Join(row, "")
		s += "\n"
	}
	return s
}

func FindStartPose(gp *GuardPath) Pose {
	directions := []string{UP, RIGHT, LEFT, DOWN}
	for i, row := range *gp {
		for _, dir := range directions {
			if j := slices.Index(row, dir); j > 0 {
				return Pose{i, j, (*gp)[i][j]}
			}
		}
	}
	panic("failed to find start pose")
}

func IsOutOfBounds(i, j, iMax, jMax int) bool {
	return i < 0 || i > iMax || j < 0 || j > jMax
}

func GetNextDirection(dir string) string {
	switch dir {
	case LEFT:
		return UP
	case UP:
		return RIGHT
	case RIGHT:
		return DOWN
	default:
		return LEFT
	}
}

func ComputeNextStep(dir string) (di int, dj int) {
	switch dir {
	case LEFT:
		return 0, -1
	case UP:
		return -1, 0
	case RIGHT:
		return 0, 1
	default:
		return 1, 0
	}
}

func IterLength[V any](s iter.Seq[V]) int {
	count := 0
	for range s {
		count++
	}
	return count
}

func ReadFileToGrid(inputPath, delimiter string) [][]string {
	file, err := os.Open(inputPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	content, err := readToGrid(file, delimiter)
	if err != nil {
		panic(err)
	}
	return content
}

func readToGrid(reader io.Reader, delimiter string) ([][]string, error) {
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
