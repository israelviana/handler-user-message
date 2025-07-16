PROTO_DIR=internal/proto
OUT_DIR=gen/proto

PROTOC_GEN_GO := $(shell which protoc-gen-go)
PROTOC_GEN_GO_GRPC := $(shell which protoc-gen-go-grpc)
PROTOC_GEN_GRPC_GATEWAY := $(shell which protoc-gen-grpc-gateway)

GOOGLEAPIS_DIR = googleapis

.PHONY: proto init-submodule check-tools

init-submodule:
	@echo ">> Initializing googleapis submodule..."
	git submodule update --init --recursive

check-tools: init-submodule
	@if [ -z "$(PROTOC_GEN_GO)" ]; then echo "protoc-gen-go not found. Install with 'go install google.golang.org/protobuf/cmd/protoc-gen-go@latest'"; exit 1; fi
	@if [ -z "$(PROTOC_GEN_GO_GRPC)" ]; then echo "protoc-gen-go-grpc not found. Install with 'go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest'"; exit 1; fi
	@if [ -z "$(PROTOC_GEN_GRPC_GATEWAY)" ]; then echo "protoc-gen-grpc-gateway not found. Install with 'go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest'"; exit 1; fi
	@if [ ! -d "$(GOOGLEAPIS_DIR)/google/api" ]; then echo "googleapis submodule not initialized. Run 'make init-submodule'"; exit 1; fi

proto: check-tools
	@echo ">> Generating Go files from proto..."
	protoc \
		-I $(PROTO_DIR) \
		-I $(GOOGLEAPIS_DIR) \
		--go_out=$(OUT_DIR) \
		--go-grpc_out=$(OUT_DIR) \
		--grpc-gateway_out=$(OUT_DIR) \
		--go_opt=paths=source_relative \
		--go-grpc_opt=paths=source_relative \
		--grpc-gateway_opt=paths=source_relative \
		$(PROTO_DIR)/whatsapp.proto