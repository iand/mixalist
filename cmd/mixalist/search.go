package main

import (
	"encoding/json"
	"net/http"

	"github.com/iand/mixalist/pkg/playlist"
	"github.com/iand/mixalist/pkg/search"
	"github.com/iand/youtube"
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

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(200)

	client := youtube.New()
	feed, err := client.VideoSearch(query)
	if err != nil {
		writeJSON(w, resp)
		return
	}

	entries := make([]playlist.Entry, len(feed.Entries))
	for i, e := range feed.Entries {
		entries[i] = playlist.Entry{
			Title: e.Title.String(),
			Ytid:  e.ID.String(),
		}
	}

	_ = entries
	writeJSON(w, feed.Entries)
}

func writeJSON(w http.ResponseWriter, v interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(v)
}
