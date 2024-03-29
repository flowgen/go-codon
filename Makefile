installdependencies:
	./scripts/dependencies.sh
	go mod tidy

installtestdependencies:
	go get github.com/stretchr/testify/assert

igenerate:
	go generate

ibuild:
	go build -o codon codon.go

clean:
	go clean

build: clean installdependencies igenerate ibuild

iinstall:
	go install codon.go

install: clean installdependencies igenerate iinstall

itest:
	./scripts/tests.sh

test: install installtestdependencies itest

.PHONY: installdependencies clean
