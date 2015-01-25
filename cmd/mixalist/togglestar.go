package main

import (
	"fmt"
	"github.com/iand/mixalist/pkg/playlist"
	"log"
	"net/http"
	"strconv"
)

func togglestar(w http.ResponseWriter, r *http.Request) {

	pidstr := r.FormValue("pid")

	if pidstr == "" {
		http.NotFound(w, r)
		return
	}

	pid, err := strconv.Atoi(pidstr)

	if err != nil {
		http.NotFound(w, r)
		return
	}

	user := getUser(w, r)

	err = database.BeginTx()
	if err != nil {
		log.Printf("failed to begin transaction: %v", err)
		http.NotFound(w, r)
		return
	}

	value, err := database.ToggleStar(user.Uid, playlist.PlaylistID(pid))

	err = database.CommitTx()
	if err != nil {
		log.Printf("failed to commit transaction: %v", err)
		http.NotFound(w, r)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	fmt.Fprintf(w, "{\"value\":\"%v\"}", value)

}
