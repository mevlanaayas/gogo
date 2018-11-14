
package main

import "fmt"

func main(){
	var x string = "Hello, World"
	// this may also work
	// var x = "Hello, World"
	fmt.Println(x)

	// different assign syntax
	var y string
	y = "Hello, World"
	fmt.Println(y)

	// changing value of string along lifetime
	var z string
	z = "first value"
	fmt.Println(z)
	z = z + " is over now"
	// z += "is over now" -> does the same job with above
	fmt.Println(z)

	// usage of ==
	fmt.Println(x == y)

	// different variable creation
	var a string = "new var with var and type"
	var b = "new var with var and without type"
	c := "new var without var and type"
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	// variable naming
	myName := "Mevlana"
	fmt.Println("My name is ", myName)

	first()
	second()

	exampleProgramFromBook()
}



// about scope
// !! in this case we write "var" to start!!
var outOfScopeVariable = "I am out of scope, so you can access me anywhere"

func first(){
	withinScopeVariable := "I am a variable of first, stay away from me"
	fmt.Println("Inside first func ", outOfScopeVariable)
	fmt.Println("Inside first func ", withinScopeVariable)
}

func second(){
	fmt.Println("Inside second func ", outOfScopeVariable)
	// we got "undefined: withinScopeVariable" error at next line
	// fmt.Println("Inside second func ", withinScopeVariable)
}

/*
Basically, the variable exists within the nearest curly braces ({ }),
or block, including any nested curly braces (blocks), but not outside of them.
*/



func constanExample(){
	const constant  = "Hello, World"
	fmt.Println(constant)
	// we got "cannot assign to constant" error at next line
	// constant = "Some other string"

}

/*
Basically, constants are variables whose values can not be changed
*/



func exampleProgramFromBook()  {
	// print is a other version of Println. It does not print new line
	fmt.Print("Enter a number: ")
	var input float64
	fmt.Scanf("%f", &input)

	output := input * 2

	fmt.Println("Example output ", output)

	// answer for question 5
	celsius := (input -32) * 5/9
	fmt.Println("Celsius result ", celsius)

	// answer for question 6
	meter := input * 1000 / 3048
	fmt.Println("Meter result ", meter)
}

/*
answers:
1-)
	var x string = ""
	var x = ""
    x := ""
2-)
    6
3-)
	scope is the living area of variable.
	its border is curly braces which variable is created in.
4-)
	var 	: changeable variables
	const 	: constant variables

5-)
	codes are above
6-)
	codes are above
*/
