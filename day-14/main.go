package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getKey(x, y int) string {
	return strconv.Itoa(x) + "," + strconv.Itoa(y)
}

func dropSand(grid map[string]string, x, y, maxY int) (position string, full bool) {
	// is y>=maxY
	if y >= maxY {
		// yes: return "", true
		return "", true
	}

	// no: is down available?
	if grid[getKey(x, y+1)] == "" {
		// yes: +1y ->continue next
		return dropSand(grid, x, y+1, maxY)
	}

	// no: is left down available?
	if grid[getKey(x-1, y+1)] == "" {
		// yes: add -1x +1y ->continue next
		return dropSand(grid, x-1, y+1, maxY)
	}

	// no: is left right available?
	if grid[getKey(x+1, y+1)] == "" {
		// yes: add +1x +1y ->continue next
		return dropSand(grid, x+1, y+1, maxY)
	}
	// no: is settled -> return position, false
	return getKey(x, y), false

}

func dropSand2(grid map[string]string, x, y, maxY int) (position string, full bool) {
	// is y>=maxY
	if y == maxY {
		// yes: return position false
		return getKey(x, y), false
	}

	// is down available?
	if grid[getKey(x, y+1)] == "" {
		// yes: +1y ->continue next
		return dropSand2(grid, x, y+1, maxY)
	}

	// no: is left down available?
	if grid[getKey(x-1, y+1)] == "" {
		// yes: add -1x +1y ->continue next
		return dropSand2(grid, x-1, y+1, maxY)
	}

	// no: is left right available?
	if grid[getKey(x+1, y+1)] == "" {
		// yes: add +1x +1y ->continue next
		return dropSand2(grid, x+1, y+1, maxY)
	}
	// no: is settled on top?
	if x == 500 && y == 0 {
		// yes: -> return position, true
		return getKey(x, y), true
	}
	// no: is full return position, false
	return getKey(x, y), false

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

	grid := map[string]string{}
	maxY := 0
	// Use the Scanner to read each line from the file
	for scanner.Scan() {
		line := scanner.Text()
		rockCorners := strings.Split(line, " -> ")
		for i := 0; i < len(rockCorners); i++ {
			// save rock point
			grid[rockCorners[i]] = "rock"
			pointsStr := strings.Split(rockCorners[i], ",")
			pointX, _ := strconv.Atoi(pointsStr[0])
			pointY, _ := strconv.Atoi(pointsStr[1])

			// save rock line
			if i+1 < len(rockCorners) {
				nextPointStr := strings.Split(rockCorners[i+1], ",")
				nextPointX, _ := strconv.Atoi(nextPointStr[0])
				nextPointY, _ := strconv.Atoi(nextPointStr[1])
				if pointX == nextPointX {
					// save points horizontlly
					if pointY < nextPointY {
						for piy := pointY + 1; piy < nextPointY; piy++ {
							grid[getKey(pointX, piy)] = "rock"
						}
					} else {
						for pix := nextPointY + 1; pix < pointY; pix++ {
							grid[getKey(pointX, pix)] = "rock"
						}

					}
				} else {
					// save points vertically
					if pointX < nextPointX {
						for pix := pointX + 1; pix < nextPointX; pix++ {
							grid[getKey(pix, pointY)] = "rock"
						}
					} else {
						for pix := nextPointX + 1; pix < pointX; pix++ {
							grid[getKey(pix, pointY)] = "rock"
						}

					}

				}

			}

			if pointY > maxY {
				maxY = pointY
			}
		}

	}

	var position string
	sandCounter := 0
	for full := false; !full; {
		sandCounter++
		position, full = dropSand(grid, 500, 0, maxY)
		grid[position] = "sand"
	}

	fmt.Println("Total sand units:", sandCounter-1)

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

	grid := map[string]string{}
	maxY := 0
	// Use the Scanner to read each line from the file
	for scanner.Scan() {
		line := scanner.Text()
		rockCorners := strings.Split(line, " -> ")
		for i := 0; i < len(rockCorners); i++ {
			// save rock point
			grid[rockCorners[i]] = "rock"
			pointsStr := strings.Split(rockCorners[i], ",")
			pointX, _ := strconv.Atoi(pointsStr[0])
			pointY, _ := strconv.Atoi(pointsStr[1])

			// save rock line
			if i+1 < len(rockCorners) {
				nextPointStr := strings.Split(rockCorners[i+1], ",")
				nextPointX, _ := strconv.Atoi(nextPointStr[0])
				nextPointY, _ := strconv.Atoi(nextPointStr[1])
				if pointX == nextPointX {
					// save points horizontlly
					if pointY < nextPointY {
						for piy := pointY + 1; piy < nextPointY; piy++ {
							grid[getKey(pointX, piy)] = "rock"
						}
					} else {
						for pix := nextPointY + 1; pix < pointY; pix++ {
							grid[getKey(pointX, pix)] = "rock"
						}

					}
				} else {
					// save points vertically
					if pointX < nextPointX {
						for pix := pointX + 1; pix < nextPointX; pix++ {
							grid[getKey(pix, pointY)] = "rock"
						}
					} else {
						for pix := nextPointX + 1; pix < pointX; pix++ {
							grid[getKey(pix, pointY)] = "rock"
						}

					}

				}

			}

			if pointY > maxY {
				maxY = pointY
			}
		}

	}

	var position string
	sandCounter := 0
	for full := false; !full; {
		sandCounter++
		position, full = dropSand2(grid, 500, 0, maxY+1)
		grid[position] = "sand"
	}

	fmt.Println("Total sand units 2:", sandCounter)

}

func main() {
	part1()
	part2()
}
