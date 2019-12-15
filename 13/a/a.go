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

	positionMode = 0
	immediateMode = 1
	relativeMode = 2

	opAdd = 1 // 3 param
	opMultiply = 2 // 3 param
	opRead = 3 // 1 param
	opWrite = 4	// 1 param
	opJumpIfTrue = 5 // 2 param
	opJumpIfFalse = 6 // 2 param
	opLessThan = 7 // 3 param
	opEquals = 8 // 3 param
	opSetRelativeBase = 9 // 1 param
	opEof = 99 // 0 param

	empty = 0
	wall = 1
	block = 2
	horizontalPaddle = 3
	ball = 4
)

type Point struct {
	x int
	y int
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

	scanner.Scan()
	
	input := scanner.Text()
	inputSlice := strings.Split(input, ",")
	code := make(map[int]int)

	for i := 0; i < len(inputSlice); i++ {
		code[i], err = strconv.Atoi(inputSlice[i])
	}

	pc := 0
	relativeBase := 0
	eof := false

	outputPointer := 0
	output := make([]int, 3)
	board := make(map[Point]int)

	for !eof {
		op, p1Mode, p2Mode, p3Mode := getOperation(code[pc])

		p1 := code[pc + 1]
		p2 := code[pc + 2]
		p3 := code[pc + 3]

		switch op {
		case opAdd:
			if p1Mode == positionMode {
				p1 = code[p1]
			} else if p1Mode == immediateMode {
			} else if p1Mode == relativeMode {
				p1 = code[relativeBase + p1]
			}

			if p2Mode == positionMode {
				p2 = code[p2]
			} else if p2Mode == immediateMode {
			} else if p2Mode == relativeMode {
				p2 = code[relativeBase + p2]
			}

			if p3Mode == positionMode {
			} else if p3Mode == relativeMode {
				p3 += relativeBase
			}

			code[p3] = p1 + p2
			pc += 4
		case opMultiply:
			if p1Mode == positionMode {
				p1 = code[p1]
			} else if p1Mode == immediateMode {
			} else if p1Mode == relativeMode {
				p1 = code[relativeBase + p1]
			}

			if p2Mode == positionMode {
				p2 = code[p2]
			} else if p2Mode == immediateMode {
			} else if p2Mode == relativeMode {
				p2 = code[relativeBase + p2]
			}

			if p3Mode == positionMode {
			} else if p3Mode == relativeMode {
				p3 += relativeBase
			}

			code[p3] = p1 * p2
			pc += 4
		case opRead:
			if p1Mode == positionMode {
			} else if p1Mode == relativeMode {
				p1 += relativeBase
			}

			// code[p1] = currentInput

			pc += 2
		case opWrite:
			if p1Mode == positionMode {
				p1 = code[p1]
			} else if p1Mode == immediateMode {
			} else if p1Mode == relativeMode {
				p1 = code[relativeBase + p1]
			}

			output[outputPointer] = p1
			outputPointer++

			if outputPointer == 3 {
				pt := Point{output[0], output[1]}
				id := output[2]

				// existing := board[pt]

				// switch existing {
				// case empty:
				// case wall:
				// case block:
				// case horizontalPaddle:
				// case ball:
				// }

				board[pt] = id

				outputPointer = 0
			}

			pc += 2
		case opJumpIfTrue:
			if p1Mode == positionMode {
				p1 = code[p1]
			} else if p1Mode == immediateMode {
			} else if p1Mode == relativeMode {
				p1 = code[relativeBase + p1]
			}

			if p2Mode == positionMode {
				p2 = code[p2]
			} else if p2Mode == immediateMode {
			} else if p2Mode == relativeMode {
				p2 = code[relativeBase + p2]
			}

			if p1 != 0 {
				pc = p2
			} else {
				pc += 3
			}
		case opJumpIfFalse: 
			if p1Mode == positionMode {
				p1 = code[p1]
			} else if p1Mode == immediateMode {
			} else if p1Mode == relativeMode {
				p1 = code[relativeBase + p1]
			}

			if p2Mode == positionMode {
				p2 = code[p2]
			} else if p2Mode == immediateMode {
			} else if p2Mode == relativeMode {
				p2 = code[relativeBase + p2]
			}

			if p1 == 0 {
				pc = p2
			} else {
				pc += 3
			}
		case opLessThan:
			if p1Mode == positionMode {
				p1 = code[p1]
			} else if p1Mode == immediateMode {
			} else if p1Mode == relativeMode {
				p1 = code[relativeBase + p1]
			}

			if p2Mode == positionMode {
				p2 = code[p2]
			} else if p2Mode == immediateMode {
			} else if p2Mode == relativeMode {
				p2 = code[relativeBase + p2]
			}

			if p3Mode == positionMode {
			} else if p3Mode == relativeMode {
				p3 += relativeBase
			}

			if p1 < p2 {
				code[p3] = 1
			} else {
				code[p3] = 0
			}

			pc += 4
		case opEquals:
			if p1Mode == positionMode {
				p1 = code[p1]
			} else if p1Mode == immediateMode {
			} else if p1Mode == relativeMode {
				p1 = code[relativeBase + p1]
			}

			if p2Mode == positionMode {
				p2 = code[p2]
			} else if p2Mode == immediateMode {
			} else if p2Mode == relativeMode {
				p2 = code[relativeBase + p2]
			}

			if p3Mode == positionMode {
			} else if p3Mode == relativeMode {
				p3 += relativeBase
			}

			if p1 == p2 {
				code[p3] = 1
			} else {
				code[p3] = 0
			}

			pc += 4
		case opSetRelativeBase:
			if p1Mode == positionMode {
				p1 = code[p1]
			} else if p1Mode == immediateMode {
			} else if p1Mode == relativeMode {
				p1 = code[relativeBase + p1]
			}

			relativeBase += p1
			pc += 2
		case opEof:
			eof = true
		}
	}

	count := 0

	for _, v := range board {
		if v == block {
			count++
		}
	}

	fmt.Println(count)
}

func getOperation(cmd int) (int, int, int, int) {
	op := cmd % 100
	p1Mode := (cmd % 1000) / 100
	p2Mode := (cmd % 10000) / 1000
	p3Mode := cmd / 10000

	return op, p1Mode, p2Mode, p3Mode
}