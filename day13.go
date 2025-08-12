package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type Machine struct {
	Ax int
	Ay int
	Bx int
	By int
	Px int
	Py int
}

func parseInputDay13() []Machine {
	dat, err := os.ReadFile("inputs/day13.txt")
	if err != nil {
		log.Fatal(err)
	}
	ret := make([]Machine, 0)
	content := strings.TrimFunc(string(dat), unicode.IsSpace)
	machineStrings := strings.Split(content, "\n\n")
	for _, machineString := range machineStrings {
		lines := strings.Split(machineString, "\n")
		componentsA := strings.Split(strings.Split(lines[0], ": ")[1], ", ")
		componentsB := strings.Split(strings.Split(lines[1], ": ")[1], ", ")
		componentsP := strings.Split(strings.Split(lines[2], ": ")[1], ", ")
		Ax, _ := strconv.Atoi(componentsA[0][2:])
		Ay, _ := strconv.Atoi(componentsA[1][2:])
		Bx, _ := strconv.Atoi(componentsB[0][2:])
		By, _ := strconv.Atoi(componentsB[1][2:])
		Px, _ := strconv.Atoi(componentsP[0][2:])
		Py, _ := strconv.Atoi(componentsP[1][2:])
		mach := Machine{Ax, Ay, Bx, By, Px, Py}
		ret = append(ret, mach)
	}
	return ret
}

func calculateDay13(machines []Machine) {
	sum := 0
	sum2 := 0
	for _, machine := range machines {
		sum += solve(machine)
		machine.Px += 10000000000000
		machine.Py += 10000000000000
		sum2 += solve(machine)
	}
	fmt.Println(sum)
	fmt.Println(sum2)
}

func solve(m Machine) int {

	if areSimilar(m) {
		if areSimilar(m) {
			log.Println("All vectors are similar for machine:", m, "Finding cheapest solution is not yet implemented in this case.")
			return 0
		}
	} else if areSimilar(m) {
		if m.Px%m.Bx == 0 {
			return m.Px / m.Bx
		} else {
			return 0
		}
	} else if areSimilar(m) {
		return 0
	}

	bNumerator := m.Py*m.Ax - m.Px*m.Ay
	bDenominator := m.By*m.Ax - m.Bx*m.Ay

	if bDenominator != 0 && bNumerator%bDenominator == 0 {
		b := bNumerator / bDenominator
		aNumerator := m.Px - b*m.Bx
		if m.Ax != 0 && aNumerator%m.Ax == 0 {
			a := aNumerator / m.Ax
			return 3*a + b
		}
	}

	if m.Ax == 0 && m.Px%m.Bx == 0 {
		b := m.Px / m.Bx
		aNumerator := m.Py - b*m.By
		if m.Ay != 0 && aNumerator%m.Ay == 0 {
			a := aNumerator / m.Ay
			return 3*a + b
		}
	}

	return 0
}

func areSimilar(m Machine) bool {
	return ((m.Ax / m.Ay) == (m.Bx / m.By)) && (m.Ax%m.Ay)*m.By == (m.Bx%m.By)*m.Ay
}
