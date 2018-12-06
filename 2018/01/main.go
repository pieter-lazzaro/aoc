package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

var input string

func main() {
	var frequency int64
	frequencies := map[int64]bool{}
	inputFrequencies := []int64{}

	lineNo := 0

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		t := scanner.Text()

		i, err := strconv.ParseInt(t, 10, 64)

		if err != nil {
			log.Fatalf("could not parse value %s on line %d", t, lineNo)
		}

		inputFrequencies = append(inputFrequencies, i)

		frequency = frequency + i

		frequencies[frequency] = true

		lineNo++
	}

	for {
		for _, i := range inputFrequencies {

			frequency = frequency + i
			if frequencies[frequency] {
				fmt.Printf("Reached %d twice", frequency)
				os.Exit(0)
			}

		}
	}

}
