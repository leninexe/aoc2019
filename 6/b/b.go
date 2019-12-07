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
	
	myOrbit := []string{"YOU"}
	santaOrbit := []string{"SAN"}

	x := orbits[myOrbit[0]]
	exists := true
	
	for exists {
		myOrbit = append([]string{x}, myOrbit...)
		x, exists = orbits[x]
	}

	x = orbits[santaOrbit[0]]
	exists = true
	
	for exists {
		santaOrbit = append([]string{x}, santaOrbit...)
		x, exists = orbits[x]
	}

	i := 0

	for i < len(myOrbit) && i < len(santaOrbit) {
		if myOrbit[i] != santaOrbit[i] {
			break
		}

		i++
	}

	distance := len(myOrbit) - 1 - i + len(santaOrbit) - 1 - i
	fmt.Println(distance)
}