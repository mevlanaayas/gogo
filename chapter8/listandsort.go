
package main

import (
	"container/list"
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age int

}

type ByName []Person
type ByAge []Person

func (ps ByName) Len() int  {
	return len(ps)

}

func (ps ByAge) Len() int  {
	return len(ps)

}

func (ps ByName) Less(i, j int) bool  {
	return ps[i].Name < ps[j].Name

}

func (ps ByAge) Less(i, j int) bool  {
	return ps[i].Age < ps[j].Age

}

func (ps ByName) Swap(i, j int)  {
	ps[i], ps[j] = ps[j], ps[i]

}

func (ps ByAge) Swap(i, j int)  {
	ps[i], ps[j] = ps[j], ps[i]

}


func main()  {
	// linked list
	var x list.List
	x.PushBack("me")
	x.PushBack("vl")
	x.PushBack("an")

	for e := x.Front(); e != nil; e = e.Next() {
		// search about this syntax
		fmt.Println(e.Value.(string))
	}

	// sort
	kids := []Person{
		{"Selami", 23},
		{"Ahmet", 22},
		{"Mevlana", 14},
	}
	fmt.Println("before sort: ", kids)
	sort.Sort(ByName(kids))
	fmt.Println("after sort by name : ", kids)
	sort.Sort(ByAge(kids))
	fmt.Println("after sort by age : ", kids)

}