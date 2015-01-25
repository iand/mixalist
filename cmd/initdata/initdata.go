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
				Title: "Going Under", SrcName: "youtube", SrcID: "CdhqVtpR2ts", Artist: "Evanescence",
			},
			&playlist.Entry{
				Title: "Bring Me to Life", SrcName: "youtube", SrcID: "3YxaaGgTQYM", Artist: "Evanescence",
			},
			&playlist.Entry{
				Title: "Everybody's Fool", SrcName: "youtube", SrcID: "jhC1pI76Rqo", Artist: "Evanescence",
			},
			&playlist.Entry{
				Title: "My Immortal", SrcName: "youtube", SrcID: "5anLPw0Efmo", Artist: "Evanescence",
			},
			&playlist.Entry{
				Title: "Haunted", SrcName: "youtube", SrcID: "_mauH_uQZRY", Artist: "Evanescence",
			},
			&playlist.Entry{
				Title: "Tourniquet", SrcName: "youtube", SrcID: "kwbIkzDVVFQ", Artist: "Evanescence",
			},
			&playlist.Entry{
				Title: "Imaginary", SrcName: "youtube", SrcID: "0xVHXJ9SUuM", Artist: "Evanescence",
			},
			&playlist.Entry{
				Title: "Taking Over Me", SrcName: "youtube", SrcID: "BrPGuARYymo", Artist: "Evanescence",
			},
			&playlist.Entry{
				Title: "Hello", SrcName: "youtube", SrcID: "9MHGtlEYZBA", Artist: "Evanescence",
			},
			&playlist.Entry{
				Title: "My Last Breath", SrcName: "youtube", SrcID: "QTYuD2kFc5c", Artist: "Evanescence",
			},
			&playlist.Entry{
				Title: "Whisper", SrcName: "youtube", SrcID: "7I5CWyzTBGU", Artist: "Evanescence",
			},
			&playlist.Entry{
				Title: "Sweet Sacrifice", SrcName: "youtube", SrcID: "XBYhQnjyrWo", Artist: "Evanescence",
			},
			&playlist.Entry{
				Title: "The Only One", SrcName: "youtube", SrcID: "bhZnKDS3H4s", Artist: "Evanescence",
			},
			&playlist.Entry{
				Title: "My Heart is Broken", SrcName: "youtube", SrcID: "f1QGnq9jUU0", Artist: "Evanescence",
			},
			&playlist.Entry{
				Title: "The Other Side", SrcName: "youtube", SrcID: "HiIvtRg7-Lc", Artist: "Evanescence",
			},
			&playlist.Entry{
				Title: "What you Want", SrcName: "youtube", SrcID: "wVWazHTunSI", Artist: "Evanescence",
			},
			&playlist.Entry{
				Title: "Lost in Paradise", SrcName: "youtube", SrcID: "3rnxlW5TrBs", Artist: "Evanescence",
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
			&playlist.Entry{Title: "Steal my girl", Artist: "One Direction", SrcName: "youtube", SrcID: "UpsKGvPjAgw"},
			&playlist.Entry{Title: "Night Changes", Artist: "One Direction", SrcName: "youtube", SrcID: "syFZfO_wfMQ"},
			&playlist.Entry{Title: "Midnight Memories", Artist: "One Direction", SrcName: "youtube", SrcID: "bkx9kCdaaMg"},
			&playlist.Entry{Title: "Best Song Ever", Artist: "One Direction", SrcName: "youtube", SrcID: "o_v9MY_FMcw"},
			&playlist.Entry{Title: "Amnesia", Artist: "Five Seconds of Summer", SrcName: "youtube", SrcID: "DCCJCILiX3o"},
			&playlist.Entry{Title: "Good Girls", Artist: "Five Seconds of Summer", SrcName: "youtube", SrcID: "0FfG_5JBVBQ"},
			&playlist.Entry{Title: "Try Hard", Artist: "Five Seconds of Summer", SrcName: "youtube", SrcID: "y7MaaEjsRB4"},
			&playlist.Entry{Title: "Don't stop", Artist: "Five Seconds of Summer", SrcName: "youtube", SrcID: "MKfzMOC19Fc"},
			&playlist.Entry{Title: "She looks so perfect", Artist: "Five Seconds of Summer", SrcName: "youtube", SrcID: "X2BYmmTI04I"},
			&playlist.Entry{Title: "Somebody to you", Artist: "The Vamps", SrcName: "youtube", SrcID: "0go2nfVXFgA"},
			&playlist.Entry{Title: "Last Night", Artist: "The Vamps", SrcName: "youtube", SrcID: "WLyHSOXhZhY"},
			&playlist.Entry{Title: "Wild Heart", Artist: "The Vamps", SrcName: "youtube", SrcID: "sCDdQwVRwxM"},
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
			&playlist.Entry{Title: "Going Under", Artist: "Evanescence", SrcName: "youtube", SrcID: "CdhqVtpR2ts"},
			&playlist.Entry{Title: "Bring Me to Life", Artist: "Evanescence", SrcName: "youtube", SrcID: "3YxaaGgTQYM"},
			&playlist.Entry{Title: "Imaginary", Artist: "Evanescence", SrcName: "youtube", SrcID: "0xVHXJ9SUuM"},
			&playlist.Entry{Title: "Taking Over Me", Artist: "Evanescence", SrcName: "youtube", SrcID: "BrPGuARYymo"},
			&playlist.Entry{Title: "Hello", Artist: "Evanescence", SrcName: "youtube", SrcID: "9MHGtlEYZBA "},
			&playlist.Entry{Title: "Hail to the King", Artist: "Avenged Sevenfold", SrcName: "youtube", SrcID: "DelhLppPSxY"},
			&playlist.Entry{Title: "My Last Breath", Artist: "Evanescence", SrcName: "youtube", SrcID: "QTYuD2kFc5c"},
			&playlist.Entry{Title: "Whisper", Artist: "Evanescence", SrcName: "youtube", SrcID: "7I5CWyzTBGU"},
			&playlist.Entry{Title: "Sweet Sacrifice", Artist: "Evanescence", SrcName: "youtube", SrcID: "XBYhQnjyrWo"},
			&playlist.Entry{Title: "The Only One", Artist: "Evanescence", SrcName: "youtube", SrcID: "bhZnKDS3H4s"},
			&playlist.Entry{Title: "My Heart is Broken", Artist: "Evanescence", SrcName: "youtube", SrcID: "f1QGnq9jUU0"},
			&playlist.Entry{Title: "The Other Side", Artist: "Evanescence", SrcName: "youtube", SrcID: "HiIvtRg7-Lc"},
			&playlist.Entry{Title: "What you Want", Artist: "Evanescence", SrcName: "youtube", SrcID: "wVWazHTunSI"},
			&playlist.Entry{Title: "Lost in Paradise", Artist: " Evanescence", SrcName: "youtube", SrcID: "3rnxlW5TrBs"},
			&playlist.Entry{Title: "Brompton Cocktail", Artist: "Avenged Sevenfold", SrcName: "youtube", SrcID: "ZPIrGXybLYU"},
			&playlist.Entry{Title: "My Immortal", Artist: "Evanescence", SrcName: "youtube", SrcID: "5anLPw0Efmo"},
			&playlist.Entry{Title: "Haunted", Artist: "Evanescence", SrcName: "youtube", SrcID: "_mauH_uQZRY"},
			&playlist.Entry{Title: "Coming Home", Artist: "Avenged Sevenfold", SrcName: "youtube", SrcID: "YK0iQlKXxEk"},
			&playlist.Entry{Title: "Everybody's Fool", Artist: "Evanescence", SrcName: "youtube", SrcID: "jhC1pI76Rqo"},
			&playlist.Entry{Title: "Tourniquet", Artist: "Evanescence", SrcName: "youtube", SrcID: "kwbIkzDVVFQ"},
			&playlist.Entry{Title: "Doing Time", Artist: "Avenged Sevenfold", SrcName: "youtube", SrcID: "ECvWXt5Hozo"},
			&playlist.Entry{Title: "Can you feel my heart?", Artist: "Bring Me The Horizon", SrcName: "youtube", SrcID: "QJJYpsA5tv8"},
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
			&playlist.Entry{Artist: "Spandau Ballet", Title: "True", SrcName: "youtube", SrcID: "AR8D2yqgQ1U"},
			&playlist.Entry{Artist: "Visage", Title: "Fade to Grey", SrcName: "youtube", SrcID: "DZiJQL9OLqI"},
			&playlist.Entry{Artist: "The Human League", Title: "Don't You Want Me", SrcName: "youtube", SrcID: "uPudE8nDog0"},
			&playlist.Entry{Artist: "Duran Duran", Title: "Rio", SrcName: "youtube", SrcID: "uPudE8nDog0"},
			&playlist.Entry{Artist: "Heaven 17", Title: "Temptation", SrcName: "youtube", SrcID: "e7_Rc7m_smQ"},
			&playlist.Entry{Artist: "Culture Club", Title: "Karma Chameleon", SrcName: "youtube", SrcID: "JmcA9LIIXWw"},
			&playlist.Entry{Artist: "Spandau Ballet", Title: "Gold", SrcName: "youtube", SrcID: "gSq8ZBdSxNU"},
			&playlist.Entry{Artist: "Simple Minds", Title: "Promised You a Miracle", SrcName: "youtube", SrcID: "tX55HEX0hb0"},
			&playlist.Entry{Artist: "Duran Duran", Title: "Girls on Film", SrcName: "youtube", SrcID: "5dDCcMRpUnc"},
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
