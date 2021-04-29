proto:
	protoc -I api/v1 \
		--go_out api/v1 \
		--go_opt paths=source_relative \
		--go-grpc_out api/v1 \
		--go-grpc_opt paths=source_relative \
		api/v1/*.proto
