package main

import (
	"fmt"
	"os"
	"log"
	"bufio"
	"strings"
	// "strconv"
)

const (
	TestParam = "-t"
	Inputfile = "input.txt"
	TestInputfile = "testinput.txt"
)

type Increment struct {
	increment float32
	sign int
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

	asteroids := make(map[int]map[int]string)
	scanner := bufio.NewScanner(file)

	y := 0

	for scanner.Scan() {
		asteroids[y] = make(map[int]string)

		line := strings.Split(scanner.Text(),"")

		for x := 0; x < len(line); x++ {
			if line[x] == "#" {
				asteroids[y][x] = line[x]
			}
		}

		y++
	}

	maxCount := 0

	for y, elements := range asteroids {
		for x, _ := range elements {
			incs := make(map[Increment]bool)
			count := 0

			for y1, otherElements := range asteroids {
				for x1, _ := range otherElements {
					if x != x1 || y != y1 {
						var sign int

						if x1 >= x && y1 >= y {
							sign = 4
						} else if x1 < x && y1 < y {
							sign = 1
						} else if x1 < x && y1 >= y {
							sign = 3
						} else {
							sign = 4
						}

						inc := Increment{(float32)(y - y1) / (float32)(x - x1), sign}

						if !incs[inc] {
							count++
						}

						incs[inc] = true
					}
				} 
			}

			if count > maxCount {
				maxCount = count
			}
		}
	}

	fmt.Println(maxCount)
}