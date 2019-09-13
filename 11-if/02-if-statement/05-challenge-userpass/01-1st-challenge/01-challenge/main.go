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
)

// ---------------------------------------------------------
// CHALLENGE #1
//  Create a user/password protected program.
//
// EXAMPLE USER
//  username: jack
//  password: 1888
//
// EXPECTED OUTPUT
//  go run main.go
//    Usage: [username] [password]
//
//  go run main.go albert
//    Usage: [username] [password]
//
//  go run main.go hacker 42
//    Access denied for "hacker".
//
//  go run main.go jack 6475
//    Invalid password for "jack".
//
//  go run main.go jack 1888
//    Access granted to "jack".
// ---------------------------------------------------------

func main() {
	const (
		usage    = "Usage: [username] [password]"
		errUsr   = "Access denied for %q.\n"
		errPwd   = "Invalid password for %q.\n"
		accessOK = "Access granted to %q.\n"

		user   = "jack"
		passwd = "1888"
	)
	if len(os.Args) < 3 {
		fmt.Println(usage)
		return
	}

	usr, pwd := os.Args[1], os.Args[2]

	if usr != user {
		fmt.Printf(errUsr, usr)
	} else if pwd != passwd {
		fmt.Printf(errPwd, usr)
	} else {
		fmt.Printf(accessOK, usr)
	}
}
