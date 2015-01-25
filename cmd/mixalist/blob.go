package main

import (
    "github.com/gorilla/mux"
    "github.com/iand/mixalist/pkg/blobstore"
    "log"
    "net/http"
    "io"
)

func blobHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    blobID := blobstore.ID(vars["blobid"])
    
    if !blobstore.DefaultStore.Exists(blobID) {
        http.NotFound(w, r)
        return
    }
    
    mimeType, err := blobstore.DefaultStore.GetType(blobID)
    if err != nil {
        log.Printf("blobHandler: failed to get type of blob %s: %s", blobID, err.Error())
        mimeType = "application/octet-stream"
    }
    
    reader, err := blobstore.DefaultStore.Open(blobID)
    if err != nil {
        log.Printf("blobHandler: failed to open blob %s: %s", blobID, err.Error())
        http.Error(w, "failed to open blob", http.StatusInternalServerError)
        return
    }
    
    w.Header().Set("Content-Type", mimeType)
    
    _, err = io.Copy(w, reader)
    reader.Close()
    if err != nil {
        log.Printf("blobHandler: failed to copy blob to response writer: %s", err.Error())
        return
    }

}
