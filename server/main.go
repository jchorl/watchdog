package main

import (
	"io/ioutil"
	"net/http"
	"time"

	"github.com/golang/protobuf/proto"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"

	"github.com/jchorl/watchdog/types"
)

func main() {
	http.HandleFunc("/ping", pingHandler)
	appengine.Main()
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	if r.Method == "POST" {
		watch := types.Watch{}
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
			http.Error(w, "Unable to store in database", http.StatusBadRequest)
			return
		}

		w.Write([]byte("ok!"))
	}
}
