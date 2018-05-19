package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/", handleHello)
	fmt.Println("serving on http://localhost:7777/hello")
	log.Fatal(http.ListenAndServe("localhost:7777", nil))
}

func handleHello(w http.ResponseWriter, req *http.Request) {
	id := strings.TrimPrefix(req.URL.Path, "/")
	log.Println("serving", req.URL)
	fmt.Fprintln(w, "Hello, " + id)
}
