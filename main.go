package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/", handleHello)
	fmt.Println("serving on http://0.0.0.0:7777/")
	log.Fatal(http.ListenAndServe("0.0.0.0:7777", nil))
}

func handleHello(w http.ResponseWriter, req *http.Request) {
	id := strings.TrimPrefix(req.URL.Path, "/")
	log.Println("serving", req.URL)
	fmt.Fprintln(w, "Hello, " + id)
}
