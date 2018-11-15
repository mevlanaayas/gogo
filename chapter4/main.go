
package main

import "fmt"

func main()  {
	// Go has only one that is "for"
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i += 1
	}

	// be careful i hold its final value
	fmt.Println("outside of loop", i)

	// different syntax
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// be careful i hold its final value
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
		} else if i % 3 == 0 {
			fmt.Println("divisible by three")
		} else if i % 4 == 0 {
			fmt.Println("divisible by four")
		}
	}

	fmt.Print("Enter a number between 0 and 10: ")
	fmt.Scanf("%d", &i)

	/*
	if
	*/
	if i == 0 {
		fmt.Println("Zero")
	} else if i == 1 {
		fmt.Println("One")
	} else if i == 2 {
		fmt.Println("Two")
	} else if i == 3 {
		fmt.Println("Three")
	} else if i == 4 {
		fmt.Println("Four")
	} else if i == 5 {
		fmt.Println("Five")
	} else {
		fmt.Println("Unknown Number")
	}

	/*
	switch
	*/
	switch i {
	case 0:
		fmt.Println("Zero")
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	case 4:
		fmt.Println("Four")
	case 5:
		fmt.Println("Five")
	default:
		fmt.Println("Unknown Number")
	}

	// calling answer func for second question
	answerTwo()

	// calling answer func for third question
	answerThree()
}

func answerTwo()  {
	for i := 1; i <= 100; i++ {
		if i % 3 == 0 {
			fmt.Println(i)
		}
	}
}

func answerThree()  {
	for i := 1; i <= 100; i++ {
		if i % 15 == 0 {
			fmt.Println("FizzBuzz ", i)
		} else if i % 3 == 0 {
			fmt.Println("Fizz ", i)
		} else if i % 5 == 0 {
			fmt.Println("Buzz ", i)
		} else {
			fmt.Println("Unhandled case ", i)
		}
	}
}

/*
answers:
1-)
	Big
2-)
    Codes are above
3-)
	Codes are above
*/
