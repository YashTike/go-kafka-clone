compile:
	protoc api/v1/*.proto \
		--proto_path=. \
		--go_out=. \
		--go-grpc_out=. \
		--go_opt=paths=source_relative \
		--go-grpc_opt=paths=source_relative
test:
	go test -race ./...