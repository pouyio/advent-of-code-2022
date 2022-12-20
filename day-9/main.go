package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Positions struct {
	head [2]int
	tail [2]int
}

type Matrix struct {
	positions Positions
	visited   map[string]bool
}

func (matrix *Matrix) MoveHead(dir string) {
	positions := &matrix.positions
	switch dir {
	case "U":
		positions.head[1]++
	case "D":
		positions.head[1]--
	case "L":
		positions.head[0]--
	case "R":
		positions.head[0]++
	}
}

func (matrix *Matrix) UpdateVisited() {
	key := strconv.Itoa(matrix.positions.tail[0]) + "_" + strconv.Itoa(matrix.positions.tail[1])
	matrix.visited[key] = true

}

func (matrix *Matrix) MoveTail() {
	positions := &matrix.positions

	distanceX := math.Abs(float64(positions.head[0] - positions.tail[0]))
	distanceY := math.Abs(float64(positions.head[1] - positions.tail[1]))

	if distanceX >= 2 {
		if positions.head[0] > positions.tail[0] {
			positions.tail[0]++
		} else {
			positions.tail[0]--
		}
		positions.tail[1] = positions.head[1]
	} else if distanceY >= 2 {
		if positions.head[1] > positions.tail[1] {
			positions.tail[1]++
		} else {
			positions.tail[1]--
		}
		positions.tail[0] = positions.head[0]
	}

	matrix.UpdateVisited()
}

func (matrix *Matrix) GetTotalVisited() int {
	return len(matrix.visited)
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
	// Use the Scanner to read each line from the file
	matrix := Matrix{
		positions: Positions{
			head: [2]int{0, 0},
			tail: [2]int{0, 0},
		},
		visited: map[string]bool{"0_0": true},
	}
	for scanner.Scan() {
		line := scanner.Text()
		movements := strings.Split(line, " ")
		times, _ := strconv.Atoi(movements[1])

		for i := 0; i < times; i++ {
			matrix.MoveHead(movements[0])
			matrix.MoveTail()
		}
	}

	totalPositions := matrix.GetTotalVisited()

	fmt.Println("Total positions the tail visited:", totalPositions)
}

func main() {
	part1()
}
