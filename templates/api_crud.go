package templates

const (
	// APICRUDTemplate writes a basic CRUD API with validation.
	APICRUDTemplate = `package {{ .Name }}api

import (
	"io/ioutil"
	"net/http"
)

// Create{{ .TitleName }} calls the create job with the data in the req body.
func Create{{ .TitleName }}(w http.ResponseWriter, r *http.Request){
	if err := validateCreate{{ .TitleName }}(r); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	_, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	// TODO kick off create job here.
	// response
	w.Header().Set("Content-Type", "text/plain")
	if _, err = w.Write([]byte("Hello {{ .TitleName }}")); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// validateCreate{{ .TitleName }} checks that requirements are met before
// kicking off the create job.
func validateCreate{{ .TitleName }}(r *http.Request) (err error) {
	// TODO validate the create request here.
	return err
}

// Get{{ .TitleName }} returns the resource data based on a query or uid.
func Get{{ .TitleName }}(w http.ResponseWriter, r *http.Request){
	// TODO query system for resource data.
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	if _, err = w.Write(body); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Patch{{ .TitleName }} will update properties of the resource.
func Patch{{ .TitleName }}(w http.ResponseWriter, r *http.Request){
	// TODO post patch of resource json.
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "merge-patch/json")
	if _, err = w.Write(body); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Delete{{ .TitleName }} will validate a request to get rid of a resource
// and delete it by ID.
func Delete{{ .TitleName }}(w http.ResponseWriter, r *http.Request){
	if err := validateDelete{{ .TitleName }}(r); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	// TODO kick off delete job here, parse id from req.
	// response
	w.Header().Set("Content-Type", "application/json")
	if _, err = w.Write(body); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// validateDelete{{ .TitleName }} checks that requirements are met before
// kicking off the create job.
func validateDelete{{ .TitleName }}(r *http.Request) (err error) {
	// TODO validate the create request here.
	return err
}`
	// APITitle is the name of the CRUD API file.
	APITitle = `{{ .Name }}api.go`
)
