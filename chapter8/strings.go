
package main

import (
	"fmt"
	"strings"
)

func main()  {
	// run "godoc strings -funcName- "
	fmt.Println(strings.Contains("test", "es"))

	fmt.Println(strings.Count("test", "t"))

	fmt.Println(strings.HasPrefix("test", "te"))

	fmt.Println(strings.HasSuffix("test", "st"))

	fmt.Println(strings.Index("test", "s"))

	fmt.Println(strings.Join([]string{"a", "b"}, "-"))

	fmt.Println(strings.Repeat("a", 5))

	fmt.Println(strings.Replace("aaaaaa", "a", "b", 2))

	fmt.Println(strings.Replace("aaaaaa", "a", "b", -1))

	fmt.Println(strings.Split("a-b-c-d-e", "-"))

	fmt.Println(strings.ToLower("MEvlaNa"))

	fmt.Println(strings.ToUpper("mevLana aYAs"))

	//convert string to binary
	arr := []byte("test")
	// convert binary to string
	str := string([]byte{'t', 'e', 's', 't'})
	fmt.Println("Binary: ", arr)
	fmt.Println("String: ", str)

}