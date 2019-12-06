package main

import (
	"fmt"
	"os"
	"log"
	"bufio"
	"strings"
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

	scanner.Scan()
	
	input := scanner.Text()
	inputSlice := strings.Split(input, ",")
	code := make([]int, len(inputSlice))

	for i := 0; i < len(inputSlice); i++ {
		code[i], err = strconv.Atoi(inputSlice[i])
	}

	// Input starts with 1
	currentInput := 1

	for i := 0; i < len(code); i++ {
		opcode := code[i] % 100
		p1mode := (code[i] % 1000) / 100
		p2mode := (code[i] % 10000) / 1000

		if opcode == 1 {
			// Add

			p1 := code[i + 1]
			p2 := code[i + 2]

			if p1mode == 0 {
				p1 = code[p1]
			}

			if p2mode == 0 {
				p2 = code[p2]
			}

			code[code[i + 3]] = p1 + p2
			i += 3
		} else if opcode == 2 {
			// Multiply

			p1 := code[i + 1]
			p2 := code[i + 2]

			if p1mode == 0 {
				p1 = code[p1]
			}

			if p2mode == 0 {
				p2 = code[p2]
			}

			code[code[i + 3]] = p1 * p2
			i += 3
		} else if opcode == 3 {
			// Store value from input
			code[code[i + 1]] = currentInput

			i += 1
		} else if opcode == 4 {
			// Write to output
			p1 := code[i + 1]

			fmt.Printf("Output: code[%d] = %d\n", p1, code[p1])

			i += 1
		} else if opcode == 99 {
			break
		}
	}
}