package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	lineNo := 0
	list := []string{}

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		t := scanner.Text()

		list = append(list, t)
		lineNo++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(checksum(list))
}

func checksum(list []string) int64 {
	var twos, threes int64

	for _, item := range list {
		hasTwo, hasThree := checkFor2or3(item)

		if hasTwo {
			twos++
		}

		if hasThree {
			threes++
		}
	}

	return twos * threes

}

func checkFor2or3(s string) (hasTwo bool, hasThree bool) {

	twos := 0
	threes := 0

	letters := map[rune]int{}

	for _, r := range s {
		letters[r]++

		if letters[r] == 2 {
			twos++
		}

		if letters[r] == 3 {
			twos--
			threes++
		}

		if letters[r] == 4 {
			threes--
		}
	}

	return twos > 0, threes > 0
}
