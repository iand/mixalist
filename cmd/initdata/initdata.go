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
		Pid: 5,
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

	&playlist.Playlist{
		Pid: 6,
		Owner: &playlist.User{
			Uid: 2,
		},
		Title: "Boy Bands",
		Tags:  []string{"pop"},
		Entries: []*playlist.Entry{
			&playlist.Entry{Title: "Steal my girl", Artist: "One Direction", Ytid: "UpsKGvPjAgw"},
			&playlist.Entry{Title: "Night Changes", Artist: "One Direction", Ytid: "syFZfO_wfMQ"},
			&playlist.Entry{Title: "Midnight Memories", Artist: "One Direction", Ytid: "bkx9kCdaaMg"},
			&playlist.Entry{Title: "Best Song Ever", Artist: "One Direction", Ytid: "o_v9MY_FMcw"},
			&playlist.Entry{Title: "Amnesia", Artist: "Five Seconds of Summer", Ytid: "DCCJCILiX3o"},
			&playlist.Entry{Title: "Good Girls", Artist: "Five Seconds of Summer", Ytid: "0FfG_5JBVBQ"},
			&playlist.Entry{Title: "Try Hard", Artist: "Five Seconds of Summer", Ytid: "y7MaaEjsRB4"},
			&playlist.Entry{Title: "Don't stop", Artist: "Five Seconds of Summer", Ytid: "MKfzMOC19Fc"},
			&playlist.Entry{Title: "She looks so perfect", Artist: "Five Seconds of Summer", Ytid: "X2BYmmTI04I"},
			&playlist.Entry{Title: "Somebody to you", Artist: "The Vamps", Ytid: "0go2nfVXFgA"},
			&playlist.Entry{Title: "Last Night", Artist: "The Vamps", Ytid: "WLyHSOXhZhY"},
			&playlist.Entry{Title: "Wild Heart", Artist: "The Vamps", Ytid: "sCDdQwVRwxM"},
		},
	},

	&playlist.Playlist{
		Pid:       7,
		ParentPid: 5,
		Owner: &playlist.User{
			Uid: 2,
		},
		Title: "Evanescence and Other",
		Tags:  []string{"rock"},
		Entries: []*playlist.Entry{
			&playlist.Entry{Title: "Going Under", Artist: "Evanescence", Ytid: "CdhqVtpR2ts"},
			&playlist.Entry{Title: "Bring Me to Life", Artist: "Evanescence", Ytid: "3YxaaGgTQYM"},
			&playlist.Entry{Title: "Imaginary", Artist: "Evanescence", Ytid: "0xVHXJ9SUuM"},
			&playlist.Entry{Title: "Taking Over Me", Artist: "Evanescence", Ytid: "BrPGuARYymo"},
			&playlist.Entry{Title: "Hello", Artist: "Evanescence", Ytid: "9MHGtlEYZBA "},
			&playlist.Entry{Title: "Hail to the King", Artist: "Avenged Sevenfold", Ytid: "DelhLppPSxY"},
			&playlist.Entry{Title: "My Last Breath", Artist: "Evanescence", Ytid: "QTYuD2kFc5c"},
			&playlist.Entry{Title: "Whisper", Artist: "Evanescence", Ytid: "7I5CWyzTBGU"},
			&playlist.Entry{Title: "Sweet Sacrifice", Artist: "Evanescence", Ytid: "XBYhQnjyrWo"},
			&playlist.Entry{Title: "The Only One", Artist: "Evanescence", Ytid: "bhZnKDS3H4s"},
			&playlist.Entry{Title: "My Heart is Broken", Artist: "Evanescence", Ytid: "f1QGnq9jUU0"},
			&playlist.Entry{Title: "The Other Side", Artist: "Evanescence", Ytid: "HiIvtRg7-Lc"},
			&playlist.Entry{Title: "What you Want", Artist: "Evanescence", Ytid: "wVWazHTunSI"},
			&playlist.Entry{Title: "Lost in Paradise", Artist: " Evanescence", Ytid: "3rnxlW5TrBs"},
			&playlist.Entry{Title: "Brompton Cocktail", Artist: "Avenged Sevenfold", Ytid: "ZPIrGXybLYU"},
			&playlist.Entry{Title: "My Immortal", Artist: "Evanescence", Ytid: "5anLPw0Efmo"},
			&playlist.Entry{Title: "Haunted", Artist: "Evanescence", Ytid: "_mauH_uQZRY"},
			&playlist.Entry{Title: "Coming Home", Artist: "Avenged Sevenfold", Ytid: "YK0iQlKXxEk"},
			&playlist.Entry{Title: "Everybody's Fool", Artist: "Evanescence", Ytid: "jhC1pI76Rqo"},
			&playlist.Entry{Title: "Tourniquet", Artist: "Evanescence", Ytid: "kwbIkzDVVFQ"},
			&playlist.Entry{Title: "Doing Time", Artist: "Avenged Sevenfold", Ytid: "ECvWXt5Hozo"},
			&playlist.Entry{Title: "Can you feel my heart?", Artist: "Bring Me The Horizon", Ytid: "QJJYpsA5tv8"},
		},
	},

	&playlist.Playlist{
		Pid: 8,
		Owner: &playlist.User{
			Uid: 4,
		},
		Title: "New Romantics",
		Tags:  []string{"eighties"},
		Entries: []*playlist.Entry{
			&playlist.Entry{Artist: "Spandau Ballet", Title: "True", Ytid: "AR8D2yqgQ1U"},
			&playlist.Entry{Artist: "Visage", Title: "Fade to Grey", Ytid: "DZiJQL9OLqI"},
			&playlist.Entry{Artist: "The Human League", Title: "Don't You Want Me", Ytid: "uPudE8nDog0"},
			&playlist.Entry{Artist: "Duran Duran", Title: "Rio", Ytid: "uPudE8nDog0"},
			&playlist.Entry{Artist: "Heaven 17", Title: "Temptation", Ytid: "e7_Rc7m_smQ"},
			&playlist.Entry{Artist: "Culture Club", Title: "Karma Chameleon", Ytid: "JmcA9LIIXWw"},
			&playlist.Entry{Artist: "Spandau Ballet", Title: "Gold", Ytid: "gSq8ZBdSxNU"},
			&playlist.Entry{Artist: "Simple Minds", Title: "Promised You a Miracle", Ytid: "tX55HEX0hb0"},
			&playlist.Entry{Artist: "Duran Duran", Title: "Girls on Film", Ytid: "5dDCcMRpUnc"},
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
	playlist.User{
		Uid:  4,
		Name: "woolly_sheep",
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
			err := database.CreatePlaylist(p)
			if err != nil {
				log.Printf("failed to create playlist %s: %v", p.Title, err)
				os.Exit(1)
			}
			err = database.CommitTx()
			if err != nil {
				log.Printf("failed to commit transaction: %v", err)
				os.Exit(1)
			}
			log.Printf("created playlist %s with pid %d", p.Title, p.Pid)
			continue
		}
		if err != nil {
			log.Printf("failed to get playlist %d: %v", p.Pid, err)
			os.Exit(1)
		}
		log.Printf("found playlist %d ok", p.Pid)
	}
}
