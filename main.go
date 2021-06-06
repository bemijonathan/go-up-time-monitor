package main

import (
	"fmt"
)

type Action interface {
	checkWebsite()
	reportShortDown()
}

// a website Property
type Website struct {
	name           string
	url            string
	expectedStatus int
	interval       int
	reportToEmail  string
	nextTime       int
}

func main() {
	websites := getWebsites()
	for index, v := range websites {
		if !v.checkWebsite() {
			fmt.Println(v.name + " is down.")
		}
		setNewTime(&websites[index])
	}
	// fmt.Println(websites)
}

func setNewTime(w *Website) {
	w.nextTime = 30
}
