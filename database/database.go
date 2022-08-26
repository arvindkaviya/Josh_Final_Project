package database

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

//Interface which has Method to check status of a Website
type WebsiteStatus interface {
	GetStatus()
}

//Struct for storing Websites with their status
type Websites struct {
	WebsitesName []string `json:"websites"`
	StatusMap    map[string]string
}

var WebsitesData Websites

//Function for checking the status of a Website with Go routine
func (websites *Websites) GetStatus() {
	for {
		var wg sync.WaitGroup
		for _, val := range websites.WebsitesName {
			wg.Add(1)
			go checkStatus(val, &wg)
		}

		time.Sleep(60 * time.Second)
		wg.Wait()

	}

}

//Helper Function for checking the status
func checkStatus(website string, wg *sync.WaitGroup) {
	defer wg.Done()
	res, err := http.Get("https://" + website)
	if err != nil {
		WebsitesData.StatusMap[website] = "DOWN"
	} else if res.StatusCode == 200 {
		WebsitesData.StatusMap[website] = "UP"
	}

	fmt.Println(website, WebsitesData.StatusMap[website])

}
