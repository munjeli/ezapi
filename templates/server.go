package templates

const (
	ServerTemplate   = `This is the server template for {{ .Name }}`
	ServerTitle      = `server.go`
	ServerAPIPath    = `{{ .TargetDir }}/apis/{{ .Name }}api/server/`
	ServerNetSrvPath = `{{ .TargetDir }}/server/`
)
