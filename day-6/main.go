package main

import (
	"bufio"
	"fmt"
	"os"
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

	// Use the Scanner to read each line from the file
	cursorPointer := 0
	length := 4
	for scanner.Scan() {

		code := strings.Split(scanner.Text(), "")

		for ok := true; ok; cursorPointer++ {
			repeated := checkRepeated(code, cursorPointer, length)
			if !repeated {
				break
			}
		}
	}

	fmt.Println("Characters processed before the first start-of-packet marker detected:", cursorPointer+length)
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

	// Use the Scanner to read each line from the file
	cursorPointer := 0
	length := 14
	for scanner.Scan() {

		code := strings.Split(scanner.Text(), "")

		for ok := true; ok; cursorPointer++ {
			repeated := checkRepeated(code, cursorPointer, length)
			if !repeated {
				break
			}
		}
	}

	fmt.Println("Characters processed before the first start-of-message marker detected:", cursorPointer+14)
}

func checkRepeated(code []string, cursorPointer int, length int) bool {
	visited := make(map[string]bool, 0)
	for i := cursorPointer; i < cursorPointer+length; i++ {
		if visited[code[i]] == true {
			return true
		} else {
			visited[code[i]] = true
		}
	}
	return false
}

func main() {
	part1()
	part2()
}
