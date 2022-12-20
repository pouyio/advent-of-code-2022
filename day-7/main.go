package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	Children *[]Node
	Parent   *Node
	Files    *[]int
	Name     string
	Size     int
}

var sum = 0

func (node *Node) calculateSize() int {
	size := 0

	for _, value := range *node.Files {
		size += value
	}

	for i := 0; i < len(*node.Children); i++ {
		size += (*node.Children)[i].calculateSize()
	}
	// not working, reference lost and size not updated
	// for _, value := range *node.Children {
	// 	size += value.calculateSize()
	// }

	node.Size += size
	return size
}

func (node *Node) AddDir(name string) {
	newDir := Node{
		Children: &[]Node{},
		Files:    &[]int{},
		Size:     0,
		Name:     name,
		Parent:   node,
	}
	*node.Children = append(*node.Children, newDir)
}

func (node *Node) AddFile(size int) {
	*node.Files = append(*node.Files, size)
}

func (node *Node) MoveTo(dir string) *Node {
	if dir == ".." {
		return node.Parent
	}
	for _, value := range *node.Children {
		if value.Name == dir {
			return &value
		}
	}
	return node.Parent
}

func part1() {
	const MAX = 100000
	// Open the file for reading
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// Create a new Scanner for the file
	scanner := bufio.NewScanner(file)
	var tree = Node{
		Children: &[]Node{},
		Files:    &[]int{},
		Name:     "/",
	}
	var cursor *Node
	// Use the Scanner to read each line from the file
	for scanner.Scan() {
		line := scanner.Text()
		if line == "$ cd /" {
			cursor = &tree
		} else if strings.HasPrefix(line, "$ ") {
			command := line[2:]
			if strings.HasPrefix(command, "cd ") {
				// command move to a dir
				cursor = cursor.MoveTo(command[3:])
			} else {
				// command list files/dir
			}
		} else if strings.HasPrefix(line, "dir ") {
			// is a directory
			cursor.AddDir(line[4:])
		} else {
			// is a file with size
			fileData := strings.Split(line, " ")
			size, _ := strconv.Atoi(fileData[0])
			cursor.AddFile(size)
		}
	}

	tree.calculateSize()

	queue := []Node{tree}
	total := 0

	for {
		current := queue[len(queue)-1]
		queue = append(queue[:len(queue)-1])

		if current.Size <= MAX {
			total += current.Size
		}

		queue = append(queue, *current.Children...)
		if len(queue) == 0 {
			break
		}
	}

	fmt.Println("Sum of the total sizes of directories:", total)
}

func part2() {
	const MAX = 100000
	// Open the file for reading
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// Create a new Scanner for the file
	scanner := bufio.NewScanner(file)
	var tree = Node{
		Children: &[]Node{},
		Files:    &[]int{},
		Name:     "/",
	}
	var cursor *Node
	// Use the Scanner to read each line from the file
	for scanner.Scan() {
		line := scanner.Text()
		if line == "$ cd /" {
			cursor = &tree
		} else if strings.HasPrefix(line, "$ ") {
			command := line[2:]
			if strings.HasPrefix(command, "cd ") {
				// command move to a dir
				cursor = cursor.MoveTo(command[3:])
			} else {
				// command list files/dir
			}
		} else if strings.HasPrefix(line, "dir ") {
			// is a directory
			cursor.AddDir(line[4:])
		} else {
			// is a file with size
			fileData := strings.Split(line, " ")
			size, _ := strconv.Atoi(fileData[0])
			cursor.AddFile(size)
		}
	}

	tree.calculateSize()

	diskAvailable := 70000000 - tree.Size
	spaceNeeded := 30000000 - diskAvailable

	smallest := int(math.Inf(1))

	queue := []Node{tree}

	for {
		current := queue[len(queue)-1]
		queue = append(queue[:len(queue)-1])

		if current.Size >= spaceNeeded {
			smallest = int(math.Min(float64(smallest), float64(current.Size)))
		}

		queue = append(queue, *current.Children...)
		if len(queue) == 0 {
			break
		}
	}

	fmt.Println("Sum of the total size of the directory to delete:", smallest)
}

func main() {
	// part1()
	part2()
}
