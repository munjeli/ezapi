package templates

const (
	// APITestCRUDTemplate is the template for tests on a CRUD API.
	APITestCRUDTemplate = `package {{ .Name }}api

import (
	"net/http"
	"testing"
)

func TestCreate{{ .TitleName }}(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func TestDelete{{ .TitleName }}(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func TestGet{{ .TitleName }}(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func TestPatch{{ .TitleName }}(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func Test_validateCreate{{ .TitleName }}(t *testing.T) {
	type args struct {
		r *http.Request
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validateCreate{{ .TitleName }}(tt.args.r); (err != nil) != tt.wantErr {
				t.Errorf("validateCreate{{ .TitleName }}() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_validateDelete{{ .TitleName }}(t *testing.T) {
	type args struct {
		r *http.Request
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validateDelete{{ .TitleName }}(tt.args.r); (err != nil) != tt.wantErr {
				t.Errorf("validateDelete{{ .TitleName }}() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}`
	// APIPath refers to the location of the API handler in the repo.
	APIPath = `{{ .TargetDir }}/apis/{{ .Name }}api/`
	// APITestFileName has the tests for the CRUD API.
	APITestFileName = `{{ .Name }}api_test.go`
)
