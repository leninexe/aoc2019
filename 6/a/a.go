package main

import (
	"fmt"
	"os"
	"log"
	"bufio"
	"strings"
)

const (
	TestParam = "-t"
	Inputfile = "input.txt"
	TestInputfile = "testinput.txt"
)

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
	orbits := make(map[string]string)

	for scanner.Scan() {
		inputSlice := strings.Split(scanner.Text(), ")")

		orbits[inputSlice[1]] = inputSlice[0]
	}

	orbitSum := 0

	for k := range orbits {
		orbitCount := 0
		x := orbits[k]
		_, exists := orbits[x]

		if exists {
			for exists {
				orbitCount++
				x, exists = orbits[x]
			}

			orbitSum += orbitCount
		} else {
			orbitSum += 1
		}
	}

	fmt.Println(orbitSum)
}