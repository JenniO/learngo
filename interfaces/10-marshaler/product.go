// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import (
	"fmt"
	"strconv"
	"time"
)

type product struct {
	Title    string
	Price    money
	Released timestamp
}

func (p *product) String() string {
	return fmt.Sprintf("%s: %s (%s)", p.Title, p.Price, p.Released)
}

func (p *product) discount(ratio float64) {
	p.Price *= money(1 - ratio)
}

func toTimestamp(v interface{}) (ts timestamp) {
	var t int

	switch v := v.(type) {
	case int:
		// book{Title: "moby dick", Price: 10, published: 118281600},
		t = v
	case string:
		// book{Title: "odyssey", Price: 15, published: "733622400"},
		t, _ = strconv.Atoi(v)
	}

	ts.Time = time.Unix(int64(t), 0)
	return ts
}
