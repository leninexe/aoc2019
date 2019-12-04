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
	numbers := strings.Split(scanner.Text(), "-")

	from, _ := strconv.Atoi(numbers[0])
	to, _ := strconv.Atoi(numbers[1])

	fmt.Printf("From %d to %d\n", from, to)

	count := 0

	for i := from; i <= to; i++ {
		sixth := i % 10
		fivth := i % 100 / 10
		fourth := i % 1000 / 100
		third := i % 10000 / 1000
		second := i % 100000 / 10000
		first := i / 100000

		if first <= second && second <= third && third <= fourth && fourth <= fivth && fivth <= sixth {
			// !decreasing

			if first == second && second != third {
				count++
			} else if first != second && second == third && third != fourth {
				count++
			} else if second != third && third == fourth && fourth != fivth {
				count++
			} else if third != fourth && fourth == fivth && fivth != sixth {
				count++
			} else if fourth != fivth && fivth == sixth {
				count++
			}
		}
	}

	fmt.Println(count)
}