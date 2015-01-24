package main

import (
	"encoding/json"
	"net/http"

	"github.com/iand/mixalist/pkg/search"
)

type SearchResponse struct {
	Results []search.Result `json:"results"`
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	resp := SearchResponse{}

	query := r.URL.Query().Get("q")
	if query == "" {
		writeJSON(w, resp)
		return
	}

	resp.Results = search.Search(query, 8)
	writeJSON(w, resp)
}

func writeJSON(w http.ResponseWriter, v interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(v)
}
