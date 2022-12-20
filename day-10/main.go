package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func calculateSumCycle(cycle, sum, register int) (int, int) {
	cycle++
	if (cycle+20)%40 == 0 {
		sum += (cycle * register)
	}
	return sum, cycle
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
	cycle := 1
	register := 1
	sum := 0
	// Use the Scanner to read each line from the file
	for scanner.Scan() {
		line := scanner.Text()
		command := strings.Split(line, " ")
		sum, cycle = calculateSumCycle(cycle, sum, register)

		if command[0] == "addx" {
			arg, _ := strconv.Atoi(command[1])
			register += arg
			sum, cycle = calculateSumCycle(cycle, sum, register)
		}

	}

	fmt.Println("Total sum of six signal strengths:", sum)
}

func addPixel(crt [][]string, cycle, register int) [][]string {
	r := int(math.Floor(float64(cycle) / 40))
	if math.Abs(float64(register-(cycle-(40*r)))) <= 1 {
		crt[r] = append(crt[r], "#")
	} else {
		crt[r] = append(crt[r], ".")
	}
	return crt
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
	cycle := 0
	register := 1
	crt := [][]string{{}, {}, {}, {}, {}, {}}
	// Use the Scanner to read each line from the file
	for scanner.Scan() {
		line := scanner.Text()
		command := strings.Split(line, " ")

		crt = addPixel(crt, cycle, register)
		cycle++

		if command[0] == "addx" {
			crt = addPixel(crt, cycle, register)
			cycle++
			arg, _ := strconv.Atoi(command[1])
			register += arg
		}

	}

	for _, line := range crt {
		fmt.Println(line)
	}

}

func main() {
	part1()
	part2()
}
