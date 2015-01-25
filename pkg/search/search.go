package search

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/iand/mixalist/pkg/db"
	"github.com/iand/youtube"
)

const (
	SourceYouTube    = "youtube"
	SourceMixalist   = "mixalist"
	SourceSoundcloud = "soundcloud"
)

var DefaultDatabase *db.Database

type SearchFunc func(query string, results chan Result, quit chan bool, done chan bool)

var (
	searchers = []SearchFunc{
		searchYouTube,
		searchEntries,
		searchSoundcloud,
	}
)

func Search(query string, max int) []Result {
	results := []Result{}

	done := make(chan bool, len(searchers))
	quit := make(chan bool)
	res := make(chan Result)

	timer := time.NewTimer(2500 * time.Millisecond)

	remaining := len(searchers)
	for _, fn := range searchers {
		go fn(query, res, quit, done)
	}

	for {
		select {
		case <-timer.C:
			log.Printf("time up")
			close(quit)
			return results
		case r := <-res:
			results = append(results, r)
			max--
			if max == 0 {
				log.Printf("max")
				return results
			}
		case <-done:
			remaining--
			if remaining == 0 {
				log.Printf("all done")
				// all done
				return results
			}
		}
	}
}

type Result struct {
	Title      string      `json:"title"`
	Source     string      `json:"source"`
	SourceID   string      `json:"sourceid"`
	MediaURL   string      `json:"mediaurl"`
	PreviewURL string      `json:"previewurl"`
	Ext        interface{} `json:"ext,omitempty"`
}

func searchYouTube(query string, results chan Result, quit chan bool, done chan bool) {
	client := youtube.New()
	feed, err := client.VideoSearch(query)
	if err != nil {
		done <- true
		return
	}

	for _, e := range feed.Entries {
		id := e.ID.String()
		if !strings.HasPrefix(id, "tag:youtube.com,2008:video:") {
			continue
		}

		ytid := id[27:]

		result := Result{
			Title:    e.Title.String(),
			Source:   SourceYouTube,
			SourceID: ytid,
			MediaURL: fmt.Sprintf("https://www.youtube.com/watch?v=%s", ytid),
		}

		bestImageName := ""

		for _, img := range e.Media.Thumbnails {
			if img.Name == "sddefault" ||
				(img.Name == "hqdefault" && bestImageName != "sddefault") ||
				(img.Name == "mqdefault" && bestImageName != "sddefault" && bestImageName != "hqdefault") ||
				(img.Name == "default " && bestImageName != "mqdefault" && bestImageName != "sddefault" && bestImageName != "hqdefault") {
				result.PreviewURL = img.URL
				bestImageName = img.Name
			}
		}

		select {
		case results <- result:
		case <-quit:
			return
		}
	}

	done <- true
}

func searchEntries(query string, results chan Result, quit chan bool, done chan bool) {
	words := strings.Split(query, " ")
	entries, err := DefaultDatabase.SearchEntries(10, 0, words...)
	if err != nil {
		done <- true
		return
	}

	for _, e := range entries {
		result := Result{
			Title:    e.Title,
			Source:   SourceMixalist,
			SourceID: e.Ytid,
		}

		select {
		case results <- result:
		case <-quit:
			return
		}
	}
	done <- true

}

type SoundcloudTrack struct {
	ID           int    `json:"id"`
	Title        string `json:"title"`
	Duration     int    `json:"duration"` //"duration": 243892,
	ArtworkURL   string `json:"artwork_url"`
	Permalink    string `json:"permalink"`     //     "permalink": "evanescence-bring-me-to-life",
	URI          string `json:"uri"`           //     "uri": "https://api.soundcloud.com/tracks/40197833",
	PermalinkURL string `json:"permalink_url"` // "permalink_url": "http://soundcloud.com/richo-adis-saputra/evanescence-bring-me-to-life",
	StreamURL    string `json:"stream_url"`    //  "stream_url": "https://api.soundcloud.com/tracks/40197833/stream",
}

func searchSoundcloud(query string, results chan Result, quit chan bool, done chan bool) {
	var tracks []SoundcloudTrack

	url := fmt.Sprintf("https://api.soundcloud.com/tracks.json?consumer_key=0fde8f66aecb751990e0a8c0af52736f&filter=all&order=default&q=%s", url.QueryEscape(query))
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("failed to fetch from soundcloud: %v", err)
		done <- true
		return
	}

	defer resp.Body.Close()

	dec := json.NewDecoder(resp.Body)

	if err := dec.Decode(&tracks); err != nil {
		log.Printf("failed to parse json from soundcloud: %v", err)
		done <- true
		return
	}

	for _, t := range tracks {
		result := Result{
			Title:      t.Title,
			Source:     SourceSoundcloud,
			SourceID:   strconv.Itoa(t.ID),
			MediaURL:   t.StreamURL,
			PreviewURL: t.ArtworkURL,
		}

		select {
		case results <- result:
		case <-quit:
			return
		}
	}
	done <- true
}
