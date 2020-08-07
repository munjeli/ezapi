package templates

const
(
	APITestCRUDTemplate = `This is the template for CRUD test {{ .Name }}`
	APIPath             = `{{ .TargetDir }}/apis/{{ .Name }}/`
	APITestFileName     = `{{ .Name }}api_test.go`
)
