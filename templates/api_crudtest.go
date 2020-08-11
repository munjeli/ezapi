package templates

const (
	// APITestCRUDTemplate is the template for tests on a CRUD API.
	APITestCRUDTemplate = `package {{ .Name }}api`
	// APIPath refers to the location of the API handler in the repo.
	APIPath = `{{ .TargetDir }}/apis/{{ .Name }}api/`
	// APITestFileName has the tests for the CRUD API.
	APITestFileName = `{{ .Name }}api_test.go`
)
