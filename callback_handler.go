package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

func callbackHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Entered handler")

	decoder := json.NewDecoder(r.Body)

	// Get the events out of the request structure
	request := &struct {
		Events []*Event `json:"events"`
	}{}

	err := decoder.Decode(&request)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	for _, event := range request.Events {
		log.Println(event.Type)
		handleEvent(*event)
	}
}

func registerRouteHandlers() {

	var endpoint_port = os.Getenv("PORT")
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("images"))))
	http.HandleFunc("/bot/", callbackHandler)

	log.Println("Listening on port " + endpoint_port)

	log.Fatal(http.ListenAndServe(":"+endpoint_port, nil)) //nil means the default Router Server is used

}

func main() {
	registerRouteHandlers()
}
