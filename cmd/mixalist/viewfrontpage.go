package main

import (
	"fmt"
	"github.com/GeertJohan/go.rice"
	"github.com/iand/mixalist/pkg/playlist"
	"html/template"
	"net/http"
)

func viewfrontpage(w http.ResponseWriter, r *http.Request) {

	pids, err := database.GetSortedPlaylistIDs(10, 0, []string{})
	if err != nil {
		msg := fmt.Sprintf("Could not get playlists: %v", err)
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	playlists := []*playlist.Playlist{}
	for _, pid := range pids {
		pl, err := database.GetPlaylist(pid)
		if err != nil {
			continue
		}

		playlists = append(playlists, pl)

	}

	box, _ := rice.FindBox("templates")

	templateData, _ := box.String("frontpage.html")
	t, _ := template.New("frontpage.html").Parse(templateData)

	user := getUser(w, r)

	data := map[string]interface{}{
		"uid":       user.Uid,
		"username":  user.Name,
		"playlists": playlists,
	}

	w.Header().Add("Content-Type", "text/html")
	t.Execute(w, data)

}
