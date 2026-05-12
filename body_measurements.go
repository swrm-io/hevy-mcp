package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type MeasurementDate struct {
	Date string `json:"date" jsonschema:"Date in YYYY-MM-DD format"`
}

func (s svc) getBodyMeasurements(ctx context.Context, req *mcp.CallToolRequest, args Fetch) (*mcp.CallToolResult, any, error) {
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

	result, err := s.client.BodyMeasurements.List(ctx, page, size)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to fetch body measurements: %v", err)
	}

	data, err := json.MarshalIndent(result.Measurements, "", "  ")
	if err != nil {
		return nil, nil, fmt.Errorf("failed to marshal body measurements: %v", err)
	}

	hasMore := "no"
	if page < result.PageCount {
		hasMore = fmt.Sprintf("yes (next page: %d)", page+1)
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{
				Text: fmt.Sprintf("Fetched %d body measurements (page: %d, size: %d, more pages: %s):\n\n%s", len(result.Measurements), page, size, hasMore, string(data)),
			},
		},
	}, nil, nil
}

func (s svc) createBodyMeasurement(ctx context.Context, req *mcp.CallToolRequest, args BodyMeasurementInput) (*mcp.CallToolResult, any, error) {
	if err := s.client.BodyMeasurements.Create(ctx, args.toLibType()); err != nil {
		return nil, nil, fmt.Errorf("failed to create body measurement: %v", err)
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: fmt.Sprintf("Created body measurement for %s", args.Date)},
		},
	}, nil, nil
}

func (s svc) updateBodyMeasurement(ctx context.Context, req *mcp.CallToolRequest, args BodyMeasurementUpdateInput) (*mcp.CallToolResult, any, error) {
	if err := s.client.BodyMeasurements.Update(ctx, args.Date, args.toUpdateLibType()); err != nil {
		return nil, nil, fmt.Errorf("failed to update body measurement: %v", err)
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: fmt.Sprintf("Updated body measurement for %s", args.Date)},
		},
	}, nil, nil
}

func (s svc) getBodyMeasurement(ctx context.Context, req *mcp.CallToolRequest, args MeasurementDate) (*mcp.CallToolResult, any, error) {
	measurement, err := s.client.BodyMeasurements.Get(ctx, args.Date)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to fetch body measurement: %v", err)
	}

	data, err := json.MarshalIndent(measurement, "", "  ")
	if err != nil {
		return nil, nil, fmt.Errorf("failed to marshal body measurement: %v", err)
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: string(data)},
		},
	}, nil, nil
}
