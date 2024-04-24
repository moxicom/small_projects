
PHONY: generate_structs
generate_structs:
	protoc -I=. --go_out=. --go_opt=paths=source_relative \
			./consignment/consignment.proto
		

PHONY: generate_rpc
generate_rpc:
	protoc -I=. --go_out=. --go_opt=paths=source_relative \
			--go-grpc_out=. --go-grpc_opt=paths=source_relative \
			./consignment/consignment.proto

PHONY: run_server
run_server:
	go run cmd/grpc_server/main.go


PHONY: run_client
run_client:
	go run cmd/grpc_client/main.go