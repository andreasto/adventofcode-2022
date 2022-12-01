package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	inputs := []string{}

	for scanner.Scan() {
		inputs = append(inputs, scanner.Text())
	}

	e := ParseCals(inputs)
	highetsCal := CalculateHigestCalCount(e)
	top3 := CalculateTop3(e)

	fmt.Printf("The elf with highest calcount has total of %v cals \n", highetsCal)
	fmt.Printf("Top3 values combined is %v cals \n", top3)
}

func ParseCals(input []string) []Elf {
	elfs := []Elf{}
	e := Elf{}
	for _, in := range input {
		convertedCal, _ := strconv.Atoi(in)

		elfs = append(elfs, e)
		e = Elf{}
		e.calItems = append(e.calItems, convertedCal)
	}

	elfs = append(elfs, e)

	return elfs
}

func CalculateTop3(elfs []Elf) int {
	totalForAllElfs := []int{}

	for _, e := range elfs {
		var total int
		for _, c := range e.calItems {
			total += c
		}
		totalForAllElfs = append(totalForAllElfs, total)
	}

	sort.Slice(totalForAllElfs, func(i, j int) bool {
		return totalForAllElfs[i] > totalForAllElfs[j]
	})

	return totalForAllElfs[0] + totalForAllElfs[1] + totalForAllElfs[2]
}

func CalculateHigestCalCount(elfs []Elf) int {
	highest := 0
	for _, e := range elfs {
		currentElf := 0
		for _, c := range e.calItems {
			currentElf += c
		}

		if currentElf > highest {
			highest = currentElf
		}
	}

	return highest
}

type Elf struct {
	calItems []int
}
