// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

func main() {
	store := list{
		book{title: "moby dick", price: 10, published: 118281600},
		book{title: "odyssey", price: 15, published: "564689451"},
		book{title: "hobbit", price: 25},
		&game{title: "minecraft", price: 20},
		&game{title: "tetris", price: 5},
		puzzle{title: "rubik", price: 5},
		&toy{title: "yoda", price: 10},
	}

	store.discount(0.8)
	store.print()

	var p printer

}
