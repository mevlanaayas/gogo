
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

