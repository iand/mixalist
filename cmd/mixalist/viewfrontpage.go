package main

import (
	"github.com/GeertJohan/go.rice"
	"github.com/iand/mixalist/pkg/playlist"
	"html/template"
	"net/http"
)

func viewfrontpage(w http.ResponseWriter, r *http.Request) {
	box, _ := rice.FindBox("templates")

	templateData, _ := box.String("frontpage.html")
	t, _ := template.New("frontpage.html").Parse(templateData)

	data := map[string]interface{}{

		"playlists": []playlist.Playlist{
			playlist.Playlist{
				Title: "Top Taylor Swift",
				Owner: playlist.User{
					Name: "agitated_weasle",
				},
				Stars: 59,
				Tags: []string{
					"love", "pop", "teen",
				},
			},
			playlist.Playlist{
				Title: "Rainy Days",
				Owner: playlist.User{
					Name: "lonely_meerkat",
				},
				Stars: 107,
				Tags: []string{
					"love", "accoustic",
				},
			},
		},
	}

	w.Header().Add("Content-Type", "text/html")
	t.Execute(w, data)

}
