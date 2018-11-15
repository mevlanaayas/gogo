
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

	fmt.Println("-slice example-")
	sliceExample()

}

func sliceExample()  {
	/*
	Slices have length and index as arrays
	But the length is changeable
	*/
	// declaring slice. we don't put length
	var x []float64
	fmt.Println("Slice default ", x)

	// TODO: need explanation
	exSlice := make([]float64, 5)
	fmt.Println("Slice with make 5 ", exSlice)

	// TODO: need explanation
	exSliceWithMaxSize := make([]float64, 5, 10)
	fmt.Println("Slice with make 5 10 ", exSliceWithMaxSize)

	arr := [5]float64{ 1, 2, 3, 4, 5 }
	exSliceWitLowAndHigh := arr[0:5]
	fmt.Println("Slice with low:high assign ", exSliceWitLowAndHigh)

	/*
	different syntax to use [low:high expression]
	exSliceWitLowAndHigh := arr[:]
	exSliceWitLowAndHigh := arr[:5]
	exSliceWitLowAndHigh := arr[0:]
	*/

	// we have two built-in for slices: copy and append

	// append example
	sliceOne := []int{ 1, 2, 3 }

	/*
	run: godoc builtin append
		docs show us usage
		append(destination slice, new elements or another slice)
	*/
	sliceTwo := append(sliceOne, 4, 5)

	fmt.Println("Slice with append ", sliceTwo)

	//copy example
	sliceThree := []int{ 1, 2, 3 }
	sliceFour := make([]int, 2)

	/*
	run: godoc builtin copy
		docs show us usage
	    copy(destination, source)
	*/
	copy(sliceFour, sliceThree)

	// we can print multiple element at once
	fmt.Println(sliceThree, sliceFour)

	/*
	we will see there is only two of elements of sliceThree
	why because sliceFour has only two gap to get item.
	*/

}