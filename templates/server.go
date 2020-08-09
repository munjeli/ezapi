package templates

const (
	ServerTemplate   = `package server`
	ServerTitle      = `server.go`
	ServerAPIPath    = `{{ .TargetDir }}/apis/{{ .Name }}api/server/`
	ServerNetSrvPath = `{{ .TargetDir }}/server/`
)
