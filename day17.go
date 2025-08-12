package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type SmallComputer struct {
	A, B, C uint64
	Program []uint64
	Out     []uint64
}

func initComputer(puzzle []string) *SmallComputer {
	var res = SmallComputer{
		A:       0,
		B:       0,
		C:       0,
		Program: nil,
		Out:     nil,
	}
	reR := regexp.MustCompile(`Register ([A|B|C]): (\d+)`)
	reP := regexp.MustCompile(`\d`)
	for _, line := range puzzle {
		if strings.Contains(line, "Program") {
			match := reP.FindAllStringSubmatch(line, -1)
			for m := 0; m < len(match); m++ {
				instruction, _ := strconv.ParseUint(match[m][0], 10, 64)
				res.Program = append(res.Program, instruction)

			}
		} else if strings.Contains(line, "Register") {
			match := reR.FindAllStringSubmatch(line, -1)

			register, _ := strconv.Atoi(match[0][2])
			if match[0][1] == "A" {
				res.A = uint64(register)
			} else if match[0][1] == "B" {
				res.B = uint64(register)
			} else if match[0][1] == "C" {
				res.C = uint64(register)
			}
		}
	}
	return &res
}

func Run(a, b, c uint64, program []uint64) []uint64 {
	var instruction uint64
	var param uint64
	out := []uint64{}
	// fmt.Printf("A: %d, B: %d, C: %d\n", a, b, c)
	for pointer := uint64(0); pointer < uint64(len(program)); pointer += 2 {
		instruction, param = program[pointer], program[pointer+1]
		// fmt.Printf("Executing: instruction=%d, param=%d, a=%d, b=%d, c=%d\n", instruction, param, a, b, c)
		combo := param
		switch param {
		case 4:
			combo = a
		case 5:
			combo = b
		case 6:
			combo = c
		}
		switch instruction {
		case 0:
			a >>= combo
		case 1:
			b ^= param
		case 2:
			b = combo % 8
		case 3:
			if a != 0 {
				pointer = param - 2
			}
		case 4:
			b ^= c
		case 5:
			out = append(out, combo%8)
		case 6:
			b = a >> combo
		case 7:
			c = a >> combo
		}
		// fmt.Printf("After execution: a=%d, b=%d, c=%d, out=%v\n", a, b, c, out)
	}
	return out
}

func readFile(fname string) []string {
	var lines []string
	file, err := os.Open(fname)
	if err != nil {
		fmt.Println(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	return lines
}

func matchesProgram(output []uint64, expected []uint64) bool {
	if len(output) != len(expected) {
		return false
	}
	for i := range output {
		if output[i] != expected[i] {
			return false
		}
	}
	return true
}

func SolveDay17(puzzle []string, part2 bool) string {
	var res string
	comp := initComputer(puzzle)
	fmt.Printf("Initial state: A=%d, B=%d, C=%d\n", comp.A, comp.B, comp.C)
	fmt.Printf("Program: %v\n", comp.Program)

	if !part2 {
		out := Run(comp.A, comp.B, comp.C, comp.Program)
		for _, num := range out {
			res += strconv.FormatUint(num, 10) + ","
		}
	} else {
		type QueueItem struct {
			a uint64
			n int
		}

		queue := list.New()
		queue.PushBack(QueueItem{a: 0, n: 1})

		for queue.Len() > 0 {
			item := queue.Remove(queue.Front()).(QueueItem)
			a, n := item.a, item.n

			if n > len(comp.Program) { // Base case
				return strconv.FormatUint(a, 10)
			}

			for i := uint64(0); i < 8; i++ {
				a2 := (a << 3) | i
				out := Run(a2, 0, 0, comp.Program)
				target := comp.Program[len(comp.Program)-n:]

				// Save correct partial solutions
				if matchesProgram(out, target) {
					queue.PushBack(QueueItem{a: a2, n: n + 1})
				}
			}
		}

		return "0" // Return 0 if no solution found
	}

	return res
}

func calculateDay17() {

	timeStart := time.Now()
	INPUT := "inputs/day17.txt"
	// INPUT := "sample.txt"

	puzzle := readFile(INPUT)

	sum_part1 := SolveDay17(puzzle, false)

	fmt.Println("Part1:", sum_part1)

	sum_part2 := SolveDay17(puzzle, true)
	fmt.Println("Part2:", sum_part2)

	fmt.Printf("Time: %.2fms\n", float64(time.Since(timeStart).Microseconds())/1000)
}
