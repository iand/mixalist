package main

import (
	"fmt"
	"github.com/GeertJohan/go.rice"
	"github.com/gorilla/mux"
	"github.com/iand/mixalist/pkg/playlist"
	"html/template"
	"net/http"
)

func main() {

	router := mux.NewRouter()
	router.PathPrefix("/css/").Handler(http.StripPrefix("/css/", http.FileServer(rice.MustFindBox("css").HTTPBox())))
	router.PathPrefix("/fonts/").Handler(http.StripPrefix("/fonts/", http.FileServer(rice.MustFindBox("fonts").HTTPBox())))
	router.PathPrefix("/js/").Handler(http.StripPrefix("/js/", http.FileServer(rice.MustFindBox("js").HTTPBox())))
	router.PathPrefix("/images/").Handler(http.StripPrefix("/images/", http.FileServer(rice.MustFindBox("images").HTTPBox())))

	router.Path("/hello").HandlerFunc(hello)
	router.Path("/").HandlerFunc(frontpage)

	server := &http.Server{
		Addr:    ":8000",
		Handler: router,
	}

	server.ListenAndServe()
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	fmt.Fprint(w, "YO")
	fmt.Fprint(w, "<br>HIII")
}
func frontpage(w http.ResponseWriter, r *http.Request) {
	box, _ := rice.FindBox("templates")

	templateData, _ := box.String("frontpage.html")
	t, _ := template.New("frontpage.html").Parse(templateData)

	data := map[string]interface{}{

		"playlists": []playlist.Playlist{
			playlist.Playlist{
				Title: "Top Taylor Swift",
				Owner: &playlist.User{
					Name: "agitated_weasle",
				},
				Stars: 59,
				Tags: []string{
					"love", "pop", "teen",
				},
			},
			playlist.Playlist{
				Title: "Rainy Days",
				Owner: &playlist.User{
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
