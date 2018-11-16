
package main

import "fmt"

func main()  {
	fmt.Println("8")

}

/*
answers:
1-)
	DRY

2-)
	other packages are able to see functions which starts w
th capital
	if starts with lowercase other packages are not able t see it

3-)
	if we import two packages with same name
	like memo from google/memo and mevlana/memo
	we use an alias to represent them
	like:
		import mem "google/memo"
		import mev "mevlana/memo"
	now we can use two packages without conflict

4-)
	nope

5-)
	put the comment line at top of its declaration
	example:
		// Writes all is well to screen
		func AllIsWell()

*/