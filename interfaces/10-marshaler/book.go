// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

type book struct {
	// embed the product type into the book type.
	// all the fields and methods of the product will be
	// available in this book type.
	product
}
