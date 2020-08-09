package templates

const (
	APITestCRUDTemplate = `package {{ .Name }}api`
	APIPath             = `{{ .TargetDir }}/apis/{{ .Name }}api/`
	APITestFileName     = `{{ .Name }}api_test.go`
)
