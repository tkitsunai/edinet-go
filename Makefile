.PHONY: all gen setup test build run
PACKAGE_DIR=edinet

gen:
	schematyper --package=edinet -o $(PACKAGE_DIR)/gen_document_list.go gen/DocumentListResponse.json

setup:
	go install github.com/tkitsunai/schematyper@latest
	go install github.com/cosmtrek/air@latest

test:
	go test -v ./...

run:
	air
