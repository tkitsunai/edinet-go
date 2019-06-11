NAME=exedinet-server
GOCMD=GO111MODULE=on go

.PHONY: all test linux darwin windows clean chmod
all: schema-v1 test linux darwin

linux: build-linux
darwin: build-darwin

build-%:
	CGO_ENABLED=0 GOOS=$* GOARCH=amd64 $(GOCMD) build -tags="$* netgo" -installsuffix netgo -o dist/$(NAME)-$* ./main.go

chmod:
	chmod +x dist/*

mod:
	$(GOCMD) mod tidy

test:
	$(GOCMD) test ./...

clean: mod
	$(GOCMD) clean

PACKAGE_DIR="edinet/api/v1"
schema-v1:
	$(GOCMD) get github.com/idubinskiy/schematyper
	schematyper --package=v1 -o $(PACKAGE_DIR)/gen_document_list.go gen/DocumentListResponse.json

run:
	PORT=8080 $(GOCMD) run main.go