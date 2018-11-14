
package main

import "fmt"

func main()  {
	// Go has only one that is "for"
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i += 1
	}

	// be careful i has hold its final value
	fmt.Println("outside of loop", i)

	// different syntax
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// be careful i has hold its final value
	fmt.Println("outside of loop", i)

}
