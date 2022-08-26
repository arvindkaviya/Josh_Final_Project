package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"example.com/routers"
)

func main() {

	fmt.Println("Server Started!!")

	r := mux.NewRouter()

	//Handling Routes
	r.HandleFunc("/websites", routers.WebsitesListHandler).Methods("POST")
	r.HandleFunc("/website", routers.SingleStatusHandler).Methods("GET")
	r.HandleFunc("/websites", routers.StatusHandler).Methods("GET")

	//Listening to the server at port 3000
	http.ListenAndServe(":3000", r)
}
