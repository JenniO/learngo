// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import (
	"fmt"
	"sort"
)

func main() {
	l := list{
		{title: "moby dick", price: 10, released: toTimestamp(118281600)},
		{title: "odyssey", price: 15, released: toTimestamp("733622400")},
		{title: "hobbit", price: 25},
	}

	l.discount(.5)
	//sort.Sort(sort.Reverse(l))
	//sort.Sort(byReleaseDate(l))
	sort.Sort(sort.Reverse(byReleaseDate(l)))
	fmt.Print(l)

	// t := &toy{product{"yoda", 150}}
	// fmt.Printf("%#v\n", t)

	// b := {"moby dick", 10}, 118281600}
	// fmt.Printf("%#v\n", b)
}
