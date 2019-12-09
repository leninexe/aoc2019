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
	
	black = 0
	white = 1
	transparent = 2
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

	var resultLayer [width * height]int

	for i := 0; i < width * height; i++ {
		result := layers[len(layers) - 1][i]

		for j := len(layers) - 2; j >= 0; j-- {
			if layers[j][i] == black || layers[j][i] == white {
				result = layers[j][i]
			}
		}

		resultLayer[i] = result
	}

	for i := 0; i < len(resultLayer); i++ {
		if resultLayer[i] == 0 {
			fmt.Print(" ")
		} else if resultLayer[i] == 1 {
			fmt.Print("#")
		} else {
			fmt.Print(" ")
		}

		if (i + 1) % width == 0 {
			fmt.Println()
		}
	}
}