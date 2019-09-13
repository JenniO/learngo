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
)

// ---------------------------------------------------------
// EXERCISE: Arg Count
//
//  1. Get arguments from command-line.
//  2. Print the expected outputs below depending on the number
//     of arguments.
//
// EXPECTED OUTPUT
//  go run main.go
//    Give me args
//
//  go run main.go hello
//    There is one: "hello"
//
//  go run main.go hi there
//    There are two: "hi there"
//
//  go run main.go i wanna be a gopher
//    There are 5 arguments
// ---------------------------------------------------------

func main() {
	n := len(os.Args)
	if n < 2 {
		fmt.Println("Give me args")
	} else if n == 2 {
		fmt.Printf("There is one: %q\n", os.Args[1])
	} else if n == 3 {
		fmt.Printf(`There are two: "%s %s"`, os.Args[1], os.Args[2])
	} else {
		fmt.Printf("There are %d arguments\n", n-1)
	}
}
