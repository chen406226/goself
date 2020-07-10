package main

import (
	"context"
	"fmt"
	"os"

	"chatgo/pkg/protocol/grpc"
	"chatgo/pkg/service/v1"
)

func main() {
	if err := grpc.RunServer(context.Background(), v1.NewChatServiceServer(), "3000"); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
