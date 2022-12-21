package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Package struct {
	value    int
	elements []*Package
	parent   *Package
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

	indexSum := 0
	index := 1
	// Use the Scanner to read each line from the file
	for scanner.Scan() {
		line1 := scanner.Text()
		scanner.Scan()
		line2 := scanner.Text()

		package1 := parsePackage(line1)
		package2 := parsePackage(line2)

		if areOrdered(package1, package2) == 1 {
			indexSum += index
		}

		index++
		scanner.Scan()
	}

	fmt.Println(indexSum)
}

func parsePackage(input string) Package {
	root := Package{-1, []*Package{}, nil}
	temp := &root

	var currentNumber string
	for _, r := range input {
		switch string(r) {
		case "[":
			newTree := Package{value: -1, elements: []*Package{}, parent: temp}
			temp.elements = append(temp.elements, &newTree)
			temp = &newTree
		case "]":
			if len(currentNumber) > 0 {
				number, _ := strconv.Atoi(currentNumber)
				temp.value = number
				currentNumber = ""
			}
			temp = temp.parent
		case ",":
			if len(currentNumber) > 0 {
				number, _ := strconv.Atoi(currentNumber)
				temp.value = number
				currentNumber = ""
			}
			temp = temp.parent
			newTree := Package{value: -1, elements: []*Package{}, parent: temp}
			temp.elements = append(temp.elements, &newTree)
			temp = &newTree
		default:
			currentNumber += string(r)
		}
	}
	return root
}

func areOrdered(first, second Package) int {
	switch {
	case len(first.elements) == 0 && len(second.elements) == 0:
		if first.value > second.value {
			return -1
		} else if first.value == second.value {
			return 0
		}
		return 1

	case first.value >= 0:
		return areOrdered(Package{-1, []*Package{&first}, nil}, second)

	case second.value >= 0:
		return areOrdered(first, Package{-1, []*Package{&second}, nil})
	default:
		var i int
		for i = 0; i < len(first.elements) && i < len(second.elements); i++ {
			ordered := areOrdered(*first.elements[i], *second.elements[i])
			if ordered != 0 {
				return ordered
			}
		}
		if i < len(first.elements) {
			return -1
		} else if i < len(second.elements) {
			return 1
		}
	}
	return 0
}

func main() {
	part1()
}
