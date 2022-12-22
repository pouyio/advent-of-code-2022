package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
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

	line := map[int]bool{}
	const LINE_Y = 2000000

	// Use the Scanner to read each line from the file
	for scanner.Scan() {
		var sensorX, sensorY, beaconX, beaconY int
		fmt.Sscanf(scanner.Text(), "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &sensorX, &sensorY, &beaconX, &beaconY)
		distanceFromBeacon := manhattan(sensorX, sensorY, beaconX, beaconY)

		distanceFromLine := int(math.Abs(float64(sensorY - LINE_Y)))

		for i := 0; i <= distanceFromBeacon-distanceFromLine; i++ {
			line[sensorX+i] = true
			line[sensorX-i] = true
		}

		if beaconY == LINE_Y {
			delete(line, beaconX)
		}
	}

	fmt.Println("Positions without beacon:", len(line))
}

func manhattan(sensorX, sensorY, beaconX, beaconY int) int {
	x := math.Abs(float64(sensorX) - float64(beaconX))
	y := math.Abs(float64(sensorY) - float64(beaconY))
	return int(x + y)
}

func main() {
	part1()
}
