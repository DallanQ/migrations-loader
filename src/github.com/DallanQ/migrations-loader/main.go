package main

import (
	"flag"
	"fmt"
	"os"
    "golang.org/x/net/context"
    "google.golang.org/cloud/datastore"
    "log"
)

const projectID = "fsmigrations"
const keyFilepath = "resources/fsmigrations-f4d1a46ae501.json"
const dbKind = "migrations"

var typ = flag.String("t", "", "(i)mmigrations or (e)migrations (required)")

type Migration struct {
	Type   string     `json:"type" datastore:",noindex"`
	Place  string     `json:"place" datastore:",noindex"`
	Year   string     `json:"year" datastore:",noindex"`
}

func main() {
	flag.Parse()
	if *typ != "i" && *typ != "e" {
		flag.Usage()
		os.Exit(1)
	}

	fmt.Println("Hello world")

    ctx := getContext(projectID, keyFilepath, datastore.ScopeDatastore, datastore.ScopeUserEmail)
    client, err := datastore.NewClient(ctx, projectID)
    if err != nil {
        log.Fatal("Error getting client", client)
    }

    key := datastore.NewKey(ctx, "testkind", "test", 0, nil)
    src := &Migration{
        Type: "e",
        Place: "place",
        Year: "1800",
    }

    retry(func() error {
   		_, err := client.Put(ctx, key, src)
   		return err
   	})

    fmt.Println("Goodbye world")
}

