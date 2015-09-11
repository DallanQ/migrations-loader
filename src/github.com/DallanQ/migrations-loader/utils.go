package main

import (
    "golang.org/x/net/context"
    "golang.org/x/oauth2/google"
    "golang.org/x/oauth2"
    "google.golang.org/cloud"
    "log"
    "io/ioutil"
)

func getContext(projectID, keyFilepath string, scopes ...string) context.Context {
	jsonKeyBytes, err := ioutil.ReadFile(keyFilepath)
	if err != nil {
		log.Fatal(err)
	}
	auth, err := google.JWTConfigFromJSON(jsonKeyBytes, scopes...)
    if err != nil {
   		log.Fatal(err)
   	}
	client := auth.Client(oauth2.NoContext)
	ctx := cloud.NewContext(projectID, client)
	return ctx
}

func retry(f func() error) error {
	var errs int
	for {
		err := f()
		if err == nil {
			return nil
		}
		errs++
		if errs > 3 {
			return err
		}
	}
}
