// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

func main() {
	var (
		mobydick  = book{title: "moby dick", price: 10}
		minecraft = game{title: "minecraft", price: 20}
		tetris    = game{title: "tetris", price: 5}
		rubik     = puzzle{title: "rubik", price: 5}
		yoda      = toy{title: "yoda", price: 10}
	)

	var store list
	store = append(store, &minecraft, &tetris, mobydick, rubik, &yoda)
	// store.print()

	var p printer

}
