
package main

import (
	"fmt"
)

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
	fmt.Println("Answer for Q3 ", exBulkArray[2:4])

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

	fmt.Println("-map example-")
	mapExample()

	fmt.Println("-answers-")
	answerFour()

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
	fmt.Println("Slice with make 5 10 and length of it ", exSliceWithMaxSize, len(exSliceWithMaxSize))

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

func mapExample()  {
	// this declaration causes to runtime error
	// var x map[string]int
	// so we use "make",
	x := make(map[string]int)
	x["key"] = 10
	fmt.Println(x, x["key"])

	/*
	we used strings as a key of map
	but we can also use integers
	*/
	mapWithIntegerKey := make(map[int]int)
	mapWithIntegerKey[10] = 101
	fmt.Println(mapWithIntegerKey)

	// maybe integer key and string value
	mapWithIntegerKeyAndStringValue := make(map[int]string)
	mapWithIntegerKeyAndStringValue[10] = "CapturedFish"
	fmt.Println(mapWithIntegerKeyAndStringValue)

	/*
	maps have non-static length. when we add new key:value pairs
	the length of map increases
	also we can use builtin delete function to delete key:value pair
	*/
	mapWithIntegerKeyAndStringValue[20] = "WithSea"
	fmt.Println("Before Delete 20 ", mapWithIntegerKeyAndStringValue)
	delete(mapWithIntegerKeyAndStringValue, 20)
	fmt.Println("After Delete 20 ", mapWithIntegerKeyAndStringValue)

	/*
	getting a bit more details
	*/
	fmt.Println("-maps continue-")
	elements := make(map[string]string)
	elements["H"] = "Hydrogen"
	elements["He"] = "Helium"
	elements["Li"] = "Lithium"
	elements["Be"] = "Beryllium"
	elements["B"] = "Boron"
	elements["C"] = "Carbon"
	elements["N"] = "Nitrogen"
	elements["O"] = "Oxygen"
	elements["F"] = "Fluorine"
	elements["Ne"] = "Neon"

	fmt.Println(elements)
	fmt.Println("Li", elements["Li"])

	// now lets try to print something is not in our map
	fmt.Println("Who is not in map", elements["Mevlana"])
	/*
	and the output seems strange, nothing next to our prompt.
	lets try to figure out
	*/
	name, result := elements["Mevlana"]
	fmt.Println("Results: ", name, result)
	/*
	we saw again a empty string (xd) but now also false
	if key does not exist in map
	it returns false (according to the type of value) and false (IsSuccessful)
	string: "" (empty string)
	integer: 0
	boolean: false
	etc.
	lets try all of them
	*/
	// we try string above so boolean now
	boolValues := map[string]bool{
		"He": true,
		"No": false,
		"Sp": true,
	}
	fmt.Println("Boolean values ", boolValues)
	boolValue, isSuccessfulBool := boolValues["Mevlana"]
	fmt.Println(boolValue, isSuccessfulBool)

	// and integer now
	intValues := map[string]int{
		"Zero": 12,
		"Sem": 45,
		"Mev": 12,
	}
	fmt.Println("Integer values ", intValues)
	intValue, isSuccessfulInt := intValues["Mevlana"]
	fmt.Println(intValue, isSuccessfulInt)

	/*
	also look at the example we try different map creation
	as we did in arrays
	*/

	// as a result we can try something like
	if elementName, isSuccessful := elements["Mevlana"]; isSuccessful {
		fmt.Println(elementName, isSuccessful)
		fmt.Println("this statement only runs when isSuccessful is true")
	}
	if elementName, isSuccessful := elements["He"]; isSuccessful {
		fmt.Println(elementName, isSuccessful)
		fmt.Println("this statement only runs when isSuccessful is true")
	}

	/*
	we can define maps as consist of maps
	(maybe we can also use other types but not yet discovered by me :))
	and be careful us different declaration for some elements
	like Li, C, N, O
	*/
	advancedElements := map[string]map[string]string{
		"H": map[string]string{
			"name": "Hydrogen",
			"state": "gas",
		},
		"He": map[string]string{
			"name": "Helium",
			"state": "gas",
		},
		"Li": {
			"name": "Lithium",
			"state": "solid",
		},
		"Be": map[string]string{
			"name": "Beryllium",
			"state": "solid",
		},
		"B": map[string]string{
			"name": "Boron",
			"state": "solid",
		},
		"C": {
			"name": "Carbon",
			"state": "solid",
		},
		"N": {
			"name": "Nitrogen",
			"state": "gas",
		},
		"O": {
			"name": "Oxygen",
			"state": "gas",
		},
	}
	fmt.Println(advancedElements)
	if advancedElement, isSuccessfulAdvanced := advancedElements["Li"]; isSuccessfulAdvanced {
		fmt.Println(advancedElement["name"], advancedElement["state"])
	}

}

func answerFour()  {
	x := []int{
		48,96,86,68,
		57,82,63,70,
		37,34,83,27,
		19,97, 9,17,
	}
	// we assume all values are smaller than 100
	// in general it is good to use +infinity
	minValue := 1000
	for _, value := range x{
		if value < minValue {
			minValue = value
		}
	}
	fmt.Println("Minimum value is ", minValue)

}

/*
answers:
1-)
	array[3]
2-)
	3
	explanation required
3-)
	["c", "d", "e"]
4-)
	Answer is above
*/