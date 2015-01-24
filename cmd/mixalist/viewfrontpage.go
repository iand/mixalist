package main

import (
	"fmt"
	"github.com/GeertJohan/go.rice"
	"github.com/iand/mixalist/pkg/db"
	"github.com/iand/mixalist/pkg/names"
	"github.com/iand/mixalist/pkg/playlist"
	"html/template"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

func viewfrontpage(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().UnixNano())

	d, err := db.Connect(true)

	if err != nil {
		msg := fmt.Sprintf("Could not connect to database: %v", err)
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	pids, err := d.GetSortedPlaylistIDs(10, 0, []string{})
	if err != nil {
		msg := fmt.Sprintf("Could not get playlists: %v", err)
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	playlists := []*playlist.Playlist{}
	for _, pid := range pids {
		pl, err := d.GetPlaylist(pid)
		if err != nil {
			continue
		}

		playlists = append(playlists, pl)

	}

	box, _ := rice.FindBox("templates")

	templateData, _ := box.String("frontpage.html")
	t, _ := template.New("frontpage.html").Parse(templateData)

	// Fake user id - temporary
	uid := strconv.FormatInt(rand.Int63n(256*256)) + "0000000"
	username := names.NewName(uid)
	data := map[string]interface{}{
		"uid":       uid,
		"username":  username,
		"playlists": playlists,
	}

	w.Header().Add("Content-Type", "text/html")
	t.Execute(w, data)

}
