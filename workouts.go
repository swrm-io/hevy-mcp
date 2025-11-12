package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/modelcontextprotocol/go-sdk/mcp"
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
	// Set defaults and limits
	page := args.Page
	if page <= 0 {
		page = 1
	}

	size := args.Size
	if size <= 0 {
		size = 5
	}
	if size > 10 {
		size = 10
	}

	workouts, nextPage, err := s.client.GetWorkouts(page, size)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to fetch workouts: %v", err)
	}

	data, err := json.MarshalIndent(workouts, "", "  ")
	if err != nil {
		return nil, nil, fmt.Errorf("failed to marshal workouts: %v", err)
	}

	hasMore := "no"
	if nextPage > 0 {
		hasMore = fmt.Sprintf("yes (next page: %d)", nextPage)
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{
				Text: fmt.Sprintf("Fetched %d workouts (page: %d, size: %d, more pages: %s):\n\n%s", len(workouts), page, size, hasMore, string(data)),
			},
		},
	}, nil, nil
}
