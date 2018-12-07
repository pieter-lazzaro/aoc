package main

import (
	"testing"
)

func Test_Checksum(t *testing.T) {

	inputList := []string{
		"abcdef",
		"bababc",
		"abbcde",
		"abcccd",
		"aabcdd",
		"abcdee",
		"ababab",
	}

	check := checksum(inputList)

	if check != 12 {
		t.Fatalf("expected %d to equal 12", check)
	}
}

func Test_Has2Or3(t *testing.T) {
	tests := []struct {
		input    string
		hasTwo   bool
		hasThree bool
	}{
		{input: "abcdef", hasTwo: false, hasThree: false},
		{input: "bababc", hasTwo: true, hasThree: true},
		{input: "abbcde", hasTwo: true, hasThree: false},
		{input: "abcccd", hasTwo: false, hasThree: true},
		{input: "aabcdd", hasTwo: true, hasThree: false},
		{input: "abcdee", hasTwo: true, hasThree: false},
		{input: "ababab", hasTwo: false, hasThree: true},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			hasTwo, hasThree := checkFor2or3(test.input)

			if hasTwo != test.hasTwo || hasThree != test.hasThree {
				t.Fatalf("expected %s to be %t:%t got %t:%t", test.input, test.hasTwo, test.hasThree, hasTwo, hasThree)
			}
		})
	}
}
