package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

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

	biggestSum := 0
	partialSum := 0
	// Use the Scanner to read each line from the file
	for scanner.Scan() {

		caloriesString := scanner.Text()
		calories, _ := strconv.Atoi(caloriesString)

		if len(caloriesString) == 0 {
			if partialSum > biggestSum {
				biggestSum = partialSum
			}
			partialSum = 0
		}
		partialSum += calories
	}

	fmt.Println("Biggest calories sum is:", biggestSum)
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

	var biggestSums []int
	partialSum := 0
	less := func(i, j int) bool {
		return biggestSums[i] > biggestSums[j]
	}
	// Use the Scanner to read each line from the file
	for scanner.Scan() {

		caloriesString := scanner.Text()
		calories, _ := strconv.Atoi(caloriesString)

		if len(caloriesString) == 0 {
			if len(biggestSums) < 3 {
				biggestSums = append(biggestSums, partialSum)
			} else {
				sort.Slice(biggestSums, less)
				for _, val := range biggestSums {
					if partialSum > val {
						biggestSums[2] = partialSum
						break
					}
				}
			}
			partialSum = 0
		}
		partialSum += calories
	}

	totalBiggestSums := 0
	for _, val := range biggestSums {
		totalBiggestSums += val
	}
	fmt.Println("3 biggest calories sum is:", totalBiggestSums)
}

func main() {
	part1()
	part2()
}
