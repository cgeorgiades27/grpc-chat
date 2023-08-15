package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"os"
	"strings"

	pb "github.com/cgeorgiades27/grpc-chat/pkg/api/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	grpcSvrIp   = flag.String("grpc-server-host", "localhost", "destination host ip")
	grpcSvrPort = flag.String("grpc-server-port", "50052", "destination host port")
)

func main() {
	flag.Parse()
	conn, err := grpc.Dial(
		*grpcSvrIp+":"+*grpcSvrPort,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	client := pb.NewChatClient(conn)
	reader := bufio.NewReader(os.Stdin)

	fmt.Println(" ---------- Welcome to ChatGPT CLI (model: GPT3Dot5Turbo) ---------- ")
	fmt.Println("\ttype your question or 'exit' to quit")
	for {
		fmt.Print("~> ")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		if text == "exit" {
			break
		}

		err := getResponseUnary(ctx, text, client)
		// err := getResponseStream(ctx, text, client)
		if err != nil {
			fmt.Printf("error getting reply: %v", err)
			continue
		}
	}
}

func getResponseUnary(ctx context.Context, input string, client pb.ChatClient) error {
	res, err := client.SendMessage(ctx, &pb.ChatRequest{UserInput: input})
	if err != nil {
		return err
	}
	fmt.Printf("gpt> %s\n", res.AssistantReply)

	return nil
}

func getResponseStream(ctx context.Context, input string, client pb.ChatClient) error {
	resStream, err := client.StreamMessage(ctx, &pb.ChatRequest{UserInput: input})
	if err != nil {
		return err
	}

	fmt.Printf("gpt> ")
	for {
		c, err := resStream.Recv()
		if err != nil {
			break
		}
		fmt.Print(c.AssistantReply)
	}
	fmt.Printf("\n")

	return nil
}
