NAME=edinet-go
GOCMD=GO111MODULE=on go

.PHONY: all test linux darwin windows clean chmod
all: schema-v1 test linux darwin windows

linux: build-linux
darwin: build-darwin
windows: windows-build

build-%:
	CGO_ENABLED=0 GOOS=$* GOARCH=amd64 $(GOCMD) build -tags="$* netgo" -installsuffix netgo -o bin/$(NAME)-$* ./main.go
windows-build:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 $(GOCMD) build -tags="$* netgo" -installsuffix netgo -o bin/$(NAME).exe ./main.go

chmod:
	chmod +x bin/*

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