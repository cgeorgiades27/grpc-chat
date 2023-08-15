
protos:
	@protoc \
		--go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		./pkg/api/v1/*.proto

run-server:
	@godotenv -f .env go run cmd/server/main.go --oai-token=$$OAI_TOKEN --grpc-port=50052

run-client:
	@go run cmd/client/main.go --grpc-server-port=50052