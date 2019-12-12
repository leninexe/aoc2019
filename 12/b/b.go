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
	originalMoons := make([]Moon, 0)

	for scanner.Scan() {
		line := scanner.Text()
		coords := strings.Split(line[1:len(line)-1], ",")
		x, _ := strconv.Atoi(strings.Split(coords[0], "=")[1])
		y, _ := strconv.Atoi(strings.Split(coords[1], "=")[1])
		z, _ := strconv.Atoi(strings.Split(coords[2], "=")[1])

		originalMoons = append(originalMoons, Moon{x, y, z, 0, 0, 0})
	}

	results := make([]int, 3)

	for position := 0; position < 3; position++ {
		moons := make([]Moon, len(originalMoons))

		for i := 0; i < len(originalMoons); i++ {
			moons[i] = Moon{originalMoons[i].x, originalMoons[i].y, originalMoons[i].z, originalMoons[i].vx, originalMoons[i].vy, originalMoons[i].vz}
		}

		duplicateFound := false
		stringHistory := make(map[string]int)
		rounds := 0

		for !duplicateFound {
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

			energy := 0
			stringValue := ""

			for i := 0; i < len(moons); i++ {
				cm := moons[i]
				energy += (int)((math.Abs((float64)(cm.x)) + math.Abs((float64)(cm.y)) + math.Abs((float64)(cm.z))) * (math.Abs((float64)(cm.vx)) + math.Abs((float64)(cm.vy)) + math.Abs((float64)(cm.vz))))
				
				if position == 0 {
					stringValue += strconv.Itoa(cm.x) + strconv.Itoa(cm.vx)
				} else if position == 1 {
					stringValue += strconv.Itoa(cm.y) + strconv.Itoa(cm.vy)
				} else {
					stringValue += strconv.Itoa(cm.z) + strconv.Itoa(cm.vz)
				}

			}

			previousRound, exists := stringHistory[stringValue]

			if exists {
				if position == 0 {
					fmt.Printf("Found a duplicate for X in round %d and %d\n", previousRound, rounds)
				} else if position == 1 {
					fmt.Printf("Found a duplicate for Y in round %d and %d\n", previousRound, rounds)
				} else {
					fmt.Printf("Found a duplicate for Z in round %d and %d\n", previousRound, rounds)
				}

				results[position] = rounds - previousRound
				duplicateFound = true
			}

			stringHistory[stringValue] = rounds

			rounds++
		}
	}

	fmt.Println(results)

	lcm1 := lcm(results[0], results[1])
	lcm2 := lcm(lcm1, results[2])

	fmt.Println(lcm2)
}

func lcm(a int, b int) int {
	return (a * b) / gcd(a, b)
}

func gcd(a int, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}

	return a
}