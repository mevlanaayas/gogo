
package main

import (
	"fmt"
	"os"
)

/*
function without args and return type
has only one statement
*/
func main() {
	fmt.Println("Hello, World")
	// exit code 0 means Success. it is an unnecessary unless it is answer of q4.
	os.Exit(0)
}

/*
run:
	go run chapter1/main.go (linux) or chapter1\main.go (windows)
    to see output of program. in our case it will be Hello, World in terminal

build:
    go build chapter1/main.go (linux) or chapter1\main.go (windows)
    to create executable.
 */

 /*
 answers:
 4-) import os, and use os.Exit(code) statement. Use godoc os Exit to see details
 */