package main

import (
	"fmt"
	"github.com/iand/mixalist/pkg/playlist"
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
	value, err := database.ToggleStar(user.Uid, playlist.PlaylistID(pid))
	w.Header().Add("Content-Type", "application/json")
	fmt.Fprintf(w, "{\"value\":\"%v\"}", value)

}
