
package main

import (
	"fmt"
	"math"
)

/*
two different declaration of struct
we can group fields whose types same.
type Circle struct {
	x float64
	y float64
	z float64

}
*/

type Circle struct {
	x, y, r float64

}

/*
	Struct methods
	special functions(but they are methods) for structs
*/

func (c *Circle) area() (area float64)  {
	fmt.Println("Inside struct func. Circle: ", c)
	fmt.Println("Inside struct func. *Circle: ", *c)
	area = math.Pi * c.r * c.r
	return

}

func (c *Circle) perimeter() (perimeter float64) {
	perimeter = 0
	return

}

//

type Rectangle struct {
	x1, y1, x2, y2 float64

}

func (r *Rectangle) area() (area float64) {
	length := distance(r.x1, r.y1, r.x1, r.y2)
	width := distance(r.x1, r.y1, r.x2, r.y1)
	area = length * width
	return

}

func (r *Rectangle) perimeter() (perimeter float64) {
	perimeter = 0
	return

}

/*
	Embedded types
	We can use one type in an another

*/

type Person struct {
	Name string

}

func (p *Person) Talk()  {
	fmt.Println("Hi, my name is ", p.Name)

}

type Android struct {
	/*
	also we can simply define like
	Person
	instead of
	Person Person

	and we don't call android has a person
	we say android is a person
	*/
	Person
	Model string

}

/*
	Interfaces

*/
type Shape interface {
	area() float64
	perimeter() float64

}


func main()  {
	/*
	var c Circle
	this will create a new Circle with default 0 value
	all the fields gets zero corresponding to the their types
	(0 for int, 0.0 for floats, "" for strings, nil for pointers, etc.)
	*/
	var circle Circle
	fmt.Println(circle)

	/*
	with new method we allocate memory for all fields with default value od 0
	and returns pointer to struct
	but by using with we cant define initial values
	so we try something like following
	*/
	newCircle := new(Circle)
	fmt.Println(newCircle)

	circleWithInitial := Circle{ x: 5, y: 6, r: 7 }
	fmt.Println(circleWithInitial)

	// or if we know the order of fields
	circleWithInitialWithoutNames := Circle{ 5, 6, 7 }
	fmt.Println(circleWithInitialWithoutNames)

	/*
	but still if we a pointer to struct,
	as new function did we use the following
	*/
	circleWithInitialAndPointer := &Circle{ 5, 6, 7 }
	fmt.Println("Expect to see address", circleWithInitialAndPointer)

	/*
		Struct fields
	*/
	circleEx := Circle{ 10, 20, 7 }
	fmt.Println("x of example circle ", circleEx.x)
	circleEx.x = 100
	fmt.Println("changed x of example circle ", circleEx.x)
	fmt.Println("area of exampl circle ", circleArea(circleEx))
	fmt.Println("area of exampl circle with struct method: ", circleEx.area())

	rectangleEx := Rectangle{ 5, 5, 4, 4 }
	fmt.Println("area of example rectangle: ", rectangleEx.area())

	//
	androidEx := new(Android)
	// inspect the next two lines and comments in the Android struct
	androidEx.Person.Talk()
	androidEx.Talk()
	// Android and Person has a is-a relationship: Android is a person


	circleShapeEx := Circle{ 4, 5, 6 }
	rectangleShapeEx := Rectangle{ 3, 4, 5, 6 }
	fmt.Println(totalArea(&circleShapeEx, &rectangleShapeEx))



}

func circleArea(c Circle) (area float64) {
	area = math.Pi * c.r * c.r
	return

}

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)

}

/*
	About interfaces
if we don't have an interfaces and if we want to
implement a method that calculates total area for
different shapes.

even if we use variadic functions to handle many params
we cant use two different type of variadic parameters
if we could the problem is the same with adding new shape type
to function. we can use array of shapes like circles and rectangles
but if we want to use squares we need to extend our function.
so interfaces comes to handle this.

*/

func totalArea(shapes ...Shape) (area float64) {
	/*
	totalArea only know shapes have area method ( via interface )
	totalArea does not know the fields of shapes
	*/
	for _, shape := range shapes {
		area += shape.area()

	}
	return

}
