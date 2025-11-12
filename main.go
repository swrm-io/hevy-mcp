package main

import (
	"context"
	"log"
	"log/slog"
	"os"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/swrm-io/go-hevy"
)

func main() {
	apiKey, ok := os.LookupEnv("HEVY_API_KEY")
	if !ok || apiKey == "" {
		slog.Error("HEVY_API_KEY environment variable is not set")
		os.Exit(1)
	}

	slog.Info("Starting Hevy MCP Server")

	client := hevy.NewClient(apiKey)
	svc := svc{
		client: client,
	}
	server := mcp.NewServer(
		&mcp.Implementation{
			Name:    "hevy",
			Version: "v0.0.1",
		},
		nil,
	)

	mcp.AddTool(
		server,
		&mcp.Tool{
			Name:        "get_workout_count",
			Description: "Get the total count of workouts",
		},
		svc.getWorkoutCount,
	)

	mcp.AddTool(
		server,
		&mcp.Tool{
			Name:        "get_workouts",
			Description: "Get Workouts from newest to oldest",
		},
		svc.getWorkouts,
	)

	// Start server with stdio transport
	if err := server.Run(context.Background(), &mcp.StdioTransport{}); err != nil {
		log.Fatal(err)
	}
}
