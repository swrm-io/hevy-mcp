package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type RoutineID struct {
	ID string `json:"id" jsonschema:"The routine ID"`
}

func (s svc) getRoutines(ctx context.Context, req *mcp.CallToolRequest, args Fetch) (*mcp.CallToolResult, any, error) {
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

	result, err := s.client.Routines.List(ctx, page, size)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to fetch routines: %v", err)
	}

	data, err := json.MarshalIndent(result.Routines, "", "  ")
	if err != nil {
		return nil, nil, fmt.Errorf("failed to marshal routines: %v", err)
	}

	hasMore := "no"
	if page < result.PageCount {
		hasMore = fmt.Sprintf("yes (next page: %d)", page+1)
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{
				Text: fmt.Sprintf("Fetched %d routines (page: %d, size: %d, more pages: %s):\n\n%s", len(result.Routines), page, size, hasMore, string(data)),
			},
		},
	}, nil, nil
}

func (s svc) getRoutine(ctx context.Context, req *mcp.CallToolRequest, args RoutineID) (*mcp.CallToolResult, any, error) {
	routine, err := s.client.Routines.Get(ctx, args.ID)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to fetch routine: %v", err)
	}

	data, err := json.MarshalIndent(routine, "", "  ")
	if err != nil {
		return nil, nil, fmt.Errorf("failed to marshal routine: %v", err)
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: string(data)},
		},
	}, nil, nil
}
