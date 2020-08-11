package templates

const (
	// APIMakeTemplate has the content for an API Makefile.
	APIMakeTemplate = `all: format vet build

vet:
	go vet

format:
	find . -name '*.go' | xargs gofmt -s -w

build:
	go build -o {{ .Name }}

clean:
	rm -f {{ .Name }}

testapi:
	cd apis/{{ .Name }}api/ && go test -v -coverprofile=cover.out && go tool cover -func=cover.out

integration:
	cd integration && go test -v -coverprofile=cover.out && go tool cover -func=cover.out

stats:
	git ls-files | xargs wc -l
`
)
