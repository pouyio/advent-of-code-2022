package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func repeatedCharacter(s1, s2 string) string {
	for _, outterCharacter := range strings.Split(s1, "") {
		for _, innerCharacter := range strings.Split(s2, "") {
			if outterCharacter == innerCharacter {
				return innerCharacter
			}
		}
	}
	return ""
}

func repeatedCharacter2(s1, s2, s3 string) string {
	repeatedCharacter := ""
	for _, s1Character := range strings.Split(s1, "") {
		for _, s2Character := range strings.Split(s2, "") {
			for _, s3Character := range strings.Split(s3, "") {
				if s1Character == s2Character && s2Character == s3Character {
					repeatedCharacter = s1Character
				}
			}
		}
	}

	return repeatedCharacter
}

func characterPriority(character string) int {
	if "A" <= character && character <= "Z" {
		return int([]rune(character)[0] - 38)
	}
	return int([]rune(character)[0] - 96)
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
		first := line[:len(line)/2]
		second := line[len(line)/2:]
		character := repeatedCharacter(first, second)
		totalSum += characterPriority(character)

	}

	fmt.Println("Sum of priorities is:", totalSum)
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
	iteration := 0
	var lines [3]string
	for scanner.Scan() {
		lines[iteration] = scanner.Text()
		iteration++

		if iteration == 3 {
			iteration = 0
			character := repeatedCharacter2(lines[0], lines[1], lines[2])
			totalSum += characterPriority(character)
		}

	}

	fmt.Println("Sum of points is:", totalSum)
}

func main() {
	part1()
	part2()
}
