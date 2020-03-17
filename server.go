package watchdog

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"cloud.google.com/go/datastore"
	"github.com/golang/glog"
	"github.com/golang/protobuf/proto"
	"google.golang.org/appengine"
	"google.golang.org/appengine/mail"
)

func main() {
	http.HandleFunc("/ping", pingHandler)
	http.HandleFunc("/check", checkHandler)
	http.HandleFunc("/remove", removeHandler)
	glog.Error(http.ListenAndServe(os.Getenv("PORT"), nil))
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	dbClient := newDBClient(r.Context())

	watch := Watch{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		glog.Errorf("reading body: %s", err)
		http.Error(w, "reading body", http.StatusBadRequest)
		return
	}

	err = proto.Unmarshal(body, &watch)
	if err != nil {
		glog.Errorf("decoding body: %s", err)
		http.Error(w, "decoding body", http.StatusBadRequest)
		return
	}

	key := datastore.NameKey("watch", watch.Name, nil)
	watch.LastSeen = time.Now().Unix()

	if _, err := dbClient.Put(ctx, key, &watch); err != nil {
		glog.Errorf(ctx, "storing in db: %s", err)
		http.Error(w, "storing in database", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func checkHandler(w http.ResponseWriter, r *http.Request) {
	dbClient := newDBClient(r.Context())

	q := dbClient.NewQuery("watch")
	for t := q.Run(ctx); ; {
		var watch Watch
		_, err := t.Next(&watch)
		if err == datastore.Done {
			break
		}
		if err != nil {
			sendErrorEmail(ctx, err)
			glog.Errorf("querying watches from database: %s", err)
			http.Error(w, "querying watches from database", http.StatusInternalServerError)
			return
		}

		if (watch.Frequency == Watch_DAILY && time.Unix(watch.LastSeen, 0).Add(time.Hour*25).Before(time.Now())) || (watch.Frequency == Watch_WEEKLY && time.Unix(watch.LastSeen, 0).Add(time.Hour*24*7).Add(time.Hour).Before(time.Now())) {
			sendServiceDownEmail(ctx, watch)
		}
	}
}

func removeHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		http.Error(w, "name query param missing", http.StatusBadRequest)
	}

	dbClient := newDBClient(r.Context())

	key := datastore.NameKey("watch", name, nil)
	err := dbClient.Delete(ctx, key)
	if err != nil {
		glog.Errorf(ctx, "deleting watch from database: %s", err)
		http.Error(w, "deleting watch from database", http.StatusInternalServerError)
	}
}

// TODO figure out email
func sendErrorEmail(ctx context.Context, err error) {
	msg := &mail.Message{
		Sender:  fmt.Sprintf("Watchdog Notifications <notifications@%s.appspotmail.com>", appengine.AppID(ctx)),
		To:      []string{Email},
		Subject: "Watchdog is down",
		Body:    fmt.Sprintf("Watchdog is down. Error: %s", err),
	}
	if err := mail.Send(ctx, msg); err != nil {
		glog.Errorf(ctx, "Couldn't send email: %s", err)
	}
}

func sendServiceDownEmail(ctx context.Context, watch Watch) {
	msg := &mail.Message{
		Sender:  fmt.Sprintf("Watchdog Notifications <notifications@%s.appspotmail.com>", appengine.AppID(ctx)),
		To:      []string{Email},
		Subject: fmt.Sprintf("%s is down", watch.Name),
		Body:    fmt.Sprintf("%s is down and was last seen %+v. The frequency is set to %s.", watch.Name, watch.LastSeen, watch.Frequency.String()),
	}
	if err := mail.Send(ctx, msg); err != nil {
		glog.Errorf(ctx, "Couldn't send email: %s", err)
	}
}

func newDBClient(ctx context.Context) *datastore.Client {
	dbClient, err := datastore.NewClient(ctx, config.ProjectID)
	if err != nil {
		glog.Fatalf("creating db client: %s", err)
	}

	return dbClient
}
