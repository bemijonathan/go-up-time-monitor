package main

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
	channel := make(chan bool)
	length := 0
	for index, v := range websites {
		println(index)
		go v.checkWebsite(channel)
		length++
	}

	for {
		x := <-channel
		println(x)
	}
}

func setNewTime(w *Website) {
	w.nextTime = 30
}
