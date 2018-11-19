
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"testing"
)

/*

create http server and serve two different text content
long.com serves "this is a long string and should be bigger than short.com"
short.com serves "short string"

test example for that long.com is bigger than short.com

*/

type testpair struct {
	values []string
	expected string
}

var BiggestPageTestCase = []testpair{
	{[]string{"http://127.0.0.1:9999/short", "http://127.0.0.1:9999/long"}, "http://127.0.0.1:9999/long"},
}

func TestExProgram(t *testing.T) {
	fmt.Println("initiated tests")
	go func() {
		http.HandleFunc("/long", LongCom)

	}()
	go func() {
		http.HandleFunc("/short", ShortCom)

	}()
	go func() {
		http.ListenAndServe(":9999", nil)
	}()
	fmt.Println("servers are ready")

	fmt.Println("starting to test the cases")
	for _, pair := range BiggestPageTestCase {
		result := ExProgram(pair.values)
		if result.URL != pair.expected {
			panic("Error")
		}
	}
	os.Exit(0)

}



func LongCom(response http.ResponseWriter, request *http.Request)  {
	response.Header().Set(
		"Content-Type",
		"text/html",
	)
	io.WriteString(
		response,
		`
		<DOCTYPE html>
		<html>
   			<head>
				<title>Long Com</title>
			</head>
			<body>
				this is a long string and should be bigger than short.com
			</body>
		</html>
		`,
	)
}

func ShortCom(response http.ResponseWriter, request *http.Request)  {
	response.Header().Set(
		"Content-Type",
		"text/html",
	)
	io.WriteString(
		response,
		`
		<DOCTYPE html>
		<html>
   			<head>
				<title>Short Com</title>
			</head>
			<body>
				short string
			</body>
		</html>
		`,
	)
}