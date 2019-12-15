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

	Ore = "ORE"
	Fuel = "FUEL"
)

type Reaction struct {
	input []RawMaterial
	output RawMaterial
}

type RawMaterial struct {
	amount int
	material string
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
	reactions := make([]Reaction, 0)

	for scanner.Scan() {
		line := scanner.Text()

		inputAndOutput := strings.Split(line, "=>")
		inputs := strings.Split(strings.Trim(inputAndOutput[0], " "), ",")
		output := strings.Split(strings.Trim(inputAndOutput[1], " "), " ")

		in := make([]RawMaterial, 0)

		for i := 0; i < len(inputs); i++ {
			splitter := strings.Split(strings.Trim(inputs[i], " "), " ")
			material := splitter[1]
			amount, _ := strconv.Atoi(splitter[0]) 

			in = append(in, RawMaterial{amount, material})
		}

		amount, _ := strconv.Atoi(output[0])

		out := RawMaterial{amount, output[1]}
		reactions = append(reactions, Reaction{in, out})
	}

	materials := make(map[string]int)
	needed := make(map[string]int)
	needed["FUEL"] = 1
	usedOre := 0

	for len(needed) > 0 {
		additionalNeeds := make(map[string]int)

		for material, amount := range needed {
			if materials[material] > 0 {
				available := materials[material]

				if available >= amount {
					amount = 0
					materials[material] -= amount
				} else {
					amount -= available
					materials[material] = 0
				}
			}

			for _, reaction := range reactions {
				if reaction.output.material == material {
					if len(reaction.input) == 1 && reaction.input[0].material == "ORE" {
						productions := amount / reaction.output.amount

						if amount % reaction.output.amount != 0 {
							productions++
						}

						usedOre += productions * reaction.input[0].amount
						materials[material] += productions * reaction.output.amount - amount
					} else {
						productions := amount / reaction.output.amount	

						if amount % reaction.output.amount != 0 {
							productions++
						}

						for _, input := range reaction.input {
							additionalNeeds[input.material] += productions * input.amount
						}

						materials[material] += (productions * reaction.output.amount) - amount
					}
				}
			}

			delete(needed, material)	
		}

		for k, v := range additionalNeeds {
			needed[k] += v
		}
	}

	fmt.Println(usedOre)
}