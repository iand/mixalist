package main

import (
	"encoding/json"
	"fmt"
	"github.com/GeertJohan/go.rice"
	"github.com/iand/mixalist/pkg/playlist"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"
)

func remixplaylist(w http.ResponseWriter, r *http.Request) {

	pidstr := r.URL.Query().Get("pid")
	pid, err := strconv.Atoi(pidstr)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	pl, err := database.GetPlaylist(playlist.PlaylistID(pid))
	if err != nil {
		msg := fmt.Sprintf("Could not get playlist: %v", err)
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	box, _ := rice.FindBox("templates")

	templateData, _ := box.String("remixplaylist.html")
	t, _ := template.New("remixplaylist.html").Parse(templateData)

	user := getUser(w, r)
	data := map[string]interface{}{
		"uid":      user.Uid,
		"username": user.Name,

		"playlist": pl,
	}

	w.Header().Add("Content-Type", "text/html")
	t.Execute(w, data)
}

func remixApiHandler(w http.ResponseWriter, r *http.Request) {
	var reqData remixApiRequestData
	err := json.NewDecoder(r.Body).Decode(&reqData)
	if err != nil {
		msg := fmt.Sprintf("Could not decode request payload: %s", err.Error())
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	user := getUser(w, r)

	entries := make([]*playlist.Entry, len(reqData.Playlist.Entries))

	for i, e := range reqData.Playlist.Entries {
		entries[i] = &playlist.Entry{
			Title:    e.Title,
			Artist:   e.Artist,
			Album:    e.Album,
			Duration: time.Duration(e.Duration) * time.Second,
			SrcName:  e.SrcName,
			SrcID:    e.SrcID,
		}
	}

	p := &playlist.Playlist{
		Title:     reqData.Playlist.Title,
		Owner:     user,
		Tags:      reqData.Playlist.Tags,
		Entries:   entries,
		ParentPid: reqData.Playlist.ParentPid,
	}

	err = database.BeginTx()
	if err != nil {
		log.Printf("remixApiHandler: database error: %s", err.Error())
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	err = database.CreatePlaylist(p)
	if err != nil {
		log.Printf("remixApiHandler: database error: %s", err.Error())
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	err = database.CommitTx()
	if err != nil {
		log.Printf("remixApiHandler: database error: %s", err.Error())
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	respData := remixApiResponseData{
		Pid: p.Pid,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(respData)
	if err != nil {
		log.Printf("remixApiHandler: failed to write response JSON: %s", err.Error())
		return
	}
}

type remixApiRequestData struct {
	Playlist remixApiPlaylist `json:"playlist"`
}

type remixApiPlaylist struct {
	Title     string              `json:"title"`
	Tags      []string            `json:"tags"`
	Entries   []*remixApiEntry    `json:"entries"`
	ParentPid playlist.PlaylistID `json:"parentPid"`
}

type remixApiEntry struct {
	Title    string `json:"title"`
	Artist   string `json:"artist"`
	Album    string `json:"album"`
	Duration int    `json:"duration"`
	SrcName  string `json:"srcName"`
	SrcID    string `json:"srcID"`
	ImageURL string `json:"imageURL"`
}

type remixApiResponseData struct {
	Pid playlist.PlaylistID `json:"pid"`
}
