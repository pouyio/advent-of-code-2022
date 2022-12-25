package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

type Valve struct {
	at       string
	time     int
	released int
	opened   []string
}
type Node struct {
	rate    int
	tunnels []string
}

func part1() {
	//Read input file
	input, _ := os.Open("input.txt")
	sc := bufio.NewScanner(input)
	defer input.Close()

	valves := map[string]Node{}
	for sc.Scan() {
		input := strings.Fields(sc.Text())
		var newNode Node
		fmt.Sscanf(input[4], "rate=%d;", &newNode.rate)
		for _, tunnel := range input[9:] {
			newNode.tunnels = append(newNode.tunnels, tunnel[:2])
		}
		valves[input[1]] = newNode
	}

	queue := []Valve{{at: "AA", time: 30, released: 0, opened: []string{}}}
	best := 0
	visited := map[string]bool{}

	for ok := true; ok; ok = len(queue) != 0 {
		current := queue[0]
		queue = queue[1:]
		identifier := current.at + "," + strings.Join(current.opened, ",")

		if current.time > 0 && !visited[identifier] {
			visited[identifier] = true
			valve := valves[current.at]
			newOpened := make([]string, len(current.opened))
			copy(newOpened, current.opened)
			for _, tunnel := range valve.tunnels {
				newTunnel := Valve{at: tunnel, time: current.time - 1, released: current.released, opened: newOpened}
				queue = append(queue, newTunnel)
			}

			isIncluded := false
			for _, openend := range newOpened {
				if openend == current.at {
					isIncluded = true
				}
			}

			if valve.rate > 0 && current.time > 1 && !isIncluded {
				released := current.released + valve.rate*(current.time-1)
				for _, tunnel := range valve.tunnels {
					queue = append(queue, Valve{at: tunnel, time: current.time - 2, released: released, opened: append(newOpened, current.at)})
				}
			}
		} else {
			best = int(math.Max(float64(best), float64(current.released)))
		}
	}
	fmt.Println(best)

}

func main() {
	part1()
}
