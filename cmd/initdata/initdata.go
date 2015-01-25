package main

import (
	"fmt"
	"github.com/iand/mixalist/pkg/db"
	"github.com/iand/mixalist/pkg/playlist"
	"log"
	"os"
)

var playlists = []playlist.Playlist{}

func main() {
	database, err := db.Connect(true)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not connect to database: %v", err)
		os.Exit(1)
	}

	for _, pl := range playlists {
		pldb, err := database.GetPlaylist(pl.Pid)
		if err != nil {
			log.Printf("could not get playlist: %v", err)
		}
	}

}
