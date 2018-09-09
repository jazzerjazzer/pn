package api

import (
	"encoding/json"
	"net/http"
	"path"
	"strings"
)

func (a *API) PromotionsHandler(w http.ResponseWriter, r *http.Request) {
	id := getID(r.URL.Path)
	found := composeResponse(Search(id, path.Join(a.currentFilepath, a.currentFilename)))
	w.Header().Set("Content-Type", "application/json")

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
	if len(splitted) < 2 {
		return nil
	}
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
