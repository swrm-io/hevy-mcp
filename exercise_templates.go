package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type ExerciseTemplateID struct {
	ID string `json:"id" jsonschema:"The exercise template ID"`
}

type FetchExerciseTemplates struct {
	Page int `json:"page" jsonschema:"Page number to fetch (default: 1)"`
	Size int `json:"size" jsonschema:"Number of exercise templates per page (default: 20, max: 100)"`
}

func (s svc) getExerciseTemplates(ctx context.Context, req *mcp.CallToolRequest, args FetchExerciseTemplates) (*mcp.CallToolResult, any, error) {
	page := args.Page
	if page <= 0 {
		page = 1
	}
	size := args.Size
	if size <= 0 {
		size = 20
	}
	if size > 100 {
		size = 100
	}

	result, err := s.client.ExerciseTemplates.List(ctx, page, size)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to fetch exercise templates: %v", err)
	}

	data, err := json.MarshalIndent(result.ExerciseTemplates, "", "  ")
	if err != nil {
		return nil, nil, fmt.Errorf("failed to marshal exercise templates: %v", err)
	}

	hasMore := "no"
	if page < result.PageCount {
		hasMore = fmt.Sprintf("yes (next page: %d)", page+1)
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{
				Text: fmt.Sprintf("Fetched %d exercise templates (page: %d, size: %d, more pages: %s):\n\n%s", len(result.ExerciseTemplates), page, size, hasMore, string(data)),
			},
		},
	}, nil, nil
}

func (s svc) getExerciseTemplate(ctx context.Context, req *mcp.CallToolRequest, args ExerciseTemplateID) (*mcp.CallToolResult, any, error) {
	template, err := s.client.ExerciseTemplates.Get(ctx, args.ID)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to fetch exercise template: %v", err)
	}

	data, err := json.MarshalIndent(template, "", "  ")
	if err != nil {
		return nil, nil, fmt.Errorf("failed to marshal exercise template: %v", err)
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: string(data)},
		},
	}, nil, nil
}
