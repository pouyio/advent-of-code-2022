package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isContained(innerStart, innerEnd, outerStart, outerEnd int) bool {
	return innerStart >= outerStart && innerEnd <= outerEnd
}

func isOverlap(start1, end1, start2, end2 int) bool {
	return start1 <= start2 && end1 >= start2
}

func part1() {
	// Open the file for reading
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// Create a new Scanner for the file
	scanner := bufio.NewScanner(file)

	totalSum := 0
	// Use the Scanner to read each line from the file
	for scanner.Scan() {

		line := scanner.Text()
		sections := strings.Split(line, ",")
		range1 := strings.Split(sections[0], "-")
		range2 := strings.Split(sections[1], "-")
		range1Start, _ := strconv.Atoi(range1[0])
		range1End, _ := strconv.Atoi(range1[1])
		range2Start, _ := strconv.Atoi(range2[0])
		range2End, _ := strconv.Atoi(range2[1])
		if isContained(range1Start, range1End, range2Start, range2End) || isContained(range2Start, range2End, range1Start, range1End) {
			totalSum++
		}

	}

	fmt.Println("Sum of complete overlap sections is:", totalSum)
}

func part2() {
	// Open the file for reading
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// Create a new Scanner for the file
	scanner := bufio.NewScanner(file)

	totalSum := 0
	// Use the Scanner to read each line from the file
	for scanner.Scan() {

		line := scanner.Text()
		sections := strings.Split(line, ",")
		range1 := strings.Split(sections[0], "-")
		range2 := strings.Split(sections[1], "-")
		range1Start, _ := strconv.Atoi(range1[0])
		range1End, _ := strconv.Atoi(range1[1])
		range2Start, _ := strconv.Atoi(range2[0])
		range2End, _ := strconv.Atoi(range2[1])
		if isOverlap(range1Start, range1End, range2Start, range2End) || isOverlap(range2Start, range2End, range1Start, range1End) {
			totalSum++
		}

	}

	fmt.Println("Sum of complete overlap sections is:", totalSum)
}

func main() {
	part1()
	part2()
}
