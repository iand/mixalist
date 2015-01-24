package db

import (
    "fmt"
    "os"
)

const driver = "postgres"

var source string

func init() {
    source = os.Getenv("MIXALIST_DB_SOURCE")
    if source == "" {
        fmt.Fprintf(os.Stderr, "Environment variable MIXALIST_DB_SOURCE is not set.\n")
        fmt.Fprintf(os.Stderr, "Please set it to a value in the form:\n")
        fmt.Fprintf(os.Stderr, "  postgres://USER:PASSWD@HOST/DBNAME\n")
        os.Exit(2)
    }
}
