package templates

const (
	// APIServerTemplate is the template for a server for
	// the CRUD API.
	APIServerTemplate = `package main

import (
	"log"
	"net/http"

	"github.com/munjeli/{{ .TargetDir }}/apis/{{ .Name }}api"
)

func main () {
	http.HandleFunc("/create", {{ .Name }}api.Create{{ .TitleName }})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
`
	// NetSrvServerTemplate is the template for a server for a
	// networked service.
	NetSrvServerTemplate = `package main

import (
	"log"
	"net/http"
)

func main () {
	http.HandleFunc("/", {{ .Name }}Handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
`
	// ServerTitle is the name of the server file for either netsrv or API.
	ServerTitle = `main.go`
	// ServerAPIPath is the path to the main.go in a generated API.
	ServerAPIPath = `{{ .TargetDir }}/apis/{{ .Name }}api/server/`
	// ServerNetSrvPath is the path to the main.go in a generated netsrv.
	ServerNetSrvPath = `{{ .TargetDir }}/server/`
)
