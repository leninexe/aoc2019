package main

import (
	"fmt"
	"os"
	"log"
	"bufio"
	"strings"
	"strconv"
	"math"
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
	minX := 0
	maxX := 0
	minY := 0
	maxY := 0

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

		if currentPos.x < minX {
			minX = currentPos.x
		}

		if currentPos.x > maxX {
			maxX = currentPos.x
		}

		if currentPos.y < minY {
			minY = currentPos.y
		}

		if currentPos.y > maxY {
			maxY = currentPos.y
		}
	}

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
			} else {
				wireMap[point{currentPos.x + dx * j, currentPos.y + dy * j}] = 2
			}
		}

		currentPos = point{currentPos.x + distance * dx, currentPos.y + distance * dy}

		if currentPos.x < minX {
			minX = currentPos.x
		}

		if currentPos.x > maxX {
			maxX = currentPos.x
		}

		if currentPos.y < minY {
			minY = currentPos.y
		}

		if currentPos.y > maxY {
			maxY = currentPos.y
		}
	}

	// Detect closest shortcircuit
	r := 1
	bestDistance := 0

	fmt.Println(int(math.Max(math.Max(math.Abs(float64(minX)), math.Abs(float64(minY))), math.Max(float64(maxX), float64(maxY)))))

	for r < int(math.Max(math.Max(math.Abs(float64(minX)), math.Abs(float64(minY))), math.Max(float64(maxX), float64(maxY)))) && bestDistance == 0 {
		x := r * -1
		y := r * -1

		// top
		for i := x; i <= r; i++ {
			if wireMap[point{i, y}] == 3 {
				distance := int(math.Abs(float64(i)) + math.Abs(float64(y)))
				fmt.Printf("Shortcircuit at [%d, %d], distance = %d\n", i, y, distance)

				if bestDistance == 0 || distance < bestDistance {
					bestDistance = distance
				}
			}
		}

		// bottom
		for i := x; i <= r; i++ {
			if wireMap[point{i, r}] == 3 {
				distance := int(math.Abs(float64(i)) + math.Abs(float64(r)))
				fmt.Printf("Shortcircuit at [%d, %d], distance = %d\n", i, r, distance)

				if bestDistance == 0 || distance < bestDistance {
					bestDistance = distance
				}
			}
		}

		// left
		for i := y + 1; i < r; i++ {
			if wireMap[point{x, i}] == 3 {
				distance := int(math.Abs(float64(x)) + math.Abs(float64(i)))
				fmt.Printf("Shortcircuit at [%d, %d], distance = %d\n", x, i, distance)

				if bestDistance == 0 || distance < bestDistance {
					bestDistance = distance
				}
			}
		}

		// right
		for i := y + 1; i < r; i++ {
			if wireMap[point{r, i}] == 3 {
				distance := int(math.Abs(float64(r)) + math.Abs(float64(i)))
				fmt.Printf("Shortcircuit at [%d, %d], distance = %d\n", r, i, distance)

				if bestDistance == 0 || distance < bestDistance {
					bestDistance = distance
				}
			}
		}

		r++
	}

	fmt.Printf("Best distance %d\n", bestDistance)
}