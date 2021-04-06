package main

import (
	"log"
	"net/http"
)

const port = ":8899"


func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	log.Println("Starting webserver on port", port)
	_ = http.ListenAndServe(port, nil)
}
