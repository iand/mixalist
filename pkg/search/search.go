package search

import (
	"time"

	"github.com/iand/youtube"
)

const (
	SourceYouTube = "yt"
)

type SearchFunc func(query string, results chan Result, quit chan bool, done chan bool)

var (
	searchers = []SearchFunc{
		searchYouTube,
	}
)

func Search(query string) []Result {
	results := []Result{}

	done := make(chan bool, len(searchers))
	quit := make(chan bool)
	res := make(chan Result)

	timer := time.NewTimer(250 * time.Millisecond)

	remaining := len(searchers)
	for _, fn := range searchers {
		go fn(query, res, quit, done)
	}

	for {
		select {
		case <-timer.C:
			close(quit)
			return results
		case r := <-res:
			results = append(results, r)
		case <-done:
			remaining--
			if remaining == 0 {
				// all done
				return results
			}
		}
	}
}

type Result struct {
	Title    string `json:"title"`
	Source   string `json:"source"`
	SourceID string `json:"sourceid"`
}

func searchYouTube(query string, results chan Result, quit chan bool, done chan bool) {
	client := youtube.New()
	feed, err := client.VideoSearch(query)
	if err != nil {
		done <- true
	}

	for _, e := range feed.Entries {
		result := Result{
			Title:  e.Title.String(),
			Source: SourceYouTube,
		}

		select {
		case results <- result:
		case <-quit:
			return
		}
	}

	done <- true
}
