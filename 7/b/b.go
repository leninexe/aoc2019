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

	for a := 5; a < 10; a++ {
		for b := 5; b < 10; b++ {
			for c := 5; c < 10; c++ {
				for d := 5; d < 10; d++ {
					for e := 5; e < 10; e++ {
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
	// for p := 0; p < 1; p++ {	
		// phaseSignal 
		amps := perms[p]
		// amps = [5]int{9,8,7,6,5}

		// Initalize amps with memory and instruction pointer
		currentAmp := 0
		var ampMemory [5][]int
		ampIp := [5]int{0, 0, 0, 0, 0}

		for amp := 0; amp < 5; amp++ {
			ampMemory[amp] = make([]int, len(programCode))

			for i := 0; i < len(programCode); i++ {
				ampMemory[amp][i] = programCode[i]
			}
		}

		firstInput := [5]bool{true, true, true, true, true}
		currentOutput := 0	
		halting := false

		for !halting {
			phaseSignal := amps[currentAmp]

			for i := ampIp[currentAmp]; i < len(ampMemory[currentAmp]); i++ {
				opcode := ampMemory[currentAmp][i] % 100
				p1mode := (ampMemory[currentAmp][i] % 1000) / 100
				p2mode := (ampMemory[currentAmp][i] % 10000) / 1000

				if opcode == 1 {
					// Add

					p1 := ampMemory[currentAmp][i + 1]
					p2 := ampMemory[currentAmp][i + 2]

					if p1mode == 0 {
						p1 = ampMemory[currentAmp][p1]
					}

					if p2mode == 0 {
						p2 = ampMemory[currentAmp][p2]
					}

					ampMemory[currentAmp][ampMemory[currentAmp][i + 3]] = p1 + p2
					i += 3
				} else if opcode == 2 {
					// Multiply

					p1 := ampMemory[currentAmp][i + 1]
					p2 := ampMemory[currentAmp][i + 2]

					if p1mode == 0 {
						p1 = ampMemory[currentAmp][p1]
					}

					if p2mode == 0 {
						p2 = ampMemory[currentAmp][p2]
					}

					ampMemory[currentAmp][ampMemory[currentAmp][i + 3]] = p1 * p2
					i += 3
				} else if opcode == 3 {
					// Store value from input
					if firstInput[currentAmp] {
						firstInput[currentAmp] = false
						ampMemory[currentAmp][ampMemory[currentAmp][i + 1]] = phaseSignal
					} else {
						ampMemory[currentAmp][ampMemory[currentAmp][i + 1]] = currentOutput
					}

					i += 1
				} else if opcode == 4 {
					// Write to output
					currentOutput = ampMemory[currentAmp][ampMemory[currentAmp][i + 1]]

					ampIp[currentAmp] = i + 2
					currentAmp = (currentAmp + 1) % 5

					break
				} else if opcode == 5 {
					// Jump if true
					p1 := ampMemory[currentAmp][i + 1]
					p2 := ampMemory[currentAmp][i + 2]

					if p1mode == 0 {
						p1 = ampMemory[currentAmp][p1]
					}

					if p2mode == 0{
						p2 = ampMemory[currentAmp][p2]
					}

					if p1 != 0 {
						i = p2 - 1
					} else {
						i += 2
					}
				} else if opcode == 6 {
					// Jump if false
					p1 := ampMemory[currentAmp][i + 1]
					p2 := ampMemory[currentAmp][i + 2]

					if p1mode == 0 {
						p1 = ampMemory[currentAmp][p1]
					}

					if p2mode == 0{
						p2 = ampMemory[currentAmp][p2]
					}

					if p1 == 0 {
						i = p2 - 1
					} else {
						i += 2
					}
				} else if opcode == 7 {
					// Less than
					p1 := ampMemory[currentAmp][i + 1]
					p2 := ampMemory[currentAmp][i + 2]

					if p1mode == 0 {
						p1 = ampMemory[currentAmp][p1]
					}

					if p2mode == 0{
						p2 = ampMemory[currentAmp][p2]
					}

					if p1 < p2 {
						ampMemory[currentAmp][ampMemory[currentAmp][i + 3]] = 1
					} else {
						ampMemory[currentAmp][ampMemory[currentAmp][i + 3]] = 0
					}

					i += 3
				} else if opcode == 8 {
					// Equals
					p1 := ampMemory[currentAmp][i + 1]
					p2 := ampMemory[currentAmp][i + 2]

					if p1mode == 0 {
						p1 = ampMemory[currentAmp][p1]
					}

					if p2mode == 0{
						p2 = ampMemory[currentAmp][p2]
					}

					if p1 == p2 {
						ampMemory[currentAmp][ampMemory[currentAmp][i + 3]] = 1
					} else {
						ampMemory[currentAmp][ampMemory[currentAmp][i + 3]] = 0
					}

					i += 3
				} else if opcode == 99 {
					halting = true

					break
				}
			}
		}

		if currentOutput > maxResult {
			maxResult = currentOutput
			bestSettings = amps
		}
	}

	fmt.Println(bestSettings)
	fmt.Println(maxResult)
}