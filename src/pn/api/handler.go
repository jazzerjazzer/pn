package api

import (
	"encoding/json"
	"net/http"
	"strings"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	id := getID(r.URL.Path)
	found := composeResponse(Start(id))
	if found == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if err := json.NewEncoder(w).Encode(found); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func composeResponse(found string) *Found {
	splitted := strings.Split(found, ",")
	return &Found{
		ID:             splitted[0],
		Price:          splitted[1],
		ExpirationDate: splitted[2],
	}
}

func getID(url string) string {
	splitted := strings.Split(url, "/")
	return splitted[len(splitted)-1]
}
