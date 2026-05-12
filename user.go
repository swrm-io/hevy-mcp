package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func (s svc) getUserInfo(ctx context.Context, req *mcp.CallToolRequest, args NoArgs) (*mcp.CallToolResult, any, error) {
	info, err := s.client.User.Info(ctx)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to fetch user info: %v", err)
	}

	data, err := json.MarshalIndent(info, "", "  ")
	if err != nil {
		return nil, nil, fmt.Errorf("failed to marshal user info: %v", err)
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: string(data)},
		},
	}, nil, nil
}
