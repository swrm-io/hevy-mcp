package main

import (
	"time"

	hevy "github.com/swrm-io/go-hevy"
)

// --- Shared set/exercise input types ---

type SetInput struct {
	Type            string   `json:"type" jsonschema:"Set type: normal, warmup, failure, or dropset"`
	WeightKg        *float64 `json:"weight_kg,omitempty" jsonschema:"Weight in kilograms"`
	Reps            *int     `json:"reps,omitempty" jsonschema:"Number of repetitions"`
	DistanceMeters  *int     `json:"distance_meters,omitempty" jsonschema:"Distance in meters (for cardio exercises)"`
	DurationSeconds *int     `json:"duration_seconds,omitempty" jsonschema:"Duration in seconds (for timed exercises)"`
	RPE             *float64 `json:"rpe,omitempty" jsonschema:"Rating of Perceived Exertion (6, 7, 7.5, 8, 8.5, 9, 9.5, or 10)"`
	CustomMetric    *float64 `json:"custom_metric,omitempty" jsonschema:"Custom metric value for custom exercise types"`
}

func (s SetInput) toWorkoutSetInput() hevy.WorkoutSetInput {
	return hevy.WorkoutSetInput{
		Type:            hevy.SetType(s.Type),
		WeightKg:        s.WeightKg,
		Reps:            s.Reps,
		DistanceMeters:  s.DistanceMeters,
		DurationSeconds: s.DurationSeconds,
		RPE:             s.RPE,
		CustomMetric:    s.CustomMetric,
	}
}

type WorkoutExerciseInput struct {
	ExerciseTemplateID string     `json:"exercise_template_id" jsonschema:"The exercise template ID (from get_exercise_templates)"`
	SupersetID         *int       `json:"superset_id,omitempty" jsonschema:"Superset group ID; exercises sharing the same ID are supersetted"`
	Notes              *string    `json:"notes,omitempty" jsonschema:"Optional notes for this exercise"`
	Sets               []SetInput `json:"sets" jsonschema:"List of sets performed"`
}

func (e WorkoutExerciseInput) toLibType() hevy.WorkoutExerciseInput {
	sets := make([]hevy.WorkoutSetInput, len(e.Sets))
	for i, s := range e.Sets {
		sets[i] = s.toWorkoutSetInput()
	}
	return hevy.WorkoutExerciseInput{
		ExerciseTemplateID: e.ExerciseTemplateID,
		SupersetID:         e.SupersetID,
		Notes:              e.Notes,
		Sets:               sets,
	}
}

// --- Workout input types ---

type WorkoutInput struct {
	Title       string                 `json:"title" jsonschema:"Workout title"`
	Description *string                `json:"description,omitempty" jsonschema:"Optional workout description"`
	StartTime   time.Time              `json:"start_time" jsonschema:"Workout start time (RFC3339)"`
	EndTime     time.Time              `json:"end_time" jsonschema:"Workout end time (RFC3339)"`
	IsPrivate   *bool                  `json:"is_private,omitempty" jsonschema:"Whether the workout is private"`
	Exercises   []WorkoutExerciseInput `json:"exercises" jsonschema:"List of exercises in the workout"`
}

func (w WorkoutInput) toLibType() hevy.WorkoutInput {
	exercises := make([]hevy.WorkoutExerciseInput, len(w.Exercises))
	for i, e := range w.Exercises {
		exercises[i] = e.toLibType()
	}
	return hevy.WorkoutInput{
		Title:       w.Title,
		Description: w.Description,
		StartTime:   w.StartTime,
		EndTime:     w.EndTime,
		IsPrivate:   w.IsPrivate,
		Exercises:   exercises,
	}
}

type UpdateWorkoutInput struct {
	ID      string       `json:"id" jsonschema:"The workout ID to update"`
	Workout WorkoutInput `json:"workout" jsonschema:"The updated workout data"`
}

// --- Routine input types ---

type RepRangeInput struct {
	Start *float64 `json:"start,omitempty" jsonschema:"Minimum reps in the range"`
	End   *float64 `json:"end,omitempty" jsonschema:"Maximum reps in the range"`
}

type RoutineSetInput struct {
	Type            string         `json:"type" jsonschema:"Set type: normal, warmup, failure, or dropset"`
	WeightKg        *float64       `json:"weight_kg,omitempty" jsonschema:"Target weight in kilograms"`
	Reps            *int           `json:"reps,omitempty" jsonschema:"Target number of repetitions"`
	RepRange        *RepRangeInput `json:"rep_range,omitempty" jsonschema:"Target rep range (e.g. 8-12 reps)"`
	DistanceMeters  *int           `json:"distance_meters,omitempty" jsonschema:"Target distance in meters"`
	DurationSeconds *int           `json:"duration_seconds,omitempty" jsonschema:"Target duration in seconds"`
	CustomMetric    *float64       `json:"custom_metric,omitempty" jsonschema:"Custom metric value for custom exercise types"`
}

func (s RoutineSetInput) toLibType() hevy.RoutineSetInput {
	out := hevy.RoutineSetInput{
		Type:            hevy.SetType(s.Type),
		WeightKg:        s.WeightKg,
		Reps:            s.Reps,
		DistanceMeters:  s.DistanceMeters,
		DurationSeconds: s.DurationSeconds,
		CustomMetric:    s.CustomMetric,
	}
	if s.RepRange != nil {
		out.RepRange = &hevy.RepRange{Start: s.RepRange.Start, End: s.RepRange.End}
	}
	return out
}

type RoutineExerciseInput struct {
	ExerciseTemplateID string            `json:"exercise_template_id" jsonschema:"The exercise template ID (from get_exercise_templates)"`
	SupersetID         *int              `json:"superset_id,omitempty" jsonschema:"Superset group ID; exercises sharing the same ID are supersetted"`
	RestSeconds        *int              `json:"rest_seconds,omitempty" jsonschema:"Rest time between sets in seconds"`
	Notes              *string           `json:"notes,omitempty" jsonschema:"Optional notes for this exercise"`
	Sets               []RoutineSetInput `json:"sets" jsonschema:"List of sets in the routine exercise"`
}

func (e RoutineExerciseInput) toLibType() hevy.RoutineExerciseInput {
	sets := make([]hevy.RoutineSetInput, len(e.Sets))
	for i, s := range e.Sets {
		sets[i] = s.toLibType()
	}
	return hevy.RoutineExerciseInput{
		ExerciseTemplateID: e.ExerciseTemplateID,
		SupersetID:         e.SupersetID,
		RestSeconds:        e.RestSeconds,
		Notes:              e.Notes,
		Sets:               sets,
	}
}

type RoutineInput struct {
	Title     string                 `json:"title" jsonschema:"Routine title"`
	FolderID  *float64               `json:"folder_id,omitempty" jsonschema:"Optional routine folder ID to place the routine in"`
	Notes     string                 `json:"notes" jsonschema:"Optional notes for the routine"`
	Exercises []RoutineExerciseInput `json:"exercises" jsonschema:"List of exercises in the routine"`
}

func (r RoutineInput) toLibType() hevy.RoutineInput {
	exercises := make([]hevy.RoutineExerciseInput, len(r.Exercises))
	for i, e := range r.Exercises {
		exercises[i] = e.toLibType()
	}
	return hevy.RoutineInput{
		Title:     r.Title,
		FolderID:  r.FolderID,
		Notes:     r.Notes,
		Exercises: exercises,
	}
}

type RoutineUpdateInput struct {
	Title     string                 `json:"title" jsonschema:"Routine title"`
	FolderID  *float64               `json:"folder_id,omitempty" jsonschema:"Optional routine folder ID"`
	Notes     *string                `json:"notes,omitempty" jsonschema:"Optional notes for the routine"`
	Exercises []RoutineExerciseInput `json:"exercises" jsonschema:"List of exercises in the routine"`
}

func (r RoutineUpdateInput) toLibType() hevy.RoutineUpdateInput {
	exercises := make([]hevy.RoutineExerciseInput, len(r.Exercises))
	for i, e := range r.Exercises {
		exercises[i] = e.toLibType()
	}
	return hevy.RoutineUpdateInput{
		Title:     r.Title,
		FolderID:  r.FolderID,
		Notes:     r.Notes,
		Exercises: exercises,
	}
}

type UpdateRoutineInput struct {
	ID      string             `json:"id" jsonschema:"The routine ID to update"`
	Routine RoutineUpdateInput `json:"routine" jsonschema:"The updated routine data"`
}

// --- Exercise template input types ---

type CreateExerciseInput struct {
	Title             string   `json:"title" jsonschema:"Exercise name"`
	ExerciseType      string   `json:"exercise_type" jsonschema:"Tracking type: weight_reps, reps_only, bodyweight_reps, bodyweight_assisted_reps, duration, weight_duration, distance_duration, or short_distance_weight"`
	EquipmentCategory string   `json:"equipment_category" jsonschema:"Equipment used: none, barbell, dumbbell, kettlebell, machine, plate, resistance_band, suspension, or other"`
	MuscleGroup       string   `json:"muscle_group" jsonschema:"Primary muscle group: abdominals, shoulders, biceps, triceps, forearms, quadriceps, hamstrings, calves, glutes, abductors, adductors, lats, upper_back, traps, lower_back, chest, cardio, neck, full_body, or other"`
	OtherMuscles      []string `json:"other_muscles,omitempty" jsonschema:"Additional muscle groups worked (same values as muscle_group)"`
}

func (e CreateExerciseInput) toLibType() hevy.CreateExerciseInput {
	others := make([]hevy.MuscleGroup, len(e.OtherMuscles))
	for i, m := range e.OtherMuscles {
		others[i] = hevy.MuscleGroup(m)
	}
	return hevy.CreateExerciseInput{
		Title:             e.Title,
		ExerciseType:      hevy.CustomExerciseType(e.ExerciseType),
		EquipmentCategory: hevy.EquipmentCategory(e.EquipmentCategory),
		MuscleGroup:       hevy.MuscleGroup(e.MuscleGroup),
		OtherMuscles:      others,
	}
}

// --- Body measurement input types ---

type BodyMeasurementInput struct {
	Date           string   `json:"date" jsonschema:"Date in YYYY-MM-DD format"`
	WeightKg       *float64 `json:"weight_kg,omitempty" jsonschema:"Body weight in kilograms"`
	LeanMassKg     *float64 `json:"lean_mass_kg,omitempty" jsonschema:"Lean body mass in kilograms"`
	FatPercent     *float64 `json:"fat_percent,omitempty" jsonschema:"Body fat percentage"`
	NeckCm         *float64 `json:"neck_cm,omitempty" jsonschema:"Neck circumference in centimeters"`
	ShoulderCm     *float64 `json:"shoulder_cm,omitempty" jsonschema:"Shoulder circumference in centimeters"`
	ChestCm        *float64 `json:"chest_cm,omitempty" jsonschema:"Chest circumference in centimeters"`
	LeftBicepCm    *float64 `json:"left_bicep_cm,omitempty" jsonschema:"Left bicep circumference in centimeters"`
	RightBicepCm   *float64 `json:"right_bicep_cm,omitempty" jsonschema:"Right bicep circumference in centimeters"`
	LeftForearmCm  *float64 `json:"left_forearm_cm,omitempty" jsonschema:"Left forearm circumference in centimeters"`
	RightForearmCm *float64 `json:"right_forearm_cm,omitempty" jsonschema:"Right forearm circumference in centimeters"`
	AbdomenCm      *float64 `json:"abdomen,omitempty" jsonschema:"Abdomen circumference in centimeters"`
	WaistCm        *float64 `json:"waist,omitempty" jsonschema:"Waist circumference in centimeters"`
	HipsCm         *float64 `json:"hips,omitempty" jsonschema:"Hips circumference in centimeters"`
	LeftThighCm    *float64 `json:"left_thigh,omitempty" jsonschema:"Left thigh circumference in centimeters"`
	RightThighCm   *float64 `json:"right_thigh,omitempty" jsonschema:"Right thigh circumference in centimeters"`
	LeftCalfCm     *float64 `json:"left_calf,omitempty" jsonschema:"Left calf circumference in centimeters"`
	RightCalfCm    *float64 `json:"right_calf,omitempty" jsonschema:"Right calf circumference in centimeters"`
}

func (b BodyMeasurementInput) toLibType() hevy.BodyMeasurement {
	return hevy.BodyMeasurement{
		Date:           b.Date,
		WeightKg:       b.WeightKg,
		LeanMassKg:     b.LeanMassKg,
		FatPercent:     b.FatPercent,
		NeckCm:         b.NeckCm,
		ShoulderCm:     b.ShoulderCm,
		ChestCm:        b.ChestCm,
		LeftBicepCm:    b.LeftBicepCm,
		RightBicepCm:   b.RightBicepCm,
		LeftForearmCm:  b.LeftForearmCm,
		RightForearmCm: b.RightForearmCm,
		AbdomenCm:      b.AbdomenCm,
		WaistCm:        b.WaistCm,
		HipsCm:         b.HipsCm,
		LeftThighCm:    b.LeftThighCm,
		RightThighCm:   b.RightThighCm,
		LeftCalfCm:     b.LeftCalfCm,
		RightCalfCm:    b.RightCalfCm,
	}
}

type BodyMeasurementUpdateInput struct {
	Date        string   `json:"date" jsonschema:"Date in YYYY-MM-DD format"`
	WeightKg    *float64 `json:"weight_kg" jsonschema:"Body weight in kilograms (null clears the value)"`
	LeanMassKg  *float64 `json:"lean_mass_kg" jsonschema:"Lean body mass in kilograms (null clears the value)"`
	FatPercent  *float64 `json:"fat_percent" jsonschema:"Body fat percentage (null clears the value)"`
	NeckCm      *float64 `json:"neck_cm" jsonschema:"Neck circumference in centimeters (null clears the value)"`
	ShoulderCm  *float64 `json:"shoulder_cm" jsonschema:"Shoulder circumference in centimeters (null clears the value)"`
	ChestCm     *float64 `json:"chest_cm" jsonschema:"Chest circumference in centimeters (null clears the value)"`
	LeftBicepCm  *float64 `json:"left_bicep_cm" jsonschema:"Left bicep circumference in centimeters (null clears the value)"`
	RightBicepCm *float64 `json:"right_bicep_cm" jsonschema:"Right bicep circumference in centimeters (null clears the value)"`
	LeftForearmCm  *float64 `json:"left_forearm_cm" jsonschema:"Left forearm circumference in centimeters (null clears the value)"`
	RightForearmCm *float64 `json:"right_forearm_cm" jsonschema:"Right forearm circumference in centimeters (null clears the value)"`
	AbdomenCm   *float64 `json:"abdomen" jsonschema:"Abdomen circumference in centimeters (null clears the value)"`
	WaistCm     *float64 `json:"waist" jsonschema:"Waist circumference in centimeters (null clears the value)"`
	HipsCm      *float64 `json:"hips" jsonschema:"Hips circumference in centimeters (null clears the value)"`
	LeftThighCm  *float64 `json:"left_thigh" jsonschema:"Left thigh circumference in centimeters (null clears the value)"`
	RightThighCm *float64 `json:"right_thigh" jsonschema:"Right thigh circumference in centimeters (null clears the value)"`
	LeftCalfCm   *float64 `json:"left_calf" jsonschema:"Left calf circumference in centimeters (null clears the value)"`
	RightCalfCm  *float64 `json:"right_calf" jsonschema:"Right calf circumference in centimeters (null clears the value)"`
}

func (b BodyMeasurementUpdateInput) toUpdateLibType() hevy.BodyMeasurementUpdate {
	return hevy.BodyMeasurementUpdate{
		WeightKg:       b.WeightKg,
		LeanMassKg:     b.LeanMassKg,
		FatPercent:     b.FatPercent,
		NeckCm:         b.NeckCm,
		ShoulderCm:     b.ShoulderCm,
		ChestCm:        b.ChestCm,
		LeftBicepCm:    b.LeftBicepCm,
		RightBicepCm:   b.RightBicepCm,
		LeftForearmCm:  b.LeftForearmCm,
		RightForearmCm: b.RightForearmCm,
		AbdomenCm:      b.AbdomenCm,
		WaistCm:        b.WaistCm,
		HipsCm:         b.HipsCm,
		LeftThighCm:    b.LeftThighCm,
		RightThighCm:   b.RightThighCm,
		LeftCalfCm:     b.LeftCalfCm,
		RightCalfCm:    b.RightCalfCm,
	}
}
