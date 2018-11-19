
package main

import (
	"fmt"
	"time"
)

func main()  {
	fmt.Println("ch")

	// var c chan string = make(chan string)
	var c = make(chan string)

	go pinger(c)
	go ponger(c)
	go printer(c)

	fmt.Print("enter a string: ")
	var input string
	fmt.Scanln(&input)

}

func pinger(c chan string)  {
	/*
	note that we declared infinite loop here
	for i:=0; ; i++ {}

	remember this was also for creating infinite loops
	for {}
	*/

	/*
	c <- "something" means to send something to channel
	"something" <- c means to receive something from channel

	we can define functions as receive or send only (channel direction - default is bidirectional)
	check ponger and printer functions.
	ponger is send only
	printer is retrieve only
	*/
	for i := 0; ; i++ {
		c <- "ping"
	}

}

func ponger(c chan<- string)  {
	for i := 0; ; i++ {
		/*
		we define ponger as send only and the
		following line will be resulted with compile error
		str := <- c
		*/
		c <- "pong"
	}

}

func printer(c <-chan string)  {
	for {
		msg := <- c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}

}

