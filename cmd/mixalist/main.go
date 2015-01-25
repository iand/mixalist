package main

import (
	"fmt"
	"github.com/GeertJohan/go.rice"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/iand/mixalist/pkg/db"
	"github.com/iand/mixalist/pkg/names"
	"github.com/iand/mixalist/pkg/playlist"
	"github.com/iand/mixalist/pkg/search"
	"log"
	"math/rand"
	"net/http"
	"os"
	"runtime"
	"time"
)

var database *db.Database
var sessionStore = sessions.NewCookieStore([]byte("something-very-secret"))

var anonymousUser = &playlist.User{
	Uid:  0,
	Name: "secret_stinkbug",
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	rand.Seed(time.Now().UnixNano())
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

func getUser(w http.ResponseWriter, r *http.Request) *playlist.User {
	session, _ := sessionStore.Get(r, "mixalist")
	uidval, exists := session.Values["uid"]
	if !exists {
		return newUser(session, w, r)
	}

	uid, ok := uidval.(int)
	if !ok {
		return newUser(session, w, r)
	}
	usernameval, _ := session.Values["username"]
	username, _ := usernameval.(string)
	return &playlist.User{
		Uid:  playlist.UserID(uid),
		Name: username,
	}
}

func newUser(session *sessions.Session, w http.ResponseWriter, r *http.Request) *playlist.User {
	username := names.NewName()
	// create a new user

	err := database.BeginTx()
	if err != nil {
		log.Printf("failed to begin transaction: %v", err)
		return anonymousUser
	}

	uid, err := database.CreateUser(username)
	if err != nil {
		log.Printf("failed to create user %s: %v", username, err)
		return anonymousUser
	}
	err = database.CommitTx()
	if err != nil {
		log.Printf("failed to commit transaction: %v", err)
		return anonymousUser
	}

	log.Printf("created new user %d/%s: %v", int(uid), username, err)
	session.Values["uid"] = int(uid)
	session.Values["username"] = username
	session.Save(r, w)

	return &playlist.User{
		Uid:  uid,
		Name: username,
	}
}
