package routers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"example.com/database"
)

//All Website Status Handler
func StatusHandler(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(&database.WebsitesData.StatusMap)
	if err != nil {
		fmt.Println(err)
	}

}

//Single Website Status Handler
func SingleStatusHandler(w http.ResponseWriter, r *http.Request) {

	website := r.URL.Query().Get("name")

	if database.WebsitesData.StatusMap[website] == "" {

		w.WriteHeader(http.StatusBadRequest)

		err := json.NewEncoder(w).Encode("Such Website does not exist in records")
		if err != nil {
			fmt.Println(err)
		}
		return
	}
	res := map[string]string{website: database.WebsitesData.StatusMap[website]}

	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(res)
	if err != nil {
		fmt.Println(err)
	}

}

//Post Handler for Saving all the Websites in the Memory Map
func WebsitesListHandler(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	err := json.NewDecoder(r.Body).Decode(&database.WebsitesData)
	if err != nil {
		fmt.Println(err)
	}

	database.WebsitesData.StatusMap = make(map[string]string)

	var websitestatus database.WebsiteStatus = &database.WebsitesData

	//json.NewEncoder(w).Encode(database.WebsitesData)

	go websitestatus.GetStatus()

}
