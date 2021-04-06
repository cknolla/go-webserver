package main

import (
	"github.com/cknolla/go-webserver/pkg/handlers"
	"log"
	"net/http"
)

const port = ":8899"


func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	log.Println("Starting webserver on port", port)
	_ = http.ListenAndServe(port, nil)
}
