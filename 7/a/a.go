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

// Phase settings 0-4 (each only used once)

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
	programCode := make([]int, len(inputSlice))

	for i := 0; i < len(inputSlice); i++ {
		programCode[i], err = strconv.Atoi(inputSlice[i])
	}

	// Permutations
	var perms [120][5]int
	count := 0

	for a := 0; a < 5; a++ {
		for b := 0; b < 5; b++ {
			for c := 0; c < 5; c++ {
				for d := 0; d < 5; d++ {
					for e := 0; e < 5; e++ {
						if a != b && a != c && a != d && a != e && b != c && b != d && b != e &&
							c != d && c != e && d != e {
							perms[count] = [5]int{a, b, c, d, e}
							count++
						}
					}
				}
			} 
		}
	}

	var bestSettings [5]int
	maxResult := 0

	for p := 0; p < len(perms); p++ {
		// phaseSignal 
		amps := perms[p]

		// inputSignal
		currentInput := 0

		for amp := 0; amp < 5; amp++ {
			// Copy code
			code := make([]int, len(programCode))

			for i := 0; i < len(programCode); i++ {
				code[i] = programCode[i]
			}

			phaseSignal := amps[amp]

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
					code[code[i + 1]] = phaseSignal
					phaseSignal = currentInput

					i += 1
				} else if opcode == 4 {
					// Write to output
					currentInput = code[code[i + 1]]

					i += 1
				} else if opcode == 5 {
					// Jump if true
					p1 := code[i + 1]
					p2 := code[i + 2]

					if p1mode == 0 {
						p1 = code[p1]
					}

					if p2mode == 0{
						p2 = code[p2]
					}

					if p1 != 0 {
						i = p2 - 1
					} else {
						i += 2
					}
				} else if opcode == 6 {
					// Jump if false
					p1 := code[i + 1]
					p2 := code[i + 2]

					if p1mode == 0 {
						p1 = code[p1]
					}

					if p2mode == 0{
						p2 = code[p2]
					}

					if p1 == 0 {
						i = p2 - 1
					} else {
						i += 2
					}
				} else if opcode == 7 {
					// Less than
					p1 := code[i + 1]
					p2 := code[i + 2]

					if p1mode == 0 {
						p1 = code[p1]
					}

					if p2mode == 0{
						p2 = code[p2]
					}

					if p1 < p2 {
						code[code[i + 3]] = 1
					} else {
						code[code[i + 3]] = 0
					}

					i += 3
				} else if opcode == 8 {
					// Equals
					p1 := code[i + 1]
					p2 := code[i + 2]

					if p1mode == 0 {
						p1 = code[p1]
					}

					if p2mode == 0{
						p2 = code[p2]
					}

					if p1 == p2 {
						code[code[i + 3]] = 1
					} else {
						code[code[i + 3]] = 0
					}

					i += 3
				} else if opcode == 99 {
					break
				}
			}
		}

		if currentInput > maxResult {
			bestSettings = perms[p]
			maxResult = currentInput
		}
	}

	fmt.Println(bestSettings)
	fmt.Println(maxResult)
}