package templates

const NetsrvMake = `all: format vet build

vet:
go vet

format:
find . -name '*.go' | xargs gofmt -s -w

build:
cd server && go build -o {{ .Name }}

clean:
cd server && rm -f {{ .Name }}.go

testapi:
cd service/{{ .Name }} && go test -v -coverprofile=cover.out && go tool cover -func=cover.out

integration:
cd server && go test -v -coverprofile=cover.out && go tool cover -func=cover.out

stats:
git ls-files | xargs wc -l
`