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
	"strconv"
)

// ---------------------------------------------------------
// EXERCISE: Sum up to N
//
//  1. Get two numbers from the command-line: min and max
//  2. Convert them to integers (using Atoi)
//  3. By using a loop, sum the numbers between min and max
//
// RESTRICTIONS
//  1. Be sure to handle the errors. So, if a user doesn't
//     pass enough arguments or she passes non-numerics,
//     then warn the user and exit from the program.
//
//  2. Also, check that, min < max.
//
// HINT
//  For converting the numbers, you can use `strconv.Atoi`.
//
// EXPECTED OUTPUT
//  Let's suppose that the user runs it like this:
//
//    go run main.go 1 10
//
//  Then it should print:
//
//    1 + 2 + 3 + 4 + 5 + 6 + 7 + 8 + 9 + 10 = 55
// ---------------------------------------------------------

func main() {
	if len(os.Args) < 3 {
		fmt.Println("I need you to give me two numbers")
		return
	}

	if min, err := strconv.Atoi(os.Args[1]); err != nil {
		fmt.Println(os.Args[1], "is not a number")
		return
	} else if max, err := strconv.Atoi(os.Args[2]); err != nil {
		fmt.Println(os.Args[2], "is not a number")
		return
	} else {
		sum := 0
		for ; min <= max; min++ {
			sum += min
			if min == max {
				fmt.Printf("%d = ", min)
				break
			}
			fmt.Printf("%d + ", min)
		}
		fmt.Print(sum)
	}

}
