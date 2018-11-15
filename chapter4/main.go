
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

	// if else statements
	for i := 1; i <= 5; i++ {
		if i % 2 == 0 {
			fmt.Println("even")
		} else {
			fmt.Println("odd")
		}
	}

	// if else if statements
	for i := 1; i <= 10; i++ {
		if i % 2 == 0 {
			fmt.Println("divisible by two")
		}else if i % 3 == 0 {
			fmt.Println("divisible by three")
		}else if i % 4 == 0 {
			fmt.Println("divisible by four")
		}
	}
}
