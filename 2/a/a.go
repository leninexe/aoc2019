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
	test := false

	if len(os.Args) > 1 && os.Args[1] == TestParam {
		test = true
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

	// pos[i] == 1 => pos[i + 3] = pos[i + 1] + pos[i + 2]
	// pos[i] == 2 => pos[i + 3] = pos[i + 1] * pos[i + 2]
	// pos[i] == 99 => EOF 

	// restore 1202 state only if not in test
	if !test {
		code[1] = 12
		code[2] = 2
	}

	for i := 0; i < len(code); i++ {
		if code[i] == 1 {
			code[code[i + 3]] = code[code[i + 1]] + code[code[i + 2]]
		} else if code[i] == 2 {
			code[code[i + 3]] = code[code[i + 1]] * code[code[i + 2]]
		} else if code[i] == 99 {
			break
		}

		fmt.Println(code)

		i += 3
	}

	fmt.Println(code[0])
}