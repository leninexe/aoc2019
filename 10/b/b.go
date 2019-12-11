package main

import (
	"fmt"
	"os"
	"log"
	"bufio"
	"strings"
	"math"
	"sort"
)

const (
	TestParam = "-t"
	Inputfile = "input.txt"
	TestInputfile = "testinput.txt"

	posX = 26
	posY = 29

	tl = 0
	tr = 1
	br = 2
	bl = 3
)

type Increment struct {
	increment float64
	sector int
	distance int
}

type Point struct {
	x int
	y int
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

	otherAsteroids := make(map[Point]Increment)

	for y1, otherElements := range asteroids {
		for x1, _ := range otherElements {
			if posX != x1 || posY != y1 {
				var sector int

				if x1 < posX && y1 < posY {
					sector = tl
				} else if x1 >= posX && y1 < posY {
					sector = tr
				} else if x1 >= posX && y1 >= posY {
					sector = br
				} else {
					sector = bl
				}

				var m float64

				if posX - x1 == 0 {
					if sector == tr || sector == bl {
						m = -1000
					} else {
						m = 1000
					}
				} else {
					m = (float64)(posY - y1) / (float64)(posX - x1)
				}

				distance := (int)(math.Abs((float64)(posY - y1)) + math.Abs((float64)(posX - x1)))
				inc := Increment{m, sector, distance}

				otherAsteroids[Point{x1, y1}] = inc
			}
		} 
	}

	sector := tr
	removed := 0
	var winner Point

	for removed < 200 {
		incs := getIncrementsForSector(otherAsteroids, sector)

		for i := 0; i < len(incs); i++ {
			nextInc := incs[i]

			var candidate Point
			minDistance := -1

			for pt, inc := range otherAsteroids {
				if inc.increment == nextInc && inc.sector == sector {
					if inc.distance < minDistance || minDistance == -1 {
						minDistance = inc.distance
						candidate = pt
					}
				}
			}

			delete(otherAsteroids, candidate)
			removed++

			if (removed == 200) {
				winner = candidate
			}
		}

		if sector == tr {
			sector = br
		} else if sector == br {
			sector = bl
		} else if sector == bl {
			sector = tl
		} else {
			sector = tr
		}
	}

	fmt.Println(winner.x * 100 + winner.y)
}

func getIncrementsForSector(incs map[Point]Increment, sector int) []float64 {
	result := make([]float64, 0)

	for _, i := range incs {
		if i.sector == sector {
			exists := false

			for _, inc := range result {
				if inc == i.increment {
					exists = true
				}
			}

			if !exists {
				result = append(result, i.increment)
			}
		}
	}

	sort.Float64s(result)
	
	return result	
}