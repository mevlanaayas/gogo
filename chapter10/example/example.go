
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type HomePage struct {
	URL string
	Size int
}

func main()  {
	urls := []string{
		"http://www.apple.com",
		"http://www.amazon.com",
		"http://www.google.com",
		"http://www.ebay.com",
		"http://www.facebook.com",
	}
	mp := ExProgram(urls)
	fmt.Println("finished, ", mp)

}



func ExProgram(urls []string) HomePage {

	results := make(chan HomePage)

	for _, url := range urls {
		go func(url string) {
			response, err := http.Get(url)
			if err != nil {
				panic(err)
			}
			defer response.Body.Close()

			content, err := ioutil.ReadAll(response.Body)
			if err != nil {
				panic(err)
			}

			hp := HomePage{
				URL: url,
				Size: len(content),
			}

			fmt.Println("site: ", url)
			fmt.Println("size: ", len(content))
			results <- hp

		}(url)
	}

	var maxPage HomePage
	maxPage.Size = 0
	for range urls {
		result := <- results
		if result.Size > maxPage.Size {
			maxPage = result
		}
	}

	fmt.Println("biggest page, ", maxPage.URL)
	fmt.Println("with size of, ", maxPage.Size)
	return maxPage

}

/*
answers:
1-)
	channel <-  "something"
	"something" <- channel

2-)
	Answer is above

3-)
	bufferedChannel = make(chan int, 20)
*/