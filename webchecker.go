package main

import (
	"net/http"
)

/**
* check the website availability of its available and equal to the expected status code
**/
func (w Website) checkWebsite(c chan bool) {
	response, err := http.Get(w.url)
	// fmt.Println(response)
	if err != nil {
		c <- false
	} else if response.StatusCode != w.expectedStatus {
		c <- false
	} else {
		c <- true
	}
}

// func (w Website) reportShortDown() bool {

// }
