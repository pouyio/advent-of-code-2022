package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

	crateIteration := 0
	// Use the Scanner to read each line from the file
	var crates [][]string
	for scanner.Scan() {

		line := scanner.Text()

		// parse crates matrix
		if strings.Contains(line, "[") {
			crates = append(crates, getCrates(line))
			crateIteration++
		}

		// format matrix, each stack is a row
		if strings.Contains(line, "1   2   3") {
			crates = transposeAndReverse(crates)
			crates = removeEmpty(crates)
		}

		// move crates
		if strings.Contains(line, "move") {
			movements := getMovements(line)
			crates = moveCrates(crates, movements)
		}
	}

	fmt.Println("Crates on top of each stack are:", getTopCrates(crates))
}

func getTopCrates(crates [][]string) string {
	topCrates := ""
	for i := 0; i < len(crates); i++ {
		topCrates = topCrates + crates[i][len(crates[i])-1]
	}
	return topCrates
}

func removeEmpty(crates [][]string) [][]string {
	var result = make([][]string, len(crates))
	for i := range crates {
		result[i] = make([]string, 0)
		for j := range crates[i] {
			if crates[i][j] != " " {
				result[i] = append(result[i], crates[i][j])
			}
		}
	}
	return result
}

func moveCrates(crates [][]string, movements [3]int) [][]string {
	from := movements[1] - 1
	to := movements[2] - 1
	for i := 0; i < movements[0]; i++ {
		crates, crateToMove := removeCrate(crates, from)
		crates = setCrate(crates, to, crateToMove)
	}
	return crates
}

func moveCratesSameOrder(crates [][]string, movements [3]int) [][]string {
	from := movements[1] - 1
	to := movements[2] - 1
	cratesToMove := []string{}
	crateToMove := ""
	for i := 0; i < movements[0]; i++ {
		crates, crateToMove = removeCrate(crates, from)
		cratesToMove = append(cratesToMove, crateToMove)
	}
	for i, j := 0, len(cratesToMove)-1; i < j; i, j = i+1, j-1 {
		cratesToMove[i], cratesToMove[j] = cratesToMove[j], cratesToMove[i]
	}
	crates = setCrates(crates, to, cratesToMove)
	return crates
}

func removeCrate(crates [][]string, column int) ([][]string, string) {
	var crate = crates[column][len(crates[column])-1]
	crates[column] = crates[column][:len(crates[column])-1]
	return crates, crate
}

func setCrate(crates [][]string, place int, crate string) [][]string {
	crates[place] = append(crates[place], crate)
	return crates
}

func setCrates(crates [][]string, place int, cratesToSet []string) [][]string {
	crates[place] = append(crates[place], cratesToSet...)
	return crates
}

func transposeAndReverse(slice [][]string) [][]string {
	xl := len(slice[0])
	yl := len(slice)
	result := make([][]string, xl)
	for i := range result {
		result[i] = make([]string, yl)
	}
	for i := 0; i < xl; i++ {
		for j := 0; j < yl; j++ {
			result[i][len(result[i])-1-j] = slice[j][i]
		}
	}
	return result
}

func getMovements(line string) [3]int {
	movements := [3]int{0, 0, 0}
	line = strings.Replace(line, "move ", "", -1)
	line = strings.Replace(line, " from ", ",", -1)
	line = strings.Replace(line, " to ", ",", -1)
	movementsString := strings.Split(line, ",")
	movements[0], _ = strconv.Atoi(movementsString[0])
	movements[1], _ = strconv.Atoi(movementsString[1])
	movements[2], _ = strconv.Atoi(movementsString[2])
	return movements
}

func getCrates(line string) []string {
	var crate []string
	for i, letter := range line {
		if i == 1 {
			crate = append(crate, string(letter))
		} else if (i-1)%(4) == 0 {
			crate = append(crate, string(letter))
		}
	}
	return crate
}

func part2() {
	/// Open the file for reading
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// Create a new Scanner for the file
	scanner := bufio.NewScanner(file)

	crateIteration := 0
	// Use the Scanner to read each line from the file
	var crates [][]string
	for scanner.Scan() {

		line := scanner.Text()

		// parse crates matrix
		if strings.Contains(line, "[") {
			crates = append(crates, getCrates(line))
			crateIteration++
		}

		// format matrix, each stack is a row
		if strings.Contains(line, "1   2   3") {
			crates = transposeAndReverse(crates)
			crates = removeEmpty(crates)
		}

		// move crates
		if strings.Contains(line, "move") {
			movements := getMovements(line)
			crates = moveCratesSameOrder(crates, movements)
		}
	}

	fmt.Println("Crates on top of each stack are:", getTopCrates(crates))
}

func main() {
	part1()
	part2()
}
