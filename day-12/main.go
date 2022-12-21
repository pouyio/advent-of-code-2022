package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x int
	y int
}

type Move struct {
	x      int
	y      int
	steps  int
	height int
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

	var start Point
	var end Point
	grid := [][]int{}
	row := 0

	// Use the Scanner to read each line from the file
	for scanner.Scan() {
		line := scanner.Text()
		lineInt := []int{}
		letters := strings.Split(line, "")
		for i, letter := range letters {
			lineInt = append(lineInt, int(letter[0]))
			if letter == "S" {
				start = Point{x: i, y: row}
			} else if letter == "E" {
				end = Point{x: i, y: row}
			}
		}
		grid = append(grid, lineInt)
		row++
	}

	queue := []Move{{x: start.x, y: start.y, steps: 0, height: 9999}}
	deltas := []Point{{x: 1, y: 0}, {x: -1, y: 0}, {x: 0, y: 1}, {x: 0, y: -1}}
	visited := map[string]bool{}

	for ok := true; ok; ok = len(queue) != 0 {

		current := queue[0]
		queue = queue[1:]
		marker := strconv.Itoa(current.x) + "_" + strconv.Itoa(current.y)
		if !visited[marker] {
			visited[marker] = true
			for _, delta := range deltas {
				target := Point{x: current.x + delta.x, y: current.y + delta.y}
				if target.x >= 0 && target.y >= 0 && target.x < len(grid[0]) && target.y < len(grid) {
					if target.x == end.x && target.y == end.y && current.height == int("z"[0]) {
						fmt.Println("Fewest steps required 1: ", current.steps+1)
						break
					}
					targetHeight := grid[target.y][target.x]

					if targetHeight <= current.height+1 {
						queue = append(queue, Move{x: target.x, y: target.y, height: targetHeight, steps: current.steps + 1})
					}
				}
			}
		}
	}

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

	var end Point
	grid := [][]int{}
	row := 0
	startHeight := int("a"[0])

	// Use the Scanner to read each line from the file
	for scanner.Scan() {
		line := scanner.Text()
		lineInt := []int{}
		letters := strings.Split(line, "")
		for i, letter := range letters {
			lineInt = append(lineInt, int(letter[0]))
			// grid = append(grid, lineInt)
			if letter == "S" {
				lineInt[i] = startHeight
			} else if letter == "E" {
				end = Point{x: i, y: row}
			}
		}
		grid = append(grid, lineInt)
		row++
	}

	queue := []Move{{x: end.x, y: end.y, steps: 0, height: 122}}
	deltas := []Point{{x: 1, y: 0}, {x: -1, y: 0}, {x: 0, y: 1}, {x: 0, y: -1}}
	visited := map[string]bool{}

	for ok := true; ok; ok = len(queue) != 0 {

		current := queue[0]
		queue = queue[1:]
		marker := strconv.Itoa(current.x) + "_" + strconv.Itoa(current.y)
		if !visited[marker] {
			visited[marker] = true
			for _, delta := range deltas {
				target := Point{x: current.x + delta.x, y: current.y + delta.y}
				if target.x >= 0 && target.y >= 0 && target.x < len(grid[0]) && target.y < len(grid) {
					targetHeight := grid[target.y][target.x]
					if targetHeight == startHeight && current.height == startHeight+1 {
						fmt.Println("Fewest steps required 2: ", current.steps+1)
						break
					}

					if targetHeight >= current.height-1 {
						queue = append(queue, Move{x: target.x, y: target.y, height: targetHeight, steps: current.steps + 1})
					}
				}
			}
		}
	}

}

func main() {
	part1()
	part2()
}
