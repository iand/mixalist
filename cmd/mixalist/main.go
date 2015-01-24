package main

import (
	"github.com/GeertJohan/go.rice"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {

	router := mux.NewRouter()
	router.PathPrefix("/css/").Handler(http.StripPrefix("/css/", http.FileServer(rice.MustFindBox("css").HTTPBox())))
	router.PathPrefix("/fonts/").Handler(http.StripPrefix("/fonts/", http.FileServer(rice.MustFindBox("fonts").HTTPBox())))
	router.PathPrefix("/js/").Handler(http.StripPrefix("/js/", http.FileServer(rice.MustFindBox("js").HTTPBox())))
	router.PathPrefix("/images/").Handler(http.StripPrefix("/images/", http.FileServer(rice.MustFindBox("images").HTTPBox())))

	router.Path("/p").HandlerFunc(viewplaylist)

	router.Path("/").HandlerFunc(viewfrontpage)

	server := &http.Server{
		Addr:    ":8000",
		Handler: router,
	}

	server.ListenAndServe()
}
