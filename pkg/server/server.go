package server

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"cloud.google.com/go/datastore"
	"github.com/golang/glog"
	"github.com/golang/protobuf/proto"
	"google.golang.org/api/iterator"

	"github.com/jchorl/watchdog/internal/email"
	pb "github.com/jchorl/watchdog/proto"
)

type Server struct {
	*http.Server
	email *email.Client
	db    *datastore.Client
}

func New(
	port string,
	emailClient *email.Client,
	datastoreClient *datastore.Client,
) *Server {
	srv := &Server{
		email: emailClient,
		db:    datastoreClient,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/ping", srv.handlePing)
	mux.HandleFunc("/check", srv.handleCheck)
	mux.HandleFunc("/remove", srv.handleRemove)

	srv.Server = &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: mux,
	}

	return srv
}

func (s *Server) handlePing(w http.ResponseWriter, r *http.Request) {
	watch := serializedWatch{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		glog.Errorf("reading body: %s", err)
		http.Error(w, "reading body", http.StatusBadRequest)
		return
	}

	err = proto.Unmarshal(body, watch.Proto)
	if err != nil {
		glog.Errorf("decoding body: %s", err)
		http.Error(w, "decoding body", http.StatusBadRequest)
		return
	}

	key := datastore.NameKey("watch", watch.Proto.GetName(), nil)
	watch.Proto.LastSeen = time.Now().Unix()

	if _, err := s.db.Put(r.Context(), key, &watch); err != nil {
		glog.Errorf("storing in db: %s", err)
		http.Error(w, "storing in database", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (s *Server) handleCheck(w http.ResponseWriter, r *http.Request) {
	q := datastore.NewQuery("watch")
	it := s.db.Run(r.Context(), q)
	var watch serializedWatch
	for {
		_, err := it.Next(&watch)
		if err == iterator.Done {
			break
		}
		if err != nil {
			s.email.SendErrorEmail(r.Context(), err)
			glog.Errorf("querying watches from database: %s", err)
			http.Error(w, "querying watches from database", http.StatusInternalServerError)
			return
		}

		if (watch.Proto.GetFrequency() == pb.Watch_DAILY &&
			time.Unix(watch.Proto.GetLastSeen(), 0).Add(time.Hour*25).Before(time.Now())) ||
			(watch.Proto.GetFrequency() == pb.Watch_WEEKLY &&
				time.Unix(watch.Proto.GetLastSeen(), 0).Add(time.Hour*24*7).Add(time.Hour).Before(time.Now())) {
			s.email.SendServiceDownEmail(r.Context(), watch.Proto)
		}
	}
}

func (s *Server) handleRemove(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		http.Error(w, "name query param missing", http.StatusBadRequest)
	}

	key := datastore.NameKey("watch", name, nil)
	err := s.db.Delete(r.Context(), key)
	if err != nil {
		glog.Errorf("deleting watch from database: %s", err)
		http.Error(w, "deleting watch from database", http.StatusInternalServerError)
	}
}
