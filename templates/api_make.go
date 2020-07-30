package templates

const (
	// APIMakeTemplate has the content for an API Makefile.
	APIMakeTemplate = `all: format vet build

vet:
go vet

format:
find . -name '*.go' | xargs gofmt -s -w

build:
cd apis/{{ .Name }}/server && go build -o {{ .Name }}

clean:
cd apis/{{ .Name }}/server && rm -f {{ .Name }}.go

testapi:
cd apis/{{ .Name }} && go test -v -coverprofile=cover.out && go tool cover -func=cover.out

integration:
cd apis/{{ .Name }}/server && go test -v -coverprofile=cover.out && go tool cover -func=cover.out

stats:
git ls-files | xargs wc -l
`
	// APIMakePath is where to put the Makefile once it's generated.
	APIMakePath = `{{ .TargetDir }}/Makefile`
)
