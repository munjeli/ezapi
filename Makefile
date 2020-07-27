all: format vet build

vet:
	go vet

format:
	find . -name '*.go' | xargs gofmt -s

build:
	go build -o eza

clean:
	rm -f eza

test:
	go test -v -coverprofile=cover.out
	go tool cover -func=cover.out