package main

import (
	"context"
	"log"
	"log/slog"
	"os"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/swrm-io/go-hevy"
)

func main() {
	apiKey, ok := os.LookupEnv("HEVY_API_KEY")
	if !ok || apiKey == "" {
		slog.Error("HEVY_API_KEY environment variable is not set")
		os.Exit(1)
	}

	slog.Info("Starting Hevy MCP Server")

	client := hevy.New(apiKey)
	svc := svc{
		client: client,
	}
	server := mcp.NewServer(
		&mcp.Implementation{
			Name:    "hevy",
			Version: "v0.0.1",
		},
		nil,
	)

	mcp.AddTool(
		server,
		&mcp.Tool{
			Name:        "get_workout_count",
			Description: "Get the total count of workouts",
		},
		svc.getWorkoutCount,
	)

	mcp.AddTool(
		server,
		&mcp.Tool{
			Name:        "get_workouts",
			Description: "Get Workouts from newest to oldest",
		},
		svc.getWorkouts,
	)

	mcp.AddTool(
		server,
		&mcp.Tool{
			Name:        "get_workout",
			Description: "Get a single workout by ID",
		},
		svc.getWorkout,
	)

	mcp.AddTool(
		server,
		&mcp.Tool{
			Name:        "get_workout_events",
			Description: "Get workout update and delete events",
		},
		svc.getWorkoutEvents,
	)

	mcp.AddTool(
		server,
		&mcp.Tool{
			Name:        "get_routines",
			Description: "Get workout routines (templates) from the user's account",
		},
		svc.getRoutines,
	)

	mcp.AddTool(
		server,
		&mcp.Tool{
			Name:        "get_routine",
			Description: "Get a single workout routine by ID",
		},
		svc.getRoutine,
	)

	mcp.AddTool(
		server,
		&mcp.Tool{
			Name:        "get_exercise_templates",
			Description: "Get exercise templates (exercise library)",
		},
		svc.getExerciseTemplates,
	)

	mcp.AddTool(
		server,
		&mcp.Tool{
			Name:        "get_exercise_template",
			Description: "Get a single exercise template by ID",
		},
		svc.getExerciseTemplate,
	)

	mcp.AddTool(
		server,
		&mcp.Tool{
			Name:        "get_exercise_history",
			Description: "Get the full set history for a given exercise template",
		},
		svc.getExerciseHistory,
	)

	mcp.AddTool(
		server,
		&mcp.Tool{
			Name:        "get_body_measurements",
			Description: "Get body measurements from newest to oldest",
		},
		svc.getBodyMeasurements,
	)

	mcp.AddTool(
		server,
		&mcp.Tool{
			Name:        "get_body_measurement",
			Description: "Get body measurements for a specific date (YYYY-MM-DD)",
		},
		svc.getBodyMeasurement,
	)

	mcp.AddTool(
		server,
		&mcp.Tool{
			Name:        "get_user_info",
			Description: "Get basic info about the authenticated Hevy user",
		},
		svc.getUserInfo,
	)

	mcp.AddTool(
		server,
		&mcp.Tool{
			Name:        "get_routine_folders",
			Description: "Get routine folders",
		},
		svc.getRoutineFolders,
	)

	mcp.AddTool(
		server,
		&mcp.Tool{
			Name:        "get_routine_folder",
			Description: "Get a single routine folder by ID",
		},
		svc.getRoutineFolder,
	)

	// Start server with stdio transport
	if err := server.Run(context.Background(), &mcp.StdioTransport{}); err != nil {
		log.Fatal(err)
	}
}
