
package main

import (
	"flag"
	"fmt"
	"math/rand"
)

func main()  {
	// command line arguments
	maxp := flag.Int("max", 6, "the max value")

	flag.Parse()

	fmt.Println("all args: ", flag.Args())
	fmt.Println(rand.Intn(*maxp))

}