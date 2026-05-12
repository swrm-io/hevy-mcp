package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type WorkoutID struct {
	ID string `json:"id" jsonschema:"The workout ID"`
}

func (s svc) getWorkoutCount(ctx context.Context, req *mcp.CallToolRequest, args NoArgs) (*mcp.CallToolResult, any, error) {
	count, err := s.client.Workouts.Count(ctx)
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

	result, err := s.client.Workouts.List(ctx, page, size)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to fetch workouts: %v", err)
	}

	data, err := json.MarshalIndent(result.Workouts, "", "  ")
	if err != nil {
		return nil, nil, fmt.Errorf("failed to marshal workouts: %v", err)
	}

	hasMore := "no"
	if page < result.PageCount {
		hasMore = fmt.Sprintf("yes (next page: %d)", page+1)
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{
				Text: fmt.Sprintf("Fetched %d workouts (page: %d, size: %d, more pages: %s):\n\n%s", len(result.Workouts), page, size, hasMore, string(data)),
			},
		},
	}, nil, nil
}

func (s svc) getWorkout(ctx context.Context, req *mcp.CallToolRequest, args WorkoutID) (*mcp.CallToolResult, any, error) {
	workout, err := s.client.Workouts.Get(ctx, args.ID)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to fetch workout: %v", err)
	}

	data, err := json.MarshalIndent(workout, "", "  ")
	if err != nil {
		return nil, nil, fmt.Errorf("failed to marshal workout: %v", err)
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: string(data)},
		},
	}, nil, nil
}

func (s svc) createWorkout(ctx context.Context, req *mcp.CallToolRequest, args WorkoutInput) (*mcp.CallToolResult, any, error) {
	workout, err := s.client.Workouts.Create(ctx, args.toLibType())
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create workout: %v", err)
	}

	data, err := json.MarshalIndent(workout, "", "  ")
	if err != nil {
		return nil, nil, fmt.Errorf("failed to marshal workout: %v", err)
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: string(data)},
		},
	}, nil, nil
}

func (s svc) updateWorkout(ctx context.Context, req *mcp.CallToolRequest, args UpdateWorkoutInput) (*mcp.CallToolResult, any, error) {
	workout, err := s.client.Workouts.Update(ctx, args.ID, args.Workout.toLibType())
	if err != nil {
		return nil, nil, fmt.Errorf("failed to update workout: %v", err)
	}

	data, err := json.MarshalIndent(workout, "", "  ")
	if err != nil {
		return nil, nil, fmt.Errorf("failed to marshal workout: %v", err)
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: string(data)},
		},
	}, nil, nil
}

func (s svc) getWorkoutEvents(ctx context.Context, req *mcp.CallToolRequest, args Fetch) (*mcp.CallToolResult, any, error) {
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

	result, err := s.client.Workouts.Events(ctx, page, size, nil)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to fetch workout events: %v", err)
	}

	data, err := json.MarshalIndent(result.Events, "", "  ")
	if err != nil {
		return nil, nil, fmt.Errorf("failed to marshal workout events: %v", err)
	}

	hasMore := "no"
	if page < result.PageCount {
		hasMore = fmt.Sprintf("yes (next page: %d)", page+1)
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{
				Text: fmt.Sprintf("Fetched %d workout events (page: %d, size: %d, more pages: %s):\n\n%s", len(result.Events), page, size, hasMore, string(data)),
			},
		},
	}, nil, nil
}
