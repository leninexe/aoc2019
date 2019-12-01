package main

import (
	"fmt"
	"os"
	"log"
	"bufio"
	"strconv"
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
	sum := 0

	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())

		if err != nil {
			log.Fatal(err)
		}

		sum += calculateFuel(num)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(sum)
}

func calculateFuel(mass int) int {
	fuel := mass / 3 - 2

	if fuel > 0 {
		return fuel + calculateFuel(fuel)
	} else {
		return 0
	}
}