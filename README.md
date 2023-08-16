# GPTerminal - grpc-chat

This repository contains a simple command-line chat application built using gRPC. The application allows multiple clients to connect to a server and send messages to each other in real-time.

## How to use

1. Clone the repository to your local machine.
2. Create an openai account and generate an [API key](https://platform.openai.com/account/api-keys)
3. Install the required dependencies by running `go mod download`.
4. Start the server by running `go run cmd/server/main.go --oai-token=OPENAI_API_KEY`.
5. In a separate terminal window, start a client by running `go run cmd/client/main.go`.
6. Follow the prompts to ask questions.

## Dependencies

- Go 1.21
- protoc (if compiling protos)

## Contributing

Contributions are welcome! Please open an issue or pull request for any bug fixes or feature requests.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
