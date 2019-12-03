package main

import (
	"fmt"
	"os"
	"log"
	"bufio"
	"strings"
	"strconv"
)

const (
	TestParam = "-t"
	Inputfile = "input.txt"
	TestInputfile = "testinput.txt"
)

type point struct {
	x, y int
}

func main () {
	filename := Inputfile

	if len(os.Args) > 1 && os.Args[1] == TestParam {
		filename = TestInputfile
	}

	file, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	wire1 := strings.Split(scanner.Text(), ",")

	scanner.Scan()
	wire2 := strings.Split(scanner.Text(), ",")

	wireMap := make(map[point]int)
	wireMap[point{0, 0}] = 0

	currentPos := point{0, 0}

	// Wire 1
	for i := 0; i < len(wire1); i++ {
		distance, _ := strconv.Atoi(wire1[i][1:len(wire1[i])])
		dx := 0
		dy := 0

		if strings.HasPrefix(wire1[i], "U") {
			dy = -1
		} else if strings.HasPrefix(wire1[i], "L") {
			dx = -1
		} else if strings.HasPrefix(wire1[i], "R") {
			dx = 1
		} else if strings.HasPrefix(wire1[i], "D") {
			dy = 1
		}

		for j := 1; j <= distance; j++ {
			wireMap[point{currentPos.x + dx * j, currentPos.y + dy * j}] = 1
		}

		currentPos = point{currentPos.x + distance * dx, currentPos.y + distance * dy}
	}

	collisions := make([]point, 0)
	currentPos = point{0, 0}

	// Wire 2
	for i := 0; i < len(wire2); i++ {
		distance, _ := strconv.Atoi(wire2[i][1:len(wire2[i])])
		dx := 0
		dy := 0

		if strings.HasPrefix(wire2[i], "U") {
			dy = -1
		} else if strings.HasPrefix(wire2[i], "L") {
			dx = -1
		} else if strings.HasPrefix(wire2[i], "R") {
			dx = 1
		} else if strings.HasPrefix(wire2[i], "D") {
			dy = 1
		}

		for j := 1; j <= distance; j++ {
			if wireMap[point{currentPos.x + dx * j, currentPos.y + dy * j}] == 1 {
				wireMap[point{currentPos.x + dx * j, currentPos.y + dy * j}] = 3

				collisions = append(collisions, point{currentPos.x + dx * j, currentPos.y + dy * j})
			} else {
				wireMap[point{currentPos.x + dx * j, currentPos.y + dy * j}] = 2
			}
		}

		currentPos = point{currentPos.x + distance * dx, currentPos.y + distance * dy}
	}

	closest := point{0, 0}
	minSteps := 0

	for _, col := range collisions {
		sum := calculateSteps(wire1, col)
		sum += calculateSteps(wire2, col)

		if minSteps == 0 || sum < minSteps {
			minSteps = sum
			closest = col
		}
	}

	fmt.Printf("Closest collision is at [%d, %d] with %d steps\n", closest.x, closest.y, minSteps)
}

func calculateSteps(path []string, col point) int {
	steps := 1
	currentPos := point{0, 0}

	for i := 0; i < len(path); i++ {
		distance, _ := strconv.Atoi(path[i][1:len(path[i])])
		dx := 0
		dy := 0

		if strings.HasPrefix(path[i], "U") {
			dy = -1
		} else if strings.HasPrefix(path[i], "L") {
			dx = -1
		} else if strings.HasPrefix(path[i], "R") {
			dx = 1
		} else if strings.HasPrefix(path[i], "D") {
			dy = 1
		}

		for j := 1; j <= distance; j++ {
			if currentPos.x + dx * j == col.x && currentPos.y + dy * j == col.y {
				return steps
			}
		
			steps++
		}

		currentPos = point{currentPos.x + distance * dx, currentPos.y + distance * dy}
	}

	return -1
}






