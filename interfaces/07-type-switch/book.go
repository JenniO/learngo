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

type book struct {
	title     string
	price     money
	published interface{}
}

func (b book) print() {
	p := " - " + format(b.published)
	fmt.Printf("%-15s: %5s %s\n", b.title, b.price.string(), p)
}

func format(v interface{}) string {
	var t int
	switch v := v.(type) {
	case int:
		t = v
	case string:
		t, _ = strconv.Atoi(v)
	default:
		return "unknown"
	}
	const layout = "2006-01"

	u := time.Unix(int64(t), 0)
	return u.Format(layout)
}
