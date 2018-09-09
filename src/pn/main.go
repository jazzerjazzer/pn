package main

import (
	"log"
	"net/http"
	"pn/api"
)

func main() {
	up := api.API{}
	http.HandleFunc("/promotions/", up.PromotionsHandler)
	http.HandleFunc("/upload", up.UploadHandler)
	log.Fatal(http.ListenAndServe(":1321", nil))
}
