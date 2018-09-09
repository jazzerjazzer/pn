package main

import (
	"log"
	"net/http"
	"pn/api"
)

func main() {
	http.HandleFunc("/", api.Handler)
	log.Fatal(http.ListenAndServe(":1321", nil))
}
