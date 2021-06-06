package main

import (
	"net/http"
)

/**
* check the website availability of its available and equal to the expected status code
**/
func (w Website) checkWebsite() bool {
	response, err := http.Get(w.url)
	// fmt.Println(response)
	if err != nil {
		return false
	}
	if response.StatusCode != w.expectedStatus {
		return false
	}
	return true
}

// func (w Website) reportShortDown() bool {

// }
