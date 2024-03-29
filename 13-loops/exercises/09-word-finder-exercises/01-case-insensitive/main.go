// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import (
	"fmt"
	"os"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Case Insensitive Search
//
//  Allow for case-insensitive searching
//
// EXAMPLE
//  Let's say that the user runs the program like this:
//    go run main.go LAZY
//
//  Or like this: go run main.go lAzY
//  Or like this: go run main.go lazy
//
//  For all cases above, the program should find
//  the "lazy" keyword.
// ---------------------------------------------------------
const corpus = "lazy cat jumps again and again and again"

func main() {
	query := os.Args[1:]
	words := strings.Fields(corpus)

	for _, r := range query {
		for i, w := range words {
			if strings.ToLower(r) == strings.ToLower(w) {
				fmt.Printf("#%-2d: %q\n", i+1, w)
				break
			}
		}
	}
}
