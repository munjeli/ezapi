package templates

const (
	APIServerTemplate   = `package main

import (
	"log"
	"net/http"

	"github.com/munjeli/{{ .TargetDir }}/apis/{{ .Name }}api"
)

func main () {
	http.HandleFunc("/", {{ .Name }}api.Create{{ .TitleName }})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
`
	NetSrvServerTemplate   = `package main

import (
	"log"
	"net/http"
)

func main () {
	http.HandleFunc("/", {{ .Name }}Handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
`
	ServerTitle      = `main.go`
	ServerAPIPath    = `{{ .TargetDir }}/`
	ServerNetSrvPath = `{{ .TargetDir }}/`

)

