GOHOSTOS:=$(shell go env GOHOSTOS)
GOPATH:=$(shell go env GOPATH)
VERSION=$(shell git describe --tags --always)
CURRENT_BRANCH=$(shell git rev-parse --abbrev-ref HEAD)
CURRENT_FOLDER=$(shell basename $(shell pwd))

PROTO_FILE_NAME=wiki_protos
ifneq ($(wildcard $(PROTO_FILE_NAME)),)
   THIRD_PARTY_PATH=$(PROTO_FILE_NAME)/third_party
else
	PROTO_FILE_NAME=api
	THIRD_PARTY_PATH=./third_party
endif

PROTO_DIRS=$(shell find $(PROTO_FILE_NAME) -type d -not -path "$(PROTO_FILE_NAME)/third_party/*")
PROTO_FILES=$(foreach dir,$(PROTO_DIRS),$(wildcard $(dir)/*.proto))
PROTO_RELATIVE_PATHS=$(subst $(PROTO_FILE_NAME)/,,$(PROTO_FILES))
PROTO_TARGETS=$(foreach file,$(PROTO_RELATIVE_PATHS),$(file)=$(CURRENT_FOLDER)/api/$(shell dirname $(file)))
PROTO_RESULT=$(shell echo $(PROTO_TARGETS) | sed 's/ /,/g')
PROTO_RESULT:=$(shell echo $(PROTO_RESULT) | sed 's/,/,M/g')
IMPORT_REPLACE=M$(PROTO_RESULT)


ifeq ($(GOHOSTOS), windows)
	#the `find.exe` is different from `find` in bash/shell.
	#to see https://docs.microsoft.com/en-us/windows-server/administration/windows-commands/find.
	#changed to use git-bash.exe to run find cli or other cli friendly, caused of every developer has a Git.
	#Git_Bash= $(subst cmd\,bin\bash.exe,$(dir $(shell where git)))
	Git_Bash=$(subst \,/,$(subst cmd\,bin\bash.exe,$(dir $(shell where git))))
	INTERNAL_PROTO_FILES=$(shell $(Git_Bash) -c "find internal -name *.proto")
	API_PROTO_FILES=$(shell $(Git_Bash) -c "find api -name *.proto")
else
	INTERNAL_PROTO_FILES=$(shell find internal -name *.proto)
	API_PROTO_FILES=$(shell find $(PROTO_FILE_NAME) -type f -name "*.proto" -not -path "$(PROTO_FILE_NAME)/third_party/*")
endif

.PHONY: init
# init env
init:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/go-kratos/kratos/cmd/kratos/v2@latest
	go install github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2@latest
	go install github.com/google/gnostic/cmd/protoc-gen-openapi@latest
	go install github.com/google/wire/cmd/wire@latest

.PHONY: config
# generate internal proto
config:
	protoc --proto_path=./$(PROTO_FILE_NAME) \
       	   --proto_path=./internal \
	       --proto_path=$(THIRD_PARTY_PATH) \
 	       --go_out=paths=source_relative,,$(IMPORT_REPLACE):./internal \
	       $(INTERNAL_PROTO_FILES)

.PHONY: api
# generate api proto
api:
	protoc --proto_path=./$(PROTO_FILE_NAME) \
		   --proto_path=$(THIRD_PARTY_PATH) \
 	       --go_out=paths=source_relative,$(IMPORT_REPLACE):./api \
 	       --go-http_out=paths=source_relative,$(IMPORT_REPLACE):./api \
 	       --go-grpc_out=paths=source_relative,$(IMPORT_REPLACE):./api \
	       --openapi_out=fq_schema_naming=true,default_response=false:. \
	       $(API_PROTO_FILES)

.PHONY: build
# build
build:
	mkdir -p bin/ && go build -ldflags "-X main.Version=$(VERSION)" -o ./bin/ ./...

.PHONY: proto
# add wiki_protos
proto:
	git submodule add http://git55.fxeyeinterface.com/public-projects/back-end/wiki_protos.git

.PHONY: dep
# update wiki_protos
dep:
	git submodule update --init --remote
	cd $(PROTO_FILE_NAME) && git fetch && git checkout -B $(CURRENT_BRANCH)  && git pull origin $(CURRENT_BRANCH)

.PHONY: generate
# generate
generate:
	go mod tidy
	go get github.com/google/wire/cmd/wire@latest
	go generate ./...

.PHONY: all
# generate all
all:
	make api;
	make config;
	make generate;

.PHONY: check
# check golang code
check:
	golangci-lint run -c .golangci.yml

# show help
help:
	@echo ''
	@echo 'Usage:'
	@echo ' make [target]'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
	helpMessage = match(lastLine, /^# (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")); \
			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
			printf "\033[36m%-22s\033[0m %s\n", helpCommand,helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help
