package main

import (
	"fmt"
	"os"
	"log"
	"bufio"
)

const (
	TestParam = "-t"
	Inputfile = "input.txt"
	TestInputfile = "testinput.txt"

	width = 25
	height = 6
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

	scanner.Scan()

	input := scanner.Text()
	var numbers []int
	
	for _, c := range input {
		numbers = append(numbers, int(c - '0'))
	}

	layer := -1
	var layers [][]int

	for i := 0; i < len(numbers); i++ {
		if i % (width * height) == 0 {
			layer++
			layers = append(layers, make([]int, 0))
		}

		layers[layer] = append(layers[layer], numbers[i])
	}

	minZeros := width * height
	onesByTwos := 0

	for i := 0; i < len(layers); i++ {
		zeros := 0
		ones := 0
		twos := 0

		for j := 0; j < len(layers[i]); j++ {
			if layers[i][j] == 0 {
				zeros++
			} else if layers[i][j] == 1 {
				ones++
			} else if layers[i][j] == 2 {
				twos++
			}
		}

		if zeros < minZeros {
			minZeros = zeros
			onesByTwos = ones * twos
		}
	}

	fmt.Println(onesByTwos)
}