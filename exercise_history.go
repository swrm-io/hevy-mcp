package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type ExerciseHistoryArgs struct {
	ExerciseTemplateID string `json:"exercise_template_id" jsonschema:"The exercise template ID to fetch history for"`
}

func (s svc) getExerciseHistory(ctx context.Context, req *mcp.CallToolRequest, args ExerciseHistoryArgs) (*mcp.CallToolResult, any, error) {
	history, err := s.client.ExerciseHistory.Get(ctx, args.ExerciseTemplateID, nil)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to fetch exercise history: %v", err)
	}

	data, err := json.MarshalIndent(history, "", "  ")
	if err != nil {
		return nil, nil, fmt.Errorf("failed to marshal exercise history: %v", err)
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{
				Text: fmt.Sprintf("Fetched %d history entries for exercise template %s:\n\n%s", len(history), args.ExerciseTemplateID, string(data)),
			},
		},
	}, nil, nil
}
