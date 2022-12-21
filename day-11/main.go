package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Operation func(old int) int

type Test struct {
	divisibleBy int
	ifTrue      int
	ifFalse     int
}

type Monkey struct {
	items        []int
	Operation    Operation
	test         Test
	inspectCount int
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
	monkeys := []Monkey{}
	// Use the Scanner to read each line from the file
	for scanner.Scan() {
		line := scanner.Text()

		// create new monkey
		if strings.HasPrefix(line, "Monkey ") {
			monkeys = append(monkeys, Monkey{inspectCount: 0})
		}

		// add items to last monkey
		if strings.HasPrefix(line, "  Starting items: ") {
			lineArr := strings.Split(line, "  Starting items: ")
			items := strings.Split(lineArr[1], ", ")
			for _, item := range items {
				itemInt, _ := strconv.Atoi(item)
				monkeys[len(monkeys)-1].items = append(monkeys[len(monkeys)-1].items, itemInt)
			}
		}

		// add operation to last monkey
		if strings.HasPrefix(line, "  Operation: new = old ") {
			lineArr := strings.Split(line, "  Operation: new = old ")
			operandsArr := strings.Split(lineArr[1], " ")
			switch operandsArr[1] {
			case "old":
				if operandsArr[0] == "+" {
					monkeys[len(monkeys)-1].Operation = func(old int) int {
						return old + old
					}
				} else {
					monkeys[len(monkeys)-1].Operation = func(old int) int {
						return old * old
					}
				}
			default:
				value, _ := strconv.Atoi(operandsArr[1])
				if operandsArr[0] == "+" {
					monkeys[len(monkeys)-1].Operation = func(old int) int {
						return old + value
					}
				} else {
					monkeys[len(monkeys)-1].Operation = func(old int) int {
						return old * value
					}
				}
			}

		}

		// add test value to last monkey
		if strings.HasPrefix(line, "  Test: divisible by ") {
			lineArr := strings.Split(line, "  Test: divisible by ")
			value, _ := strconv.Atoi(lineArr[1])
			monkeys[len(monkeys)-1].test.divisibleBy = value
		}

		// add next monkey if test is true
		if strings.HasPrefix(line, "    If true: throw to monkey ") {
			lineArr := strings.Split(line, "    If true: throw to monkey ")
			value, _ := strconv.Atoi(lineArr[1])
			monkeys[len(monkeys)-1].test.ifTrue = value
		}
		// add next monkey if test is false
		if strings.HasPrefix(line, "    If false: throw to monkey ") {
			lineArr := strings.Split(line, "    If false: throw to monkey ")
			value, _ := strconv.Atoi(lineArr[1])
			monkeys[len(monkeys)-1].test.ifFalse = value
		}

	}

	// execute rounds and update monkeys
	for round := 0; round < 20; round++ {
		for i, monkey := range monkeys {
			for _, item := range monkey.items {
				monkeys[i].inspectCount++
				newItem := math.Floor(float64(monkey.Operation(item)) / 3)
				to := monkey.test.ifFalse
				if int(newItem)%monkey.test.divisibleBy == 0 {
					to = monkey.test.ifTrue
				}
				monkeys[to].items = append(monkeys[to].items, int(newItem))
			}
			monkeys[i].items = []int{}
		}
	}

	// sort monkeys by inspected items count
	inspectCounts := []int{}
	for _, monkey := range monkeys {
		inspectCounts = append(inspectCounts, monkey.inspectCount)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(inspectCounts)))
	fmt.Println(inspectCounts[0] * inspectCounts[1])
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
	monkeys := []Monkey{}
	// Use the Scanner to read each line from the file
	for scanner.Scan() {
		line := scanner.Text()

		// create new monkey
		if strings.HasPrefix(line, "Monkey ") {
			monkeys = append(monkeys, Monkey{inspectCount: 0})
		}

		// add items to last monkey
		if strings.HasPrefix(line, "  Starting items: ") {
			lineArr := strings.Split(line, "  Starting items: ")
			items := strings.Split(lineArr[1], ", ")
			for _, item := range items {
				itemInt, _ := strconv.Atoi(item)
				monkeys[len(monkeys)-1].items = append(monkeys[len(monkeys)-1].items, itemInt)
			}
		}

		// add operation to last monkey
		if strings.HasPrefix(line, "  Operation: new = old ") {
			lineArr := strings.Split(line, "  Operation: new = old ")
			operandsArr := strings.Split(lineArr[1], " ")
			switch operandsArr[1] {
			case "old":
				if operandsArr[0] == "+" {
					monkeys[len(monkeys)-1].Operation = func(old int) int {
						return old + old
					}
				} else {
					monkeys[len(monkeys)-1].Operation = func(old int) int {
						return old * old
					}
				}
			default:
				value, _ := strconv.Atoi(operandsArr[1])
				if operandsArr[0] == "+" {
					monkeys[len(monkeys)-1].Operation = func(old int) int {
						return old + value
					}
				} else {
					monkeys[len(monkeys)-1].Operation = func(old int) int {
						return old * value
					}
				}
			}

		}

		// add test value to last monkey
		if strings.HasPrefix(line, "  Test: divisible by ") {
			lineArr := strings.Split(line, "  Test: divisible by ")
			value, _ := strconv.Atoi(lineArr[1])
			monkeys[len(monkeys)-1].test.divisibleBy = value
		}

		// add next monkey if test is true
		if strings.HasPrefix(line, "    If true: throw to monkey ") {
			lineArr := strings.Split(line, "    If true: throw to monkey ")
			value, _ := strconv.Atoi(lineArr[1])
			monkeys[len(monkeys)-1].test.ifTrue = value
		}
		// add next monkey if test is false
		if strings.HasPrefix(line, "    If false: throw to monkey ") {
			lineArr := strings.Split(line, "    If false: throw to monkey ")
			value, _ := strconv.Atoi(lineArr[1])
			monkeys[len(monkeys)-1].test.ifFalse = value
		}

	}

	commonMultiple := 1
	for _, monkey := range monkeys {
		commonMultiple *= monkey.test.divisibleBy
	}

	// execute rounds and update monkeys
	for round := 0; round < 10000; round++ {
		for i, monkey := range monkeys {
			for _, item := range monkey.items {
				monkeys[i].inspectCount++
				newItem := monkey.Operation(item) % commonMultiple
				to := monkey.test.ifFalse
				if int(newItem)%monkey.test.divisibleBy == 0 {
					to = monkey.test.ifTrue
				}
				monkeys[to].items = append(monkeys[to].items, int(newItem))
			}
			monkeys[i].items = []int{}
		}
	}

	// sort monkeys by inspected items count
	inspectCounts := []int{}
	for _, monkey := range monkeys {
		inspectCounts = append(inspectCounts, monkey.inspectCount)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(inspectCounts)))
	fmt.Println(inspectCounts[0] * inspectCounts[1])
}

func main() {
	part1()
	part2()
}
