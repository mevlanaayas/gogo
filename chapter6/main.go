
package main

import (
	"fmt"
)

func main()  {
	scores := []float64{ 90, 80, 70, 60, 50 }
	fmt.Println(average(scores))

	fmt.Println(returnHello())
	fmt.Println(returnMultiple())

	/*
		Variadic functions:
		pass input whose length is dynamic
	*/
	fmt.Println(variadicO(5, 90, 80, 70, 60, 50))

	/*
		Closure functions
		define function inside other function's scope
		so it can access local variables
	*/
	// local variable count to use in closure func
	count := 0
	// closure functions
	increment := func() int {
		count++
		return count
	}
	decrement := func() int {
		count--
		return count
	}
	// call
	fmt.Println(increment())
	fmt.Println(decrement())
	/*
	why we cant use ???
	func () (count int) {
		count ++
		return
	}
	*/

	/*
	need more detailed reading
	*/
	nextEven := makeEvenGeneratorByClosure()
	fmt.Println(nextEven())
	fmt.Println(nextEven())
	fmt.Println(nextEven())

	/*
		Recursive functions:
		function which calls itself
	*/
	fmt.Println(recursiveEx(6))

	/*
	starting to talk about defer, panic and recover
	*/
	fmt.Println("-Defer, Panic, Recover-")

	/*
	defer masterPiece()
	when we run this statement in any function.
	masterPiece function runs at the end of function
	*/
	deferTry()
	deferTryNo1st()

	/*
	recover and panic
	panic is to throw errors
	recover is to recover this errors :)

	*/
	panicRecoverEx()

	/*
		Pointers
	when we pass an variable to function
	we send a copy of that object not itself
	see the next lines and look at the functions itself
	*/
	withoutPointerEx()
	pointerEx()

	/*
	when we pass pointer to the function we change the value
	from address of variable not from the copy of it.
	if we print number inside both setZero and setPtrZero we
	see the value of number inside setZero and address of number inside
	setPtrZero

	pointers are represented with "*" (asterisk)
	and we use & operator to find address of variable

	we use asterisk to reach a value that address hold
	example:
		x = 10 // initialize x sets it to 10, put 10 address of x
		y = &x // put address of x inside y
		z = *y // finds value that y store. y is address now and that address holds 10
	*/

	fmt.Println("-new builtin-")
	newEx()

}

func average(scores []float64) float64  {
	result := 0.0
	total := 0.0
	for _, score := range scores {
		total += score
	}
	result = total / float64(len(scores))
	return result

}

func returnHello() (result string) {
	result = "Mevlana"
	return

}

func returnMultiple() (first int, second int) {
	first, second = 5, 10
	return

}

func variadicO(first int, args ...int) (result int)  {
	// args comes here as an array
	fmt.Println(args)
	for _, value := range args {
		result += value
	}
	result /= first
	return

}

func makeEvenGeneratorByClosure() func() uint  {
	i := uint(0)
	returnValue := func() (ret uint){
		ret = i
		i += 2
		return
	}
	return returnValue

}

func recursiveEx(number uint) uint {
	if number == 0 {
		return 1
	}
	return number * recursiveEx(number - 1)

}

func first()  {
	fmt.Println("1st")

}

func second()  {
	fmt.Println("2nd")

}
func deferTry()  {
	defer second()
	first()

}

func deferTryNo1st()  {
	defer second()
	/*
	we said the function next to defer keyword will be called
	when function ends. so if we return before line "first()"
	we expect to see "2nd" on screen and no "1st".
	*/
	return
	first()

}

func panicRecoverEx()  {
	defer func() {
		message := recover()
		fmt.Println("Panic message ", message)
	}()
	panic("OMG")

}

func pointerEx()  {
	variable := 100
	setPtrZero(&variable)
	fmt.Println("We are expecting to see 0 now ;) ", variable)

}

func withoutPointerEx()  {
	variable := 10
	setZero(variable)
	fmt.Println("We are excepting to see 10 again ", variable)

}

func setZero(number int)  {
	fmt.Println("expected to see value ", number)
	number = 0

}

func setPtrZero(number *int)  {
	fmt.Println("expected to see adress", number)
	*number = 0

}

func newEx()  {
	xPtr := new(int)
	*xPtr = 10
	fmt.Println("We are expecting to see address ", xPtr)
	fmt.Println("We are expecting to see value ", *xPtr)

}