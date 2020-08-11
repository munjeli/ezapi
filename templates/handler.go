package templates

const (
	HandlerTemplate = `package server

import (
	"io/ioutil"
	"net/http"
)

func {{ .Name }}Handler(w http.ResponseWriter, r *http.Request) {
    if err := validate(r); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	_, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	// TODO Call service here
	w.Header().Set("Content-Type", "text/plain")
	if _, err = w.Write([]byte("{{ .TitleName }} is Running!")); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func validate(r *http.Request) (err error) {
	return err
}
`
	HandlerTitle    = `handler.go`
)
