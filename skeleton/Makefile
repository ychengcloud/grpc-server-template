app = server
BUF_VERSION:=0.43.2

.PHONY: run
run:
	go run ./cmd/$(app) ;

.PHONY: watch
watch:
	gowatch

.PHONY: test
test: 
	go test -v ./$(app)/... -covermode=count -coverprofile=dist/cover-$(app).out ;

.PHONY: build
build: 
	go mod tidy; \
	GOOS=linux GOARCH="amd64" go build -o dist/$(app)-linux-amd64 ./cmd/$(app)/; 

.PHONY: cover
cover: test
	go tool cover -html=dist/cover-$(app).out; 

.PHONY: mock
mock:
	mockgen -package repo -destination repositories/mock/repository.go   -source repositories/repository.go

.PHONY: lint
lint:
	golint ./...
	buf lint
	buf breaking --against 'https://github.com/johanbrandhorst/grpc-gateway-boilerplate.git#branch=master'
.PHONY: proto
proto:
	# protoc -I api/proto ./api/proto/* --go_out=plugins=grpc:api/proto
	buf generate

install:
	go install \
		google.golang.org/protobuf/cmd/protoc-gen-go \
		google.golang.org/grpc/cmd/protoc-gen-go-grpc \
		github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
		github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2
	curl -sSL \
    	"https://github.com/bufbuild/buf/releases/download/v${BUF_VERSION}/buf-$(shell uname -s)-$(shell uname -m)" \
    	-o "$(shell go env GOPATH)/bin/buf" && \
  	chmod +x "$(shell go env GOPATH)/bin/buf"

all: lint cover install
