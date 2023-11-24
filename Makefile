.PHONY: all gen
PACKAGE_DIR=edinet

gen:
	go install github.com/tkitsunai/schematyper@latest
	schematyper --package=edinet -o $(PACKAGE_DIR)/gen_document_list.go gen/DocumentListResponse.json
