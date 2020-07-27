all: format vet build

vet:
	go vet

format:
	find . -name '*.go' | xargs gofmt -s -w

build:
	go build -o eza

clean:
	rm -f eza

test:
	go test -v -coverprofile=cover.out
	go tool cover -func=cover.out

stats:
	git ls-files | xargs wc -l