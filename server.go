package watchdog

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/golang/protobuf/proto"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/mail"
)

func main() {
	http.HandleFunc("/ping", pingHandler)
	http.HandleFunc("/check", checkHandler)
	http.HandleFunc("/remove", removeHandler)
	appengine.Main()
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	if r.Method == "POST" {
		watch := Watch{}
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Errorf(ctx, "Unable to read body: %s", err)
			http.Error(w, "Unable to read body", http.StatusBadRequest)
			return
		}

		err = proto.Unmarshal(body, &watch)
		if err != nil {
			log.Errorf(ctx, "Unable to decode body: %s", err)
			http.Error(w, "Unable to decode body", http.StatusBadRequest)
			return
		}

		key := datastore.NewKey(ctx, "watch", watch.Name, 0, nil)
		watch.LastSeen = time.Now().Unix()

		if _, err := datastore.Put(ctx, key, &watch); err != nil {
			log.Errorf(ctx, "Unable to store in database: %s", err)
			http.Error(w, "Unable to store in database", http.StatusInternalServerError)
			return
		}

		w.Write([]byte("ok!"))
	}
}

func checkHandler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	q := datastore.NewQuery("watch")
	for t := q.Run(ctx); ; {
		var watch Watch
		_, err := t.Next(&watch)
		if err == datastore.Done {
			break
		}
		if err != nil {
			sendErrorEmail(ctx, err)
			log.Errorf(ctx, "Unable to query watches from database: %s", err)
			http.Error(w, "Unable to query watches from database", http.StatusInternalServerError)
			return
		}

		if (watch.Frequency == Watch_DAILY && time.Unix(watch.LastSeen, 0).Add(time.Hour*25).Before(time.Now())) || (watch.Frequency == Watch_WEEKLY && time.Unix(watch.LastSeen, 0).Add(time.Hour*24*7).Add(time.Hour).Before(time.Now())) {
			sendServiceDownEmail(ctx, watch)
		}
	}
}

func removeHandler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	names := r.URL.Query()["name"]
	if len(names) == 0 {
		http.Error(w, "Send the name as a query parameter", http.StatusBadRequest)
	}

	for _, name := range names {
		key := datastore.NewKey(ctx, "watch", name, 0, nil)
		err := datastore.Delete(ctx, key)
		if err != nil {
			log.Errorf(ctx, "Unable to delete watch from database: %s", err)
			http.Error(w, "Unable to delete watch from database", http.StatusInternalServerError)
		}
		return
	}
}

func sendErrorEmail(ctx context.Context, err error) {
	msg := &mail.Message{
		Sender:  fmt.Sprintf("Watchdog Notifications <notifications@%s.appspotmail.com>", appengine.AppID(ctx)),
		To:      []string{Email},
		Subject: "Watchdog is down",
		Body:    fmt.Sprintf("Watchdog is down. Error: %s", err),
	}
	if err := mail.Send(ctx, msg); err != nil {
		log.Errorf(ctx, "Couldn't send email: %s", err)
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
		log.Errorf(ctx, "Couldn't send email: %s", err)
	}
}
