package main

import (
	"fmt"
	"github.com/GeertJohan/go.rice"
	"github.com/gorilla/mux"
	"github.com/iand/mixalist/pkg/db"
	"github.com/iand/mixalist/pkg/search"
	"net/http"
	"os"
	"runtime"
)

var database *db.Database

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	var err error
	database, err = db.Connect(true)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not connect to database: %v", err)
		os.Exit(1)
	}
	search.DefaultDatabase = database

	router := mux.NewRouter()
	router.PathPrefix("/css/").Handler(http.StripPrefix("/css/", http.FileServer(rice.MustFindBox("css").HTTPBox())))
	router.PathPrefix("/fonts/").Handler(http.StripPrefix("/fonts/", http.FileServer(rice.MustFindBox("fonts").HTTPBox())))
	router.PathPrefix("/js/").Handler(http.StripPrefix("/js/", http.FileServer(rice.MustFindBox("js").HTTPBox())))
	router.PathPrefix("/images/").Handler(http.StripPrefix("/images/", http.FileServer(rice.MustFindBox("images").HTTPBox())))

	router.Path("/s").HandlerFunc(searchHandler)
	router.Path("/p/{pid:[0-9]+}").HandlerFunc(viewplaylist)
	router.Path("/r").HandlerFunc(remixplaylist)
	router.Path("/about").HandlerFunc(viewabout)
	router.Path("/").HandlerFunc(viewfrontpage)

	server := &http.Server{
		Addr:    ":8000",
		Handler: router,
	}

	server.ListenAndServe()
}
