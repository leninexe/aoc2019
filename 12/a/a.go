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

type Moon struct {
	x int
	y int
	z int
	vx int
	vy int
	vz int
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
	moons := make([]Moon, 0)

	for scanner.Scan() {
		line := scanner.Text()
		coords := strings.Split(line[1:len(line)-1], ",")
		x, _ := strconv.Atoi(strings.Split(coords[0], "=")[1])
		y, _ := strconv.Atoi(strings.Split(coords[1], "=")[1])
		z, _ := strconv.Atoi(strings.Split(coords[2], "=")[1])

		moons = append(moons, Moon{x, y, z, 0, 0, 0})
	}

	for rounds := 0; rounds < 1000; rounds++ {
		for i := 0; i < len(moons); i++ {
			cm := moons[i]

			for j := 0; j < len(moons); j++ {
				if i != j {
					om := moons[j]

					if cm.x > om.x {
						cm.vx--
					} else if cm.x < om.x {
						cm.vx++
					}

					if cm.y > om.y {
						cm.vy--
					} else if cm.y < om.y {
						cm.vy++
					}

					if cm.z > om.z {
						cm.vz--
					} else if cm.z < om.z {
						cm.vz++
					}
				}
			}

			moons[i] = cm
		}

		for i := 0; i < len(moons); i++ {
			cm := moons[i]		

			cm.x += cm.vx
			cm.y += cm.vy
			cm.z += cm.vz

			moons[i] = cm
		}
	}

	energy := 0

	for i := 0; i < len(moons); i++ {
		cm := moons[i]
		energy += (int)((math.Abs((float64)(cm.x)) + math.Abs((float64)(cm.y)) + math.Abs((float64)(cm.z))) * (math.Abs((float64)(cm.vx)) + math.Abs((float64)(cm.vy)) + math.Abs((float64)(cm.vz))))
	}

	fmt.Println(energy)
}