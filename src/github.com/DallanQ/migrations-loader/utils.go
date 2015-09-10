package main
import (
    "golang.org/x/net/context"
    "io/ioutil"
    "log"
    "golang.org/x/oauth2/google"
    "golang.org/x/oauth2"
    "google.golang.org/cloud"
)

func getContext(projectID, jsonKeyFilepath string, scopes ...string) context.Context {
	jsonKeyBytes, err := ioutil.ReadFile(jsonKeyFilepath)
	if err != nil {
		log.Fatal(err)
	}
	conf, err := google.JWTConfigFromJSON(jsonKeyBytes, scopes...)
	if err != nil {
		log.Fatal(err)
	}
	client := conf.Client(oauth2.NoContext)
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
