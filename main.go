package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"encoding/json"
)

type Hey struct {
	Name    string
	Message string
}

func main() {
	http.HandleFunc("/", handleHello)
	fmt.Println("serving on http://0.0.0.0:7777/")
	log.Fatal(http.ListenAndServe("0.0.0.0:7777", nil))
}

func handleHello(w http.ResponseWriter, req *http.Request) {
	hey := Hey{strings.TrimPrefix(req.URL.Path, "/"), "Hey!"}
	log.Println("serving", req.URL)
	js, err := json.Marshal(hey)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
