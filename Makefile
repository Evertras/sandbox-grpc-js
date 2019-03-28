BINARY_NAME=sandbox-grpc-js
TS_FILES=$(shell find front/src -name "*.ts")
WASM_FILES=$(shell find lib -name "*.go" ! -path "lib/static/*" ! -path "lib/server/*" ! -name "*_test.go")
GO_PROTO_BUILD_DIR=lib/messages

all: test build

clean:
	rm -f $(BINARY_NAME)
	rm -f front/game.js
	rm -f front/game.js.map
	rm -rf $(GO_PROTO_BUILD_DIR)
	rm -rf front/src/proto

test: node_modules protos
	npx tslint -p .
	npm test
	go test -v ./lib/...

build: protos front/game.js
	CG_ENABLED=0 go build -o $(BINARY_NAME) -v ./cmd/server/main.go

bench:
	go test -v -benchmem -bench . ./lib/...

run-dev:
	go run -race ./cmd/server/main.go -d

docker: clean build
	docker build --rm -t evertras/$(BINARY_NAME) .

protos: $(GO_PROTO_BUILD_DIR) front/src/proto

# These are not files, so always run them when asked to
.PHONY: all clean test build bench run-dev protos docker

# Actual files/directories that must be generated
front/game.js: node_modules front/src/proto $(TS_FILES)
	npx webpack || (rm -f front/game.js && exit 1)

node_modules:
	npm install

$(GO_PROTO_BUILD_DIR): messages/*.proto
	rm -rf $(GO_PROTO_BUILD_DIR)
	mkdir -p $(GO_PROTO_BUILD_DIR)
	protoc --twirp_out=./lib --go_out=./lib ./messages/*.proto || (rmdir $(GO_PROTO_BUILD_DIR) && exit 1)

front/src/proto: node_modules messages/*.proto
	rm -rf front/src/proto
	mkdir front/src/proto
	npx pbjs -t static-module -w commonjs messages/echo.proto -o front/src/proto/echo.pb.js || (rm -rf front/src/proto && exit 1)
	npx pbts -o front/src/proto/echo.pb.d.ts front/src/proto/echo.pb.js || (rm -rf front/src/proto && exit 1)
	protoc --twirp_typescript_out=library=pbjs:./front/src/proto messages/echo.proto
