BINDIR := ./bin
BIN    := snowflake-service

GOBIN    := $(shell go env GOPATH)/bin
GOSRC    := $(shell find . -type f -name '*.go' -print) go.mod go.sum
PROTOSRC := $(shell find . -type f -name '*.proto' -print)
SQLSRC   := $(shell find . -type f -name '*.sql' -print)

GOGEN     := $(GOBIN)/protoc-gen-go
GOGRPC    := $(GOBIN)/protoc-gen-go-grpc
VTPROTO   := $(GOBIN)/protoc-gen-go-vtproto
GOIMPORTS := $(GOBIN)/goimports
SQLC      := $(GOBIN)/sqlc

PROTODIR := ./api
PROTOGEN := $(PROTODIR)/pb
PROTODEF := $(patsubst $(PROTODIR)/%,%,$(PROTOSRC))

SQLDIR := ./build/package/snowflake-service
SQLGEN := ./internal/db/sqlc

VTFLAGS := marshal+unmarshal+size
LDFLAGS := -w -s

COUNT ?= 1

# -----------------------------------------------------------------
#  build

.PHONY: all
all: build

.PHONY: build
build: $(BINDIR)/$(BIN)

$(BINDIR)/$(BIN): $(GOSRC)
	go build -trimpath -ldflags '$(LDFLAGS)' -o $(BINDIR)/$(BIN) ./cmd/$(BIN)

# -----------------------------------------------------------------
#  test

.PHONY: test
test:
	go test -race -v -count=$(COUNT) ./...

# -----------------------------------------------------------------
#  generate

.PHONY: generate
generate: $(VTPROTO) $(GOGEN) $(GOGRPC) $(SQLC) $(PROTODIR)/pb/.protogen $(SQLGEN)/.sqlgen

$(PROTOGEN)/.protogen: $(PROTOSRC)
	protoc --proto_path=$(PROTODIR)                                  \
		--go_out=.         --plugin protoc-gen-go=$(GOGEN)           \
		--go-grpc_out=.    --plugin protoc-gen-go-grpc=$(GOGRPC)     \
		--go-vtproto_out=. --plugin protoc-gen-go-vtproto=$(VTPROTO) \
		--go-vtproto_opt=features=$(VTFLAGS)                         \
		$(PROTODEF)
	@touch $(PROTOGEN)/.protogen

$(SQLGEN)/.sqlgen: $(SQLSRC)
	$(SQLC) -f $(SQLDIR)/sqlc.yaml generate
	@touch $(SQLGEN)/.sqlgen

# -----------------------------------------------------------------
#  dependencies

$(GOGEN):
	( cd /; go install google.golang.org/protobuf/cmd/protoc-gen-go@latest)

$(GOGRPC):
	( cd /; go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest)

$(VTPROTO):
	( cd /; go install github.com/planetscale/vtprotobuf/cmd/protoc-gen-go-vtproto@latest)

$(GOIMPORTS):
	(cd /; go install golang.org/x/tools/cmd/goimports@latest)

$(SQLC):
	(cd /; go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest)

# -----------------------------------------------------------------
#  misc

.PHONY: format
format: $(GOIMPORTS)
	GO111MODULE=on go list -f '{{.Dir}}' ./... | xargs $(GOIMPORTS) -w -local github.com/uplite/snowflake-service

.PHONY: migration
migration: $(MIGRATE)
	$(MIGRATE) create -ext sql -dir build/package/snowflake-service/migrations/ -seq -digits 4 $(NAME)

.PHONY: clean
clean:
	rm -rf $(PROTOGEN)
	rm -rf $(SQLGEN)
	rm -rf $(BINDIR)
