package blobstore

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

const IDLength = 16

const idChars = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"

type ID string

type Store string

func NewStore(dir string) Store {
	return Store(dir)
}

func (store Store) mkID() (id ID) {
	rand.Seed(time.Now().UnixNano())
	buf := make([]byte, IDLength)
	for i := range buf {
		buf[i] = idChars[rand.Intn(len(idChars))]
	}
	return ID(buf)
}

func (store Store) getFilename(id ID) (filename string) {
	return filepath.Join(string(store), string(id[:2]), string(id[2:4]), string(id[4:6]), string(id)+".dat")
}

func (store Store) mkUnusedID() (id ID) {
	id = store.mkID()
	for fileExists(store.getFilename(id)) {
		id = store.mkID()
	}

	return id
}

func (store Store) Create() (id ID, w io.WriteCloser, err error) {
	id = store.mkUnusedID()
	blobFile := store.getFilename(id)
	blobDir := filepath.Dir(blobFile)
	err = os.MkdirAll(blobDir, 0755)
	if err != nil {
		return "", nil, err
	}

	f, err := os.Create(blobFile)
	return id, f, err
}

func (store Store) Open(id ID) (r io.ReadCloser, err error) {
	blobFile := store.getFilename(id)
	f, err := os.Open(blobFile)
	return f, err
}

func (store Store) Delete(id ID) (err error) {
	blobFile := store.getFilename(id)
	return os.Remove(blobFile)
}

func (store Store) Download(url string) (id ID, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	if resp.StatusCode != 200 {
		return "", fmt.Errorf("HTTP request returned %s", resp.Status)
	}
	
	id, w, err := store.Create()
	if err != nil {
		resp.Body.Close()
		return "", err
	}
	
	_, err = io.Copy(w, resp.Body)
	w.Close()
	resp.Body.Close()
	if err != nil {
		store.Delete(id)
		return "", err
	}
	
	return id, nil
}

func fileExists(filename string) (exists bool) {
	_, err := os.Stat(filename)
	return err == nil
}

var DefaultStore Store

func init() {
	DefaultStore = Store(os.Getenv("MIXALIST_BS_DIR"))
	if DefaultStore == "" {
		fmt.Fprintf(os.Stderr, "Environment variable MIXALIST_BS_DIR is not set.\n")
		fmt.Fprintf(os.Stderr, "Please set it to the blobstore directory.\n")
		os.Exit(2)
	}
}
