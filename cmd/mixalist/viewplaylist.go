package main

import (
	"github.com/GeertJohan/go.rice"
	"github.com/iand/mixalist/pkg/playlist"
	"html/template"
	"net/http"
)

func viewplaylist(w http.ResponseWriter, r *http.Request) {
	box, _ := rice.FindBox("templates")

	templateData, _ := box.String("playlist.html")
	t, _ := template.New("playlist.html").Parse(templateData)

	data := map[string]interface{}{

		"playlist": playlist.Playlist{
			Title: "Top Taylor Swift",
			Owner: playlist.User{
				Name: "agitated_weasle",
			},
			Stars: 59,
			Tags: []string{
				"love", "pop", "teen",
			},
			Entries: []playlist.Entry{
				playlist.Entry{
					Title:  "Track1",
					Artist: "Artist1",
				},
				playlist.Entry{
					Title:  "Track2",
					Artist: "Artist2",
				},
				playlist.Entry{
					Title:  "Track3",
					Artist: "Artist3",
				},
			},
		},
	}

	w.Header().Add("Content-Type", "text/html")
	t.Execute(w, data)
}
