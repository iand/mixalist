package main

import (
	"fmt"
	"github.com/GeertJohan/go.rice"
	"github.com/iand/mixalist/pkg/db"
	"github.com/iand/mixalist/pkg/playlist"
	"html/template"
	"net/http"
	"strconv"
)

func remixplaylist(w http.ResponseWriter, r *http.Request) {

	pidstr := r.URL.Query().Get("pid")
	pid, err := strconv.Atoi(pidstr)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	d, err := db.Connect(true)

	if err != nil {
		msg := fmt.Sprintf("Could not connect to database: %v", err)
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	pl, err := d.GetPlaylist(playlist.PlaylistID(pid))
	if err != nil {
		msg := fmt.Sprintf("Could not get playlist: %v", err)
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	box, _ := rice.FindBox("templates")

	templateData, _ := box.String("mixplaylist.html")
	t, _ := template.New("remixplaylist.html").Parse(templateData)

	data := map[string]interface{}{

		"playlist": pl,
	}

	w.Header().Add("Content-Type", "text/html")
	t.Execute(w, data)
}
