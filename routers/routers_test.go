package routers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"example.com/database"
	"github.com/stretchr/testify/require"
)

func TestWebsitesListHandler(t *testing.T) {

	data := database.Websites{
		WebsitesName: []string{

			"www.google.com",
			"www.swiggy.com",
			"www.fakewebsite.com",
		},
	}

	jsonStr, _ := json.Marshal(data)

	resp, err := http.Post("http://127.0.0.1:3000/websites", "application/json", bytes.NewBuffer(jsonStr))

	if err != nil {
		t.Errorf("got error on api request %s", err)
	}

	defer resp.Body.Close()

	if resp.Status != "200 OK" {
		t.Errorf("Got unexpected resonse %v", resp.Status)
	}

}

func TestStatusHandler(t *testing.T) {
	time.Sleep(2 * time.Second)
	client := http.Client{}
	resp, err := client.Get("http://127.0.0.1:3000/websites")
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	response, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println(resp.StatusCode, string(response))

	// Check the response body is what we expect.
	expected := `{"www.google.com" : "UP","www.swiggy.com" : "UP","www.fakewebsite.com" : "DOWN"}`

	require.JSONEq(t, expected, string(response))
}

func TestSingleStatusHandler(t *testing.T) {
	time.Sleep(2 * time.Second)

	tests := []struct {
		name    string
		website string
	}{
		{"website-1", "www.google.com"},
		{"website-2", "www.fakewebsite.com"},
		{"website-3", "www.swiggy.com"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := http.Client{}
			resp, err := client.Get("http://127.0.0.1:3000/website?name=" + tt.website)
			if err != nil {
				fmt.Println(err)
			}

			defer resp.Body.Close()

			response, _ := ioutil.ReadAll(resp.Body)

			// Check the response body is what we expect.
			if tt.website == "www.fakewebsite.com" {
				expectString := map[string]string{tt.website: "DOWN"}
				expected, err := json.Marshal(expectString)
				if err != nil {
					fmt.Println(err)
				}
				require.JSONEq(t, string(expected), string(response))
			} else {
				expectString := map[string]string{tt.website: "UP"}
				expected, err := json.Marshal(expectString)
				if err != nil {
					fmt.Println(err)
				}
				require.JSONEq(t, string(expected), string(response))
			}

		})
	}
}

// client := http.Client{}
// resp, err := client.Get("http://127.0.0.1:3000/website?name=www.google.com")
// if err != nil {
// 	fmt.Println(err)
// }
// fmt.Println("hello")

// defer resp.Body.Close()

// response, _ := ioutil.ReadAll(resp.Body)
// fmt.Println(resp.StatusCode, string(response))

// // Check the response body is what we expect.
// expected := `{"www.google.com":"UP"}`

//require.JSONEq(t, expected, string(response))
