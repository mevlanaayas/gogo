
package main

import (
	"net/http"
	"io"
)

func main()  {
	http.HandleFunc("/hello", hello)
	http.Handle(
		"/assets/",
		http.StripPrefix(
			"/assets/",
			http.FileServer(http.Dir("assets")),
		),
	)
	http.ListenAndServe(":9000", nil)
}

func hello(response http.ResponseWriter, request *http.Request)  {
	response.Header().Set(
		"Content-Type",
		"text/html",
	)
	io.WriteString(
		response,
		`<DOCTYPE html>
			<html>
				<head>
					<title>Hello, World</title>
				</head>
				<body>
					Hello Go server
				</body>
			</html>`,
	)

}