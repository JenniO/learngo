// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import "fmt"

type printer interface {
	print()
}

type list []printer

func (l list) print() {
	// `list` acts like a `[]game`
	if len(l) == 0 {
		fmt.Println("Sorry. We're waiting for delivery 🚚.")
		return
	}

	for _, it := range l {
		it.print()
	}
}
