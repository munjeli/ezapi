package templates

const (
	APIServerTemplate   = `package server

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
	NetSrvServerTemplate   = `package server

import (
	"log"
	"net/http"
)

func main () {
	http.HandleFunc("/", server.{{ .TitleName }}Handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
`
	ServerTitle      = `server.go`
	ServerAPIPath    = `{{ .TargetDir }}/apis/{{ .Name }}api/server/`
	ServerNetSrvPath = `{{ .TargetDir }}/server/`

)

