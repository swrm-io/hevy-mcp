package main

import "github.com/swrm-io/go-hevy"

type svc struct {
	client *hevy.Client
}

// NoArgs is an empty struct for tools that do not require input.
type NoArgs struct{}

// Fetch represents the input parameters for fetching workouts.
type Fetch struct {
	Page int `json:"page" jsonschema:"Page number to fetch (default: 1)"`
	Size int `json:"size" jsonschema:"Number of workouts per page (default: 5, max: 10)"`
}
