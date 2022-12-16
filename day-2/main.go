package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// rock, paper, scissors
// a,x   b,y    c,z

const (
	ROCK     = "A"
	PAPER    = "B"
	SCISSORS = "C"
)

const (
	ROCK_ME     = "X"
	PAPER_ME    = "Y"
	SCISSORS_ME = "Z"
)

var toWin = map[string]string{
	ROCK:     PAPER,
	PAPER:    SCISSORS,
	SCISSORS: ROCK,
}

var toLose = map[string]string{
	ROCK:     SCISSORS,
	PAPER:    ROCK,
	SCISSORS: PAPER,
}

func isWin(oponent, me string) bool {
	return (me == ROCK_ME && oponent == SCISSORS) || (me == PAPER_ME && oponent == ROCK) || (me == SCISSORS_ME && oponent == PAPER)
}

func isDraw(oponent, me string) bool {
	return (me == ROCK_ME && oponent == ROCK) || (me == PAPER_ME && oponent == PAPER) || (me == SCISSORS_ME && oponent == SCISSORS)
}

func gamePoints(oponent, me string) int {
	if isWin(oponent, me) {
		return 6
	}

	if isDraw(oponent, me) {
		return 3
	}

	return 0
}

func handPoints(hand string) int {
	switch hand {
	case ROCK_ME:
	case ROCK:
		return 1
	case PAPER_ME:
	case PAPER:
		return 2
	case SCISSORS_ME:
	case SCISSORS:
		return 3
	}
	return 0
}

func gamePoints2(result string) int {
	switch result {
	case "Y":
		return 3
	case "Z":
		return 6
	}
	return 0
}

func calculateMyHand(oponent, result string) string {
	if result == "Y" {
		return oponent
	}
	if result == "Z" {
		return toWin[oponent]
	}
	return toLose[oponent]

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

		play := scanner.Text()
		hands := strings.Split(play, " ")
		totalSum += gamePoints(hands[0], hands[1]) + handPoints(hands[1])

	}

	fmt.Println("Sum of points is:", totalSum)
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

		play := scanner.Text()
		inputs := strings.Split(play, " ")
		totalSum += gamePoints2(inputs[1]) + handPoints(calculateMyHand(inputs[0], inputs[1]))

	}

	fmt.Println("Sum of points is:", totalSum)
}

func main() {
	part1() // 13675
	part2()
}
