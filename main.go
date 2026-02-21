package main

import (
	"context"
	"log"
	"os"

	"github.com/mark3labs/mcp-go/server"
	"github.com/simenandre/fiken-mcp-server/internal/fiken"
	"github.com/simenandre/fiken-mcp-server/tools"
)

func main() {
	apiKey := os.Getenv("FIKEN_API_KEY")
	if apiKey == "" {
		log.Fatal("FIKEN_API_KEY environment variable is required")
	}

	client := fiken.NewClient(apiKey)

	s := server.NewMCPServer(
		"fiken-mcp-server",
		"1.0.0",
		server.WithToolCapabilities(true),
	)

	tools.RegisterAll(s, client)

	stdio := server.NewStdioServer(s)
	if err := stdio.Listen(context.Background(), os.Stdin, os.Stdout); err != nil {
		log.Fatal(err)
	}
}
