package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"time"

	pb "github.com/cgeorgiades27/grpc-chat/pkg/api/v1"
	"github.com/sashabaranov/go-openai"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	grpcPort    = flag.String("grpc-port", "50052", "grpc server port")
	oaiToken    = flag.String("oai-token", "", "openai token")
	chatContext []openai.ChatCompletionMessage
)

const (
	MaxTokens = 25
)

type chatserver struct {
	oaiClient *openai.Client
	pb.UnimplementedChatServer
}

func (c chatserver) SendMessage(ctx context.Context, req *pb.ChatRequest) (*pb.ChatResponse, error) {
	chatContext = append(chatContext, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: req.UserInput,
	})

	res, err := c.oaiClient.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model:    openai.GPT3Dot5Turbo,
			Messages: chatContext,
		},
	)

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	content := res.Choices[0].Message.Content
	chatContext = append(chatContext, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleAssistant,
		Content: content,
	})

	return &pb.ChatResponse{AssistantReply: content}, nil
}

func (c chatserver) StreamMessage(req *pb.ChatRequest, stream pb.Chat_StreamMessageServer) error {
	chatContext = append(chatContext, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: req.UserInput,
	})

	res, err := c.oaiClient.CreateChatCompletion(
		stream.Context(),
		openai.ChatCompletionRequest{
			Model:    openai.GPT3Dot5Turbo,
			Messages: chatContext,
		},
	)

	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	content := res.Choices[0].Message.Content
	chatContext = append(chatContext, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleAssistant,
		Content: content,
	})

	for _, c := range content {
		if err := stream.Send(
			&pb.ChatResponse{
				AssistantReply: string(c),
			},
		); err != nil {
			return status.Error(codes.Internal, err.Error())
		}
		// delay to better demonstrate stream
		time.Sleep(time.Millisecond * 20)
	}

	return nil
}

func main() {
	flag.Parse()
	oaiClient := openai.NewClient(*oaiToken)
	chatSvr := chatserver{oaiClient: oaiClient}

	s := grpc.NewServer()
	pb.RegisterChatServer(s, chatSvr)

	// Start gRPC server
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", *grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
