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

	fmt.Println(findLetters(list))
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

func compare(a, b string) int {

	diff := -1

	uA := []rune(a)
	uB := []rune(b)

	if len(a) != len(b) {
		return -1
	}

	for i, r := range uA {

		if uB[i] == r {
			continue
		}

		if diff != -1 {
			return -1
		}

		diff = i
	}

	return diff
}

func findLetters(list []string) string {
	for i, word := range list {
		for j := i + 1; j < len(list); j++ {
			diff := compare(word, list[j])

			if diff != -1 {
				uWord := []rune(word)
				return string(append(uWord[:diff], uWord[diff+1:]...))
			}
		}
	}

	return ""
}
