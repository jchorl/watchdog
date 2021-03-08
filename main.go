package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"cloud.google.com/go/datastore"
	"github.com/golang/glog"

	"github.com/jchorl/watchdog/internal/email"
	"github.com/jchorl/watchdog/pkg/server"
)

func main() {
	flag.Set("logtostderr", "true") // for glog
	flag.Parse()

	ctx := context.Background()

	emailClient := &email.Client{
		SendGridAPIKey: sendGridAPIKey,
		Domain:         os.Getenv("DOMAIN"),
		FromEmail:      fmt.Sprintf("alerts@%s", os.Getenv("DOMAIN")),
		ToEmail:        os.Getenv("ALERT_EMAIL"),
	}

	datastoreClient, err := datastore.NewClient(ctx, os.Getenv("GOOGLE_CLOUD_PROJECT"))
	if err != nil {
		glog.Fatalf("creating db client: %+v", err)
	}

	srv := server.New(os.Getenv("PORT"), emailClient, datastoreClient)

	glog.Error(srv.ListenAndServe())
}
