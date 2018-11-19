
package main

import (
	"fmt"
	"time"
)

func main()  {
	/*
	we created two channel and send "from 1" and "from 2" messages every 2 and 3 seconds respectively.
	note that they are all running as goroutines

	in last go routines we use select statement
	it works like switch statement but special to channels
	it selects a channel which is ready
	*/

	channelOne := make(chan string)
	channelTwo := make(chan string)

	go func() {
		for {
			channelOne <- "from 1"
			time.Sleep(time.Second * 2)
		}

	}()

	go func() {
		for {
			channelTwo <- "from 2"
			time.Sleep(time.Second * 3)
		}

	}()

	go func() {
		var i = 0
		for {
			select {
			case msg1 := <- channelOne:
				fmt.Println("message from one: ", msg1)
				i++
			case msg2 := <- channelTwo:
				fmt.Println("message from two: ", msg2)
				i++
			/*
			we can use timeout case without default case.
			following two lines waits for second. if it does not take any
			value from channelOne or channelTwo within 1 sec it prints timeout.

			case <- time.After(time.Second):
				fmt.Println("Timeout. ")
			*/
			default:
				fmt.Println("nothing is ready: ", i)
			}
		}

	}()

	var input string
	fmt.Scanln(&input)

}
