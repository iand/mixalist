package main

import (
	"fmt"
	"github.com/iand/mixalist/pkg/db"
	"github.com/iand/mixalist/pkg/playlist"
	"log"
	"os"
)

var playlists = []*playlist.Playlist{

	&playlist.Playlist{
		Pid: 1,
		Owner: &playlist.User{
			Uid: 2,
		},
		Title: "Evanescence",
		Tags:  []string{"rock"},
		Entries: []*playlist.Entry{
			&playlist.Entry{
				Title: "Going Under", Ytid: "CdhqVtpR2ts", Artist: "Evanescence",
			},
			&playlist.Entry{
				Title: "Bring Me to Life", Ytid: "3YxaaGgTQYM", Artist: "Evanescence",
			},
			&playlist.Entry{
				Title: "Everybody's Fool", Ytid: "jhC1pI76Rqo", Artist: "Evanescence",
			},
			&playlist.Entry{
				Title: "My Immortal", Ytid: "5anLPw0Efmo", Artist: "Evanescence",
			},
			&playlist.Entry{
				Title: "Haunted", Ytid: "_mauH_uQZRY", Artist: "Evanescence",
			},
			&playlist.Entry{
				Title: "Tourniquet", Ytid: "kwbIkzDVVFQ", Artist: "Evanescence",
			},
			&playlist.Entry{
				Title: "Imaginary", Ytid: "0xVHXJ9SUuM", Artist: "Evanescence",
			},
			&playlist.Entry{
				Title: "Taking Over Me", Ytid: "BrPGuARYymo", Artist: "Evanescence",
			},
			&playlist.Entry{
				Title: "Hello", Ytid: "9MHGtlEYZBA", Artist: "Evanescence",
			},
			&playlist.Entry{
				Title: "My Last Breath", Ytid: "QTYuD2kFc5c", Artist: "Evanescence",
			},
			&playlist.Entry{
				Title: "Whisper", Ytid: "7I5CWyzTBGU", Artist: "Evanescence",
			},
			&playlist.Entry{
				Title: "Sweet Sacrifice", Ytid: "XBYhQnjyrWo", Artist: "Evanescence",
			},
			&playlist.Entry{
				Title: "The Only One", Ytid: "bhZnKDS3H4s", Artist: "Evanescence",
			},
			&playlist.Entry{
				Title: "My Heart is Broken", Ytid: "f1QGnq9jUU0", Artist: "Evanescence",
			},
			&playlist.Entry{
				Title: "The Other Side", Ytid: "HiIvtRg7-Lc", Artist: "Evanescence",
			},
			&playlist.Entry{
				Title: "What you Want", Ytid: "wVWazHTunSI", Artist: "Evanescence",
			},
			&playlist.Entry{
				Title: "Lost in Paradise", Ytid: "3rnxlW5TrBs", Artist: "Evanescence",
			},
		},
	},
}

var users = []playlist.User{
	playlist.User{
		Uid:  1,
		Name: "rauyran",
	},
	playlist.User{
		Uid:  2,
		Name: "archinot",
	},
	playlist.User{
		Uid:  3,
		Name: "kierdavis",
	},
}

func main() {
	database, err := db.Connect(true)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not connect to database: %v", err)
		os.Exit(1)
	}

	for _, u := range users {
		_, err = database.GetUser(u.Uid)
		if err == db.InvalidUserError {
			log.Printf("did not find user %d, creating...", u.Uid)
			err = database.BeginTx()
			if err != nil {
				log.Printf("failed to begin transaction: %v", err)
				os.Exit(1)
			}
			uid, err := database.CreateUser(u.Name)
			if err != nil {
				log.Printf("failed to create user %s: %v", u.Name, err)
				os.Exit(1)
			}
			err = database.CommitTx()
			if err != nil {
				log.Printf("failed to commit transaction: %v", err)
				os.Exit(1)
			}
			log.Printf("created user %s with uid %d", u.Name, uid)
			continue
		}
		if err != nil {
			log.Printf("failed to get playlist %d: %v", u.Uid, err)
			os.Exit(1)
		}
		log.Printf("found user %d ok", u.Uid)
	}

	for _, p := range playlists {
		_, err = database.GetPlaylist(p.Pid)
		if err == db.InvalidPlaylistError {
			log.Printf("did not find playlist %d, creating...", p.Pid)
			err = database.BeginTx()
			if err != nil {
				log.Printf("failed to begin transaction: %v", err)
				os.Exit(1)
			}
			pid, err := database.CreatePlaylist(p)
			if err != nil {
				log.Printf("failed to create playlist %s: %v", p.Title, err)
				os.Exit(1)
			}
			err = database.CommitTx()
			if err != nil {
				log.Printf("failed to commit transaction: %v", err)
				os.Exit(1)
			}
			log.Printf("created playlist %s with pid %d", p.Title, pid)
			continue
		}
		if err != nil {
			log.Printf("failed to get playlist %d: %v", p.Pid, err)
			os.Exit(1)
		}
		log.Printf("found playlist %d ok", p.Pid)
	}
}
