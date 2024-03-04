package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/muzzapp/devops-interview-task/pkg/muzz"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, dialErr := grpc.Dial("127.0.0.1:9876", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if dialErr != nil {
		slog.Error("Failed to dial gRPC service: %v", "err", dialErr)
		os.Exit(1)
	}

	client := muzz.NewServiceClient(conn)

	response, respErr := client.Echo(
		context.Background(),
		&muzz.EchoRequest{Message: "Hello, world!"},
	)
	if respErr != nil {
		slog.Error("Failed to call gRPC service: %v", "err", respErr)
		os.Exit(1)
	}

	slog.Info("Response: %s", "message", response.Message)
}
