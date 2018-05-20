package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"encoding/json"
)

type Hey struct {
	Name    string `json:"name"`
	Message string `json:"message"`
}

func main() {
	http.HandleFunc("/", handleHello)
	fmt.Println("serving on http://0.0.0.0:7777/")
	log.Fatal(http.ListenAndServe("0.0.0.0:7777", nil))
}

func handleHello(w http.ResponseWriter, req *http.Request) {
	username := strings.TrimPrefix(req.URL.Path, "/")
	message := Hey{username, "Hey!"}
	log.Println("serving", req.URL, username)
	js, err := json.Marshal(message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
