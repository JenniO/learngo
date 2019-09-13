package main

import "fmt"

type item struct {
	id    int
	name  string
	price int
}

type game struct {
	item
	genre string
}

func newGame(id, price int, name, genre string) game {
	g := game{
		item:  item{id: id, name: name, price: price},
		genre: genre,
	}
	return g
}

func addGame(games []game, g game) []game {
	games = append(games, g)
	return games
}

func load() (games []game) {
	games = addGame(games, newGame(1, 50, "god of war", "action adventure"))
	games = addGame(games, newGame(2, 40, "x-com 2", "strategy"))
	games = addGame(games, newGame(3, 20, "minecraft", "sandbox"))
	return
}

func indexByID(games []game) map[int]game {
	byID := make(map[int]game)
	for _, g := range games {
		byID[g.id] = g
	}
	return byID
}

func printGame(g game) {
	fmt.Printf("#%d: %-15q %-20s $%d\n",
		g.id, g.name, "("+g.genre+")", g.price)
}
