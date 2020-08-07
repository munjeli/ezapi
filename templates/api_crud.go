package templates

const APICRUDTemplate = `package {{ .Name }}api

import (
	"io/ioutil"
	"net/http"
)

// Create{{ .Name }} calls the create job with the data in the req body.
func Create{{ .Name }}(w http.ResponseWriter, r *http.Request){
	if err := validateCreate{{ .Name }}(r); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	// TODO kick off create job here.
	// response
	w.Header().Set("Content-Type", "application/json")
	if _, err = w.Write(body); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// validateCreate{{ .Name }} checks that requirements are met before
// kicking off the create job.
func validateCreate{{ .Name }}(r *http.Request) (err error) {
	// TODO validate the create request here.
	return err
}

// Get{{ .Name }} returns the resource data based on a query or uid.
func Get{{ .Name }}(w http.ResponseWriter, r *http.Request){
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

// Patch{{ .Name }} will update properties of the resource.
func Patch{{ .Name }}(w http.ResponseWriter, r *http.Request){
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

// Delete{{ .Name }} will validate a request to get rid of a resource
// and delete it by ID.
func Delete{{ .Name }}(w http.ResponseWriter, r *http.Request){
	if err := validateDelete{{ .Name }}(r); err != nil {
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

// validateDelete{{ .Name }} checks that requirements are met before
// kicking off the create job.
func validateDelete{{ .Name }}(r *http.Request) (err error) {
	// TODO validate the create request here.
	return err
}`
