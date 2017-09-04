package serve

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/munnerz/metaldata/pkg/authorisation"
	"github.com/munnerz/metaldata/pkg/registry"
	"github.com/munnerz/metaldata/pkg/util/errors"
)

const (
	httpRefHeader = "X-Set-Ref"
)

type Listener struct {
	auth     authorisation.Interface
	registry registry.Interface
}

func NewListener(auth authorisation.Interface, registry registry.Interface) Listener {
	return Listener{auth, registry}
}

func (l *Listener) Serve() error {
	r := mux.NewRouter()

	r.PathPrefix("/").Methods("GET").HandlerFunc(l.handleGet)
	r.PathPrefix("/").Methods("POST").HandlerFunc(l.handlePost)

	return http.ListenAndServe(":8080", r)
}

func (l *Listener) handleGet(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	ref, err := l.auth.Authorize(r)

	if err != nil {
		log.Printf("error authorising request: %s", err.Error())
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(err.Error()))
		return
	}

	val, err := l.registry.Get(ref, registry.Key(r.RequestURI))

	if errors.IsNotFound(err) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(fmt.Sprintf("key '%s' not found", err.Error())))
		return
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("error reading value for ref '%s': %s", ref, err.Error())
		return
	}

	w.Write([]byte(val))
}

func (l *Listener) handlePost(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	ref := registry.SourceRef(r.Header.Get(httpRefHeader))

	if len(ref) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("%s header must be set", httpRefHeader)))
		return
	}

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Printf("error reading POST body: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	key := registry.Key(r.RequestURI)
	val := string(body)

	err = l.registry.Set(ref, key, val)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("error writing key '%s' (value: '%s') for ref '%s': %s", key, val, ref, err.Error())
		return
	}
}
