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
			Description: "Get the total number of workouts logged by the user.",
		},
		svc.getWorkoutCount,
	)

	mcp.AddTool(
		server,
		&mcp.Tool{
			Name:        "get_workouts",
			Description: "List workouts from newest to oldest, paginated. Default page size is 5, max is 10. Use page and size to paginate through results.",
		},
		svc.getWorkouts,
	)

	mcp.AddTool(
		server,
		&mcp.Tool{
			Name:        "get_workout",
			Description: "Get full details of a single workout by its ID, including all exercises and sets. Workout IDs are returned by get_workouts.",
		},
		svc.getWorkout,
	)

	mcp.AddTool(
		server,
		&mcp.Tool{
			Name:        "create_workout",
			Description: "Log a new completed workout. Each exercise requires an exercise_template_id — use get_exercise_templates to find valid IDs. start_time and end_time must be RFC3339 timestamps.",
		},
		svc.createWorkout,
	)

	mcp.AddTool(
		server,
		&mcp.Tool{
			Name:        "update_workout",
			Description: "Replace all fields of an existing workout. Requires the full workout payload, not just changed fields. Workout IDs are returned by get_workouts.",
		},
		svc.updateWorkout,
	)

	mcp.AddTool(
		server,
		&mcp.Tool{
			Name:        "get_workout_events",
			Description: "List workout change events (updates and deletions), useful for syncing. Events are paginated newest-first. Each event includes the type (updated or deleted) and the affected workout data or ID.",
		},
		svc.getWorkoutEvents,
	)

	mcp.AddTool(
		server,
		&mcp.Tool{
			Name:        "get_routines",
			Description: "List the user's saved workout routines (reusable templates), paginated. Default page size is 5, max is 10.",
		},
		svc.getRoutines,
	)

	mcp.AddTool(
		server,
		&mcp.Tool{
			Name:        "get_routine",
			Description: "Get full details of a single routine by its ID, including all exercises and sets. Routine IDs are returned by get_routines.",
		},
		svc.getRoutine,
	)

	mcp.AddTool(
		server,
		&mcp.Tool{
			Name:        "create_routine",
			Description: "Create a new workout routine (reusable template). Each exercise requires an exercise_template_id — use get_exercise_templates to find valid IDs. Optionally assign to a folder using a folder_id from get_routine_folders.",
		},
		svc.createRoutine,
	)

	mcp.AddTool(
		server,
		&mcp.Tool{
			Name:        "update_routine",
			Description: "Replace all fields of an existing routine. Requires the full routine payload, not just changed fields. Routine IDs are returned by get_routines. Note: the routine's folder cannot be changed via update (the Hevy API rejects folder_id on this endpoint) — a routine can only be assigned to a folder at creation time via create_routine.",
		},
		svc.updateRoutine,
	)

	mcp.AddTool(
		server,
		&mcp.Tool{
			Name:        "get_exercise_templates",
			Description: "List the exercise library (built-in and custom exercises), paginated. Default page size is 20, max is 100. Returns exercise_template_id values needed for create_workout and create_routine.",
		},
		svc.getExerciseTemplates,
	)

	mcp.AddTool(
		server,
		&mcp.Tool{
			Name:        "get_exercise_template",
			Description: "Get details of a single exercise template by its ID, including muscle groups and equipment. IDs are returned by get_exercise_templates.",
		},
		svc.getExerciseTemplate,
	)

	mcp.AddTool(
		server,
		&mcp.Tool{
			Name:        "create_exercise_template",
			Description: "Create a custom exercise in the user's exercise library. Returns the new exercise template ID, which can then be used in create_workout or create_routine.",
		},
		svc.createExerciseTemplate,
	)

	mcp.AddTool(
		server,
		&mcp.Tool{
			Name:        "get_exercise_history",
			Description: "Get the full logged set history for a specific exercise across all past workouts. Requires an exercise_template_id from get_exercise_templates. Useful for tracking progress over time.",
		},
		svc.getExerciseHistory,
	)

	mcp.AddTool(
		server,
		&mcp.Tool{
			Name:        "get_body_measurements",
			Description: "List body measurement entries from newest to oldest, paginated. Default page size is 5, max is 10. Measurements include weight, body fat, and circumference measurements.",
		},
		svc.getBodyMeasurements,
	)

	mcp.AddTool(
		server,
		&mcp.Tool{
			Name:        "get_body_measurement",
			Description: "Get the body measurement entry for a specific date (YYYY-MM-DD format).",
		},
		svc.getBodyMeasurement,
	)

	mcp.AddTool(
		server,
		&mcp.Tool{
			Name:        "create_body_measurement",
			Description: "Log a new body measurement entry for a date. Returns an error if an entry already exists for that date — use update_body_measurement instead.",
		},
		svc.createBodyMeasurement,
	)

	mcp.AddTool(
		server,
		&mcp.Tool{
			Name:        "update_body_measurement",
			Description: "Update body measurements for a specific date (YYYY-MM-DD). This is a full replace: any field not provided will be set to null on the server, clearing its value.",
		},
		svc.updateBodyMeasurement,
	)

	mcp.AddTool(
		server,
		&mcp.Tool{
			Name:        "get_user_info",
			Description: "Get the authenticated user's Hevy account info (ID, display name, and profile URL).",
		},
		svc.getUserInfo,
	)

	mcp.AddTool(
		server,
		&mcp.Tool{
			Name:        "get_routine_folders",
			Description: "List the user's routine folders, paginated. Default page size is 5, max is 10. Folders are used to organise routines.",
		},
		svc.getRoutineFolders,
	)

	mcp.AddTool(
		server,
		&mcp.Tool{
			Name:        "get_routine_folder",
			Description: "Get a single routine folder by its ID. Folder IDs are returned by get_routine_folders.",
		},
		svc.getRoutineFolder,
	)

	mcp.AddTool(
		server,
		&mcp.Tool{
			Name:        "create_routine_folder",
			Description: "Create a new routine folder with the given title. Returns the new folder, including its ID which can be used in create_routine or update_routine.",
		},
		svc.createRoutineFolder,
	)

	// Start server with stdio transport
	if err := server.Run(context.Background(), &mcp.StdioTransport{}); err != nil {
		log.Fatal(err)
	}
}
