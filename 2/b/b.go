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
	// TestParam = "-t"
	Inputfile = "input.txt"
	// TestInputfile = "testinput.txt"
	Add = 1
	Multiply = 2
	EOF = 99
	ExitCode = 19690720
)

func main () {
	filename := Inputfile

	// if len(os.Args) > 1 && os.Args[1] == TestParam {
	// 	filename = TestInputfile
	// }

	file, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	
	input := scanner.Text()
	inputSlice := strings.Split(input, ",")
	originalCode := make([]int, len(inputSlice))

	for i := 0; i < len(inputSlice); i++ {
		originalCode[i], err = strconv.Atoi(inputSlice[i])
	}

	// pos[i] == 1 => pos[i + 3] = pos[i + 1] + pos[i + 2]
	// pos[i] == 2 => pos[i + 3] = pos[i + 1] * pos[i + 2]
	// pos[i] == 99 => EOF 

	foundSolution := false

	for noun := 0; noun <= 99 && !foundSolution; noun++ {
		for verb := 0; verb <= 99 && !foundSolution; verb++ {
			code := make([]int, len(originalCode))

			for i:= 0; i < len(originalCode); i++ {
				code[i] = originalCode[i]
			}

			code[1] = noun
			code[2] = verb

			for i := 0; i < len(code); i++ {
				if code[i] == Add {
					code[code[i + 3]] = code[code[i + 1]] + code[code[i + 2]]
				} else if code[i] == Multiply {
					code[code[i + 3]] = code[code[i + 1]] * code[code[i + 2]]
				} else if code[i] == EOF {
					break
				}

				i += 3
			}

			if code[0] == ExitCode {
				foundSolution = true

				fmt.Printf("Found ExitCode with noun = %d and verb = %d, Solution is %d\n", noun, verb, 100 * noun + verb)
			}
		}
	}
}