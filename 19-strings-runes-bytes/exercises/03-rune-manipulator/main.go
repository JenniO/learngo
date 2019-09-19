package main

import (
	"fmt"
	"unicode/utf8"
)

// ---------------------------------------------------------
// EXERCISE: Rune Manipulator
//
//  Please read the comments inside the following code.
//
// EXPECTED OUTPUT
//  Please run the solution.
// ---------------------------------------------------------

func main() {
	strings := []string{
		"cool",
		"güzel",
		"jīntiān",
		"今天",
		"read 🤓",
	}

	_ = strings

	// Print the byte and rune length of the strings
	// Hint: Use len and utf8.RuneCountInString
	for _, s := range strings {
		fmt.Printf("Byte length: %-3d, rune length: %-3d\n", len(s),
			utf8.RuneCountInString(s))
	}

	// Print the bytes of the strings in hexadecimal
	// Hint: Use % x verb

	// Print the runes of the strings in hexadecimal
	// Hint: Use % x verb

	// Print the runes of the strings as rune literals
	// Hint: Use for range

	// Print the first rune and its byte size of the strings
	// Hint: Use utf8.DecodeRuneInString

	// Print the last rune of the strings
	// Hint: Use utf8.DecodeLastRuneInString

	// Slice and print the first two runes of the strings

	// Slice and print the last two runes of the strings

	// Convert the string to []rune
	// Print the first and last two runes
}
