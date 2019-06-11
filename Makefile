GOCMD=GO111MODULE=on go

.PHONY: all gen
all: gen

PACKAGE_DIR="api/edinet/api/v1"
gen:
	$(GOCMD) get -u github.com/idubinskiy/schematyper
	schematyper --package=v1 -o $(PACKAGE_DIR)/gen_document_list.go gen/DocumentListResponse.json
