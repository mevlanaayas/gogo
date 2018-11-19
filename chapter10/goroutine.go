
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main()  {
	/*
	goroutines to run programs concurrently
	we use go keyword to run functions in concurrent
	when you run the following program you will see
	the output of first function and the output of second
	function in mixed order.
	that means they runs concurrently

	note that: we call scanln function at the end
	why because:
	the main program runs one line and goes to another.
	It does not wait to finish a job for goroutines!!
	if you remove Scan line you will see it quits program immediately

	*/
	// first example
	go ConcurrentFunc(0)
	go ConcurrentFunc(1)
	var input string
	fmt.Print("Enter a string: ")
	fmt.Scanln(&input)

	// second example sync goroutines
	for i := 0; i < 10; i++ {
		go ConcurrentFunc(i)
	}
	fmt.Print("Enter a string: ")
	fmt.Scanln(&input)

}

func ConcurrentFunc(n int)  {
	/*
	in first example increase for loop limit to 50
	and comment out last two rounds of for loop
	*/
	for i := 0; i < 10; i++ {
		fmt.Println(n, " : ", i)
		/*
		following lines to sync goroutines
		*/
		amt := time.Duration(rand.Intn(250))
		time.Sleep(time.Millisecond * amt)
	}

}