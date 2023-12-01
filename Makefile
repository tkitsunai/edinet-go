.PHONY: all gen setup test build run
PACKAGE_DIR=edinet

gen:
	schematyper --package=edinet --root-type=EdinetDocumentResponse -o $(PACKAGE_DIR)/gen_edinet_response.go gen/edinet_response.json

init:
	cp -af .edinet-apikey.yml.example .edinet-apikey.yml

setup:
	go install github.com/tkitsunai/schematyper@latest
	go install github.com/cosmtrek/air@latest

test:
	go test -v ./...

run:
	air
