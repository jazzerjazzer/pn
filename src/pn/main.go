package main

import (
	"log"
	"net/http"
	"pn/api"
)

func main() {
	api := api.API{}
	http.HandleFunc("/promotions/", api.PromotionsHandler)
	http.HandleFunc("/upload", api.UploadHandler)
	log.Fatal(http.ListenAndServe(":1321", nil))
}
