package main

import (
	"context"
	"log"
	"os"

	"github.com/mark3labs/mcp-go/server"
	"github.com/simenandre/fiken-mcp/internal/fiken"
	"github.com/simenandre/fiken-mcp/tools"
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
		server.WithInstructions("This is an accounting MCP server that integrates with Fiken, a Norwegian accounting system. "+
			"When working with financial data, you must ensure all numbers are correct. "+
			"Always verify amounts, quantities, and calculations either by using code to compute them or by referencing the exact values returned from the MCP tools. "+
			"Never guess or approximate financial figures. "+
			"IMPORTANT: All monetary values in the Fiken API are represented in the smallest currency unit. "+
			"For Norwegian Krone (NOK), this means values are in øre (1 NOK = 100 øre). "+
			"For example, 10000 means 100 NOK, not 10000 NOK. "+
			"Always convert monetary values by dividing by 100 when displaying amounts to the user."),
	)

	tools.RegisterAll(s, client)

	stdio := server.NewStdioServer(s)
	if err := stdio.Listen(context.Background(), os.Stdin, os.Stdout); err != nil {
		log.Fatal(err)
	}
}
