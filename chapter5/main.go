
package main

import "fmt"

func main()  {
	// arrays holds set of variables whose values are same
	var x [5]int
	x[4] = 100
	fmt.Println(x)

	// initializing array
	var exArray [5]float64
	exArray[0] = 90
	exArray[1] = 80
	exArray[2] = 70
	exArray[3] = 60
	exArray[4] = 50

	var total float64 = 0

	// 5 is hardcoded, for i := 0; i < 5; i++ {
	for i := 0; i < len(exArray); i++ {
		total += exArray[i]
	}
	fmt.Println(exArray)
	fmt.Println("Total ", total)

	// also here 5 is hardcoded, fmt.Println("Average ", total / 5)
	/*
	when you use 'fmt.Println("Average ", total / len(exArray))'
	total / len(exArray) throw an mismatched types error
	because len(exArray) returns integer and total is defined as float64
	convert len(exArray) to float64
		float64(len(exArray))
		This is an example of a type conversion. In general, to convert between types, you use
		the type name like a function.
	*/
	fmt.Println("Average ", total / float64(len(exArray)))

	// different syntax of for loop when we use arrays
	total = 0.0
	for i, value := range exArray {
		total += value
		fmt.Println("Index", i)
		fmt.Println("Value", value)
	}
	fmt.Println("Average ", total / float64(len(exArray)))

	/*
	I used 'i', which represents current index, to print to console.
	But other usage of for loop is shown below
	In this case I don't want to print index
	And program throws an error: declared and not used
	Because we don't use it inside for loop
	*/
	total = 0.0
	for _, value := range exArray {
		total += value
	}
	// I used _ (underscore). _ is to tell the compiler "I wont use this variable"
	// Ignored values let's say.

	fmt.Println("Average ", total / float64(len(exArray)))

	// We can use other declaration syntax for arrays
	exBulkArray := [5]float64{ 90, 80, 70, 60, 50 }

	// also this is possible
	exMultiLineBulkArray := [5]float64{
		90,
		80,
		70,
		60,
		50,
		// trailing comma is required here
	}

	fmt.Println(exBulkArray)
	fmt.Println(exMultiLineBulkArray)

}
