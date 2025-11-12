package main

import "github.com/swrm-io/go-hevy"

type svc struct {
	client *hevy.Client
}

// NoArgs is an empty struct for tools that do not require input.
type NoArgs struct{}

// Fetch represents the input parameters for fetching workouts.
type Fetch struct {
	Count int `json:"count" jsonschema:"Number of items to fetch"`
}
