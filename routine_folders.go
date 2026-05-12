package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type RoutineFolderID struct {
	ID int `json:"id" jsonschema:"The routine folder ID"`
}

func (s svc) getRoutineFolders(ctx context.Context, req *mcp.CallToolRequest, args Fetch) (*mcp.CallToolResult, any, error) {
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

	result, err := s.client.RoutineFolders.List(ctx, page, size)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to fetch routine folders: %v", err)
	}

	data, err := json.MarshalIndent(result.RoutineFolders, "", "  ")
	if err != nil {
		return nil, nil, fmt.Errorf("failed to marshal routine folders: %v", err)
	}

	hasMore := "no"
	if page < result.PageCount {
		hasMore = fmt.Sprintf("yes (next page: %d)", page+1)
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{
				Text: fmt.Sprintf("Fetched %d routine folders (page: %d, size: %d, more pages: %s):\n\n%s", len(result.RoutineFolders), page, size, hasMore, string(data)),
			},
		},
	}, nil, nil
}

func (s svc) getRoutineFolder(ctx context.Context, req *mcp.CallToolRequest, args RoutineFolderID) (*mcp.CallToolResult, any, error) {
	folder, err := s.client.RoutineFolders.Get(ctx, args.ID)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to fetch routine folder: %v", err)
	}

	data, err := json.MarshalIndent(folder, "", "  ")
	if err != nil {
		return nil, nil, fmt.Errorf("failed to marshal routine folder: %v", err)
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: string(data)},
		},
	}, nil, nil
}
