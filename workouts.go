package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/swrm-io/go-hevy"
)

func (s svc) getWorkoutCount(ctx context.Context, req *mcp.CallToolRequest, args NoArgs) (*mcp.CallToolResult, any, error) {
	count, err := s.client.WorkoutCount()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to fetch workout count: %v", err)
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{
				Text: fmt.Sprintf("Total workouts: %d", count),
			},
		},
	}, nil, nil
}

func (s svc) getWorkouts(ctx context.Context, req *mcp.CallToolRequest, args Fetch) (*mcp.CallToolResult, any, error) {
	var workouts []hevy.Workout
	for workout := range s.client.Workouts() {
		workouts = append(workouts, workout)
		if len(workouts) >= args.Count {
			break
		}
	}

	data, err := json.MarshalIndent(workouts, "", "  ")
	if err != nil {
		return nil, nil, fmt.Errorf("failed to marshal workouts: %v", err)
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{
				Text: fmt.Sprintf("Fetched %d workouts:\n\n%s", len(workouts), string(data)),
			},
		},
	}, nil, nil

}
