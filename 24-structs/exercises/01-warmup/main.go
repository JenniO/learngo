// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Warm Up
//
//  Starting with this exercise, you'll build a command-line
//  game store.
//
//  1. Declare the following structs:
//
//     + item: id (int), name (string), price (int)
//
//     + game: embed the item, genre (string)
//
//
//  2. Create a game slice using the following data:
//
//     id  name          price    genre
//
//     1   god of war    50       action adventure
//     2   x-com 2       30       strategy
//     3   minecraft     20       sandbox
//
//
//  3. Print all the games.
//
// EXPECTED OUTPUT
//  Please run the solution to see the output.
// ---------------------------------------------------------

type item struct {
	id    int
	name  string
	price int
}

type game struct {
	item
	genre string
}

func main() {
	games := []game{
		{
			item: item{
				id:    1,
				name:  "god of war",
				price: 50,
			},
			genre: "action adventure",
		},
		{
			item: item{
				id:    2,
				name:  "x-com 2",
				price: 30,
			},
			genre: "strategy",
		},
		{
			item: item{
				id:    3,
				name:  "minecraft",
				price: 20,
			},
			genre: "sandbox",
		},
	}

	fmt.Printf("%-3s %-15s %-8s %-20s\n", "id", "name", "price", "genre")
	for _, game := range games {
		fmt.Printf("%-3d %-15s %-8d %-20s\n",
			game.id, game.name, game.price, game.genre)
	}
}
