package main

import (
	"encoding/json"
	"testing"
)

// Test helper to create string pointers
func stringPtr(s string) *string {
	return &s
}

// Test data helper
func createValidJsonExercise() jsonExercise {
	return jsonExercise{
		SourceID:         "exercise-123",
		Name:             "bench press",
		Force:            stringPtr("push"),
		Level:            "beginner",
		Mechanic:         stringPtr("compound"),
		Equipment:        stringPtr("barbell"),
		PrimaryMuscles:   []string{"chest", "triceps"},
		SecondaryMuscles: []string{"shoulders"},
		Instructions:     []string{"Lie on a flat bench.", "Grip the barbell."},
		Category:         "strength",
		Images:           []string{"image1.jpg", "image2.jpg"},
	}
}

func TestValidateMetadata(t *testing.T) {
	t.Run("valid metadata", func(t *testing.T) {
		input := createValidJsonExercise()
		result, err := validateMetadata(input)
		if err != nil {
			t.Errorf("validateMetadata() unexpected error: %v", err)
		}
		// Check normalization
		if result.SourceID != "Exercise-123" {
			t.Errorf("validateMetadata() SourceID = %q, expected %q", result.SourceID, "Exercise-123")
		}
		if result.Name != "Bench Press" {
			t.Errorf("validateMetadata() Name = %q, expected %q", result.Name, "Bench Press")
		}
		if result.Level != "Beginner" {
			t.Errorf("validateMetadata() Level = %q, expected %q", result.Level, "Beginner")
		}
		if result.Category != "Strength" {
			t.Errorf("validateMetadata() Category = %q, expected %q", result.Category, "Strength")
		}
		if *result.Force != "Push" {
			t.Errorf("validateMetadata() Force = %q, expected %q", *result.Force, "Push")
		}
		if *result.Mechanic != "Compound" {
			t.Errorf("validateMetadata() Mechanic = %q, expected %q", *result.Mechanic, "Compound")
		}
		if *result.Equipment != "Barbell" {
			t.Errorf("validateMetadata() Equipment = %q, expected %q", *result.Equipment, "Barbell")
		}
		if len(result.PrimaryMuscles) != 2 || result.PrimaryMuscles[0] != "Chest" || result.PrimaryMuscles[1] != "Triceps" {
			t.Errorf("validateMetadata() PrimaryMuscles = %v, expected %v", result.PrimaryMuscles, []string{"Chest", "Triceps"})
		}
		if len(result.SecondaryMuscles) != 1 || result.SecondaryMuscles[0] != "Shoulders" {
			t.Errorf("validateMetadata() SecondaryMuscles = %v, expected %v", result.SecondaryMuscles, []string{"Shoulders"})
		}
		if len(result.Instructions) != 2 || result.Instructions[0] != "Lie On A Flat Bench." || result.Instructions[1] != "Grip The Barbell." {
			t.Errorf("validateMetadata() Instructions = %v, expected normalized instructions", result.Instructions)
		}
	})

	t.Run("empty source id", func(t *testing.T) {
		input := createValidJsonExercise()
		input.SourceID = ""
		_, err := validateMetadata(input)
		if err == nil {
			t.Error("validateMetadata() expected error for empty source id")
		} else if err.Error() != "source id is empty" {
			t.Errorf("validateMetadata() error = %v, expected %q", err, "source id is empty")
		}
	})

	t.Run("whitespace source id", func(t *testing.T) {
		input := createValidJsonExercise()
		input.SourceID = "   "
		_, err := validateMetadata(input)
		if err == nil {
			t.Error("validateMetadata() expected error for whitespace source id")
		} else if err.Error() != "source id is empty" {
			t.Errorf("validateMetadata() error = %v, expected %q", err, "source id is empty")
		}
	})

	t.Run("empty name", func(t *testing.T) {
		input := createValidJsonExercise()
		input.Name = ""
		_, err := validateMetadata(input)
		if err == nil {
			t.Error("validateMetadata() expected error for empty name")
		} else if err.Error() != "name is empty" {
			t.Errorf("validateMetadata() error = %v, expected %q", err, "name is empty")
		}
	})

	t.Run("empty level", func(t *testing.T) {
		input := createValidJsonExercise()
		input.Level = ""
		_, err := validateMetadata(input)
		if err == nil {
			t.Error("validateMetadata() expected error for empty level")
		} else if err.Error() != "level is empty" {
			t.Errorf("validateMetadata() error = %v, expected %q", err, "level is empty")
		}
	})

	t.Run("empty category", func(t *testing.T) {
		input := createValidJsonExercise()
		input.Category = ""
		_, err := validateMetadata(input)
		if err == nil {
			t.Error("validateMetadata() expected error for empty category")
		} else if err.Error() != "category is empty" {
			t.Errorf("validateMetadata() error = %v, expected %q", err, "category is empty")
		}
	})

	t.Run("nil force pointer", func(t *testing.T) {
		input := createValidJsonExercise()
		input.Force = nil
		result, err := validateMetadata(input)
		if err != nil {
			t.Errorf("validateMetadata() unexpected error for nil force: %v", err)
		}
		if result.Force != nil {
			t.Errorf("validateMetadata() Force should be nil, got %v", result.Force)
		}
	})

	t.Run("empty force string", func(t *testing.T) {
		input := createValidJsonExercise()
		empty := ""
		input.Force = &empty
		_, err := validateMetadata(input)
		if err == nil {
			t.Error("validateMetadata() expected error for empty force string")
		} else if err.Error() != "force is empty string" {
			t.Errorf("validateMetadata() error = %v, expected %q", err, "force is empty string")
		}
	})

	t.Run("whitespace force string", func(t *testing.T) {
		input := createValidJsonExercise()
		whitespace := "   \t\n  "
		input.Force = &whitespace
		_, err := validateMetadata(input)
		if err == nil {
			t.Error("validateMetadata() expected error for whitespace force string")
		} else if err.Error() != "force is empty string" {
			t.Errorf("validateMetadata() error = %v, expected %q", err, "force is empty string")
		}
	})

	t.Run("nil mechanic pointer", func(t *testing.T) {
		input := createValidJsonExercise()
		input.Mechanic = nil
		result, err := validateMetadata(input)
		if err != nil {
			t.Errorf("validateMetadata() unexpected error for nil mechanic: %v", err)
		}
		if result.Mechanic != nil {
			t.Errorf("validateMetadata() Mechanic should be nil, got %v", result.Mechanic)
		}
	})

	t.Run("empty mechanic string", func(t *testing.T) {
		input := createValidJsonExercise()
		empty := ""
		input.Mechanic = &empty
		_, err := validateMetadata(input)
		if err == nil {
			t.Error("validateMetadata() expected error for empty mechanic string")
		} else if err.Error() != "mechanic is empty string" {
			t.Errorf("validateMetadata() error = %v, expected %q", err, "mechanic is empty string")
		}
	})

	t.Run("nil equipment pointer", func(t *testing.T) {
		input := createValidJsonExercise()
		input.Equipment = nil
		result, err := validateMetadata(input)
		if err != nil {
			t.Errorf("validateMetadata() unexpected error for nil equipment: %v", err)
		}
		if result.Equipment != nil {
			t.Errorf("validateMetadata() Equipment should be nil, got %v", result.Equipment)
		}
	})

	t.Run("empty equipment string", func(t *testing.T) {
		input := createValidJsonExercise()
		empty := ""
		input.Equipment = &empty
		_, err := validateMetadata(input)
		if err == nil {
			t.Error("validateMetadata() expected error for empty equipment string")
		} else if err.Error() != "equipment is empty string" {
			t.Errorf("validateMetadata() error = %v, expected %q", err, "equipment is empty string")
		}
	})

	t.Run("empty primary muscle", func(t *testing.T) {
		input := createValidJsonExercise()
		input.PrimaryMuscles = []string{"chest", "", "triceps"}
		_, err := validateMetadata(input)
		if err == nil {
			t.Error("validateMetadata() expected error for empty primary muscle")
		} else if err.Error() != "primary muscles[1] is empty" {
			t.Errorf("validateMetadata() error = %v, expected %q", err, "primary muscles[1] is empty")
		}
	})

	t.Run("whitespace primary muscle", func(t *testing.T) {
		input := createValidJsonExercise()
		input.PrimaryMuscles = []string{"chest", "   ", "triceps"}
		_, err := validateMetadata(input)
		if err == nil {
			t.Error("validateMetadata() expected error for whitespace primary muscle")
		} else if err.Error() != "primary muscles[1] is empty" {
			t.Errorf("validateMetadata() error = %v, expected %q", err, "primary muscles[1] is empty")
		}
	})

	t.Run("empty secondary muscle", func(t *testing.T) {
		input := createValidJsonExercise()
		input.SecondaryMuscles = []string{"shoulders", ""}
		_, err := validateMetadata(input)
		if err == nil {
			t.Error("validateMetadata() expected error for empty secondary muscle")
		} else if err.Error() != "secondary muscles[1] is empty" {
			t.Errorf("validateMetadata() error = %v, expected %q", err, "secondary muscles[1] is empty")
		}
	})

	t.Run("filter empty instructions", func(t *testing.T) {
		input := createValidJsonExercise()
		input.Instructions = []string{"First step.", "", "   ", "Second step."}
		result, err := validateMetadata(input)
		if err != nil {
			t.Errorf("validateMetadata() unexpected error: %v", err)
		}
		if len(result.Instructions) != 2 {
			t.Errorf("validateMetadata() Instructions length = %d, expected %d", len(result.Instructions), 2)
		}
		if result.Instructions[0] != "First Step." {
			t.Errorf("validateMetadata() Instructions[0] = %q, expected %q", result.Instructions[0], "First Step.")
		}
		if result.Instructions[1] != "Second Step." {
			t.Errorf("validateMetadata() Instructions[1] = %q, expected %q", result.Instructions[1], "Second Step.")
		}
	})

	t.Run("empty image path", func(t *testing.T) {
		input := createValidJsonExercise()
		input.Images = []string{"image1.jpg", "", "image2.jpg"}
		_, err := validateMetadata(input)
		if err == nil {
			t.Error("validateMetadata() expected error for empty image path")
		} else if err.Error() != "images[1] path is empty" {
			t.Errorf("validateMetadata() error = %v, expected %q", err, "images[1] path is empty")
		}
	})

	t.Run("whitespace image path", func(t *testing.T) {
		input := createValidJsonExercise()
		input.Images = []string{"image1.jpg", "   ", "image2.jpg"}
		_, err := validateMetadata(input)
		if err == nil {
			t.Error("validateMetadata() expected error for whitespace image path")
		} else if err.Error() != "images[1] path is empty" {
			t.Errorf("validateMetadata() error = %v, expected %q", err, "images[1] path is empty")
		}
	})

	t.Run("normalization of all fields", func(t *testing.T) {
		input := jsonExercise{
			SourceID:         "  EXERCISE-456  ",
			Name:             "  dead lift  ",
			Force:            stringPtr("  PULL  "),
			Level:            "  ADVANCED  ",
			Mechanic:         stringPtr("  COMPOUND  "),
			Equipment:        stringPtr("  BARBELL  "),
			PrimaryMuscles:   []string{"  BACK  ", "HAMSTRINGS  "},
			SecondaryMuscles: []string{" GLUTES ", " CALVES "},
			Instructions:     []string{"  stand with feet shoulder-width apart.  ", "  grip the barbell.  "},
			Category:         "  STRENGTH  ",
			Images:           []string{"  deadlift.jpg  ", "deadlift2.jpg"},
		}
		result, err := validateMetadata(input)
		if err != nil {
			t.Errorf("validateMetadata() unexpected error: %v", err)
		}
		if result.SourceID != "Exercise-456" {
			t.Errorf("validateMetadata() SourceID = %q, expected %q", result.SourceID, "Exercise-456")
		}
		if result.Name != "Dead Lift" {
			t.Errorf("validateMetadata() Name = %q, expected %q", result.Name, "Dead Lift")
		}
		if *result.Force != "Pull" {
			t.Errorf("validateMetadata() Force = %q, expected %q", *result.Force, "Pull")
		}
		if result.Level != "Advanced" {
			t.Errorf("validateMetadata() Level = %q, expected %q", result.Level, "Advanced")
		}
		if *result.Mechanic != "Compound" {
			t.Errorf("validateMetadata() Mechanic = %q, expected %q", *result.Mechanic, "Compound")
		}
		if *result.Equipment != "Barbell" {
			t.Errorf("validateMetadata() Equipment = %q, expected %q", *result.Equipment, "Barbell")
		}
		if len(result.PrimaryMuscles) != 2 || result.PrimaryMuscles[0] != "Back" || result.PrimaryMuscles[1] != "Hamstrings" {
			t.Errorf("validateMetadata() PrimaryMuscles = %v, expected %v", result.PrimaryMuscles, []string{"Back", "Hamstrings"})
		}
		if len(result.SecondaryMuscles) != 2 || result.SecondaryMuscles[0] != "Glutes" || result.SecondaryMuscles[1] != "Calves" {
			t.Errorf("validateMetadata() SecondaryMuscles = %v, expected %v", result.SecondaryMuscles, []string{"Glutes", "Calves"})
		}
		if result.Category != "Strength" {
			t.Errorf("validateMetadata() Category = %q, expected %q", result.Category, "Strength")
		}
		if result.Images[0] != "deadlift.jpg" {
			t.Errorf("validateMetadata() Images[0] = %q, expected %q", result.Images[0], "deadlift.jpg")
		}
		if result.Images[1] != "deadlift2.jpg" {
			t.Errorf("validateMetadata() Images[1] = %q, expected %q", result.Images[1], "deadlift2.jpg")
		}
	})
}

func TestNewExercise(t *testing.T) {
	t.Run("create exercise with normalization", func(t *testing.T) {
		force := "push"
		mechanic := "compound"
		equipment := "barbell"
		images := []Image{
			{ImageBlob: []byte("fake image"), MimeType: "image/jpeg"},
		}

		exercise := NewExercise(
			"exercise-789",
			"  squat  ",
			&force,
			"  intermediate  ",
			&mechanic,
			&equipment,
			[]string{"quads", " glutes "},
			[]string{" hamstrings ", "calves"},
			[]string{"stand with feet shoulder-width apart", "lower your body"},
			"strength",
			images,
		)

		if exercise.SourceId != "exercise-789" {
			t.Errorf("NewExercise() SourceId = %q, expected %q", exercise.SourceId, "exercise-789")
		}
		if exercise.Name != "Squat" {
			t.Errorf("NewExercise() Name = %q, expected %q", exercise.Name, "Squat")
		}
		if *exercise.Force != "Push" {
			t.Errorf("NewExercise() Force = %q, expected %q", *exercise.Force, "Push")
		}
		if exercise.Level != "Intermediate" {
			t.Errorf("NewExercise() Level = %q, expected %q", exercise.Level, "Intermediate")
		}
		if *exercise.Mechanic != "Compound" {
			t.Errorf("NewExercise() Mechanic = %q, expected %q", *exercise.Mechanic, "Compound")
		}
		if *exercise.Equipment != "Barbell" {
			t.Errorf("NewExercise() Equipment = %q, expected %q", *exercise.Equipment, "Barbell")
		}
		if len(exercise.PrimaryMuscles) != 2 || exercise.PrimaryMuscles[0] != "Quads" || exercise.PrimaryMuscles[1] != "Glutes" {
			t.Errorf("NewExercise() PrimaryMuscles = %v, expected %v", exercise.PrimaryMuscles, []string{"Quads", "Glutes"})
		}
		if len(exercise.SecondaryMuscles) != 2 || exercise.SecondaryMuscles[0] != "Hamstrings" || exercise.SecondaryMuscles[1] != "Calves" {
			t.Errorf("NewExercise() SecondaryMuscles = %v, expected %v", exercise.SecondaryMuscles, []string{"Hamstrings", "Calves"})
		}
		if len(exercise.Instructions) != 2 {
			t.Errorf("NewExercise() Instructions length = %d, expected %d", len(exercise.Instructions), 2)
		}
		if exercise.Category != "Strength" {
			t.Errorf("NewExercise() Category = %q, expected %q", exercise.Category, "Strength")
		}
		if len(exercise.Images) != 1 {
			t.Errorf("NewExercise() Images length = %d, expected %d", len(exercise.Images), 1)
		}
	})

	t.Run("create exercise with nil pointers", func(t *testing.T) {
		exercise := NewExercise(
			"exercise-999",
			"pull-up",
			nil,
			"beginner",
			nil,
			nil,
			[]string{"back", "biceps"},
			[]string{"forearms"},
			[]string{"grip the bar", "pull yourself up"},
			"bodyweight",
			nil,
		)

		if exercise.SourceId != "exercise-999" {
			t.Errorf("NewExercise() SourceId = %q, expected %q", exercise.SourceId, "exercise-999")
		}
		if exercise.Name != "Pull-Up" {
			t.Errorf("NewExercise() Name = %q, expected %q", exercise.Name, "Pull-Up")
		}
		if exercise.Force != nil {
			t.Errorf("NewExercise() Force should be nil, got %v", exercise.Force)
		}
		if exercise.Level != "Beginner" {
			t.Errorf("NewExercise() Level = %q, expected %q", exercise.Level, "Beginner")
		}
		if exercise.Mechanic != nil {
			t.Errorf("NewExercise() Mechanic should be nil, got %v", exercise.Mechanic)
		}
		if exercise.Equipment != nil {
			t.Errorf("NewExercise() Equipment should be nil, got %v", exercise.Equipment)
		}
		if len(exercise.PrimaryMuscles) != 2 || exercise.PrimaryMuscles[0] != "Back" || exercise.PrimaryMuscles[1] != "Biceps" {
			t.Errorf("NewExercise() PrimaryMuscles = %v, expected %v", exercise.PrimaryMuscles, []string{"Back", "Biceps"})
		}
		if len(exercise.SecondaryMuscles) != 1 || exercise.SecondaryMuscles[0] != "Forearms" {
			t.Errorf("NewExercise() SecondaryMuscles = %v, expected %v", exercise.SecondaryMuscles, []string{"Forearms"})
		}
		if exercise.Category != "Bodyweight" {
			t.Errorf("NewExercise() Category = %q, expected %q", exercise.Category, "Bodyweight")
		}
		if len(exercise.Images) != 0 {
			t.Errorf("NewExercise() Images length = %d, expected %d", len(exercise.Images), 0)
		}
	})

	t.Run("create exercise with empty arrays", func(t *testing.T) {
		exercise := NewExercise(
			"exercise-000",
			"test exercise",
			stringPtr("test"),
			"test",
			stringPtr("test"),
			stringPtr("test"),
			[]string{},
			[]string{},
			[]string{},
			"test",
			[]Image{},
		)

		if len(exercise.PrimaryMuscles) != 0 {
			t.Errorf("NewExercise() PrimaryMuscles length = %d, expected %d", len(exercise.PrimaryMuscles), 0)
		}
		if len(exercise.SecondaryMuscles) != 0 {
			t.Errorf("NewExercise() SecondaryMuscles length = %d, expected %d", len(exercise.SecondaryMuscles), 0)
		}
		if len(exercise.Instructions) != 0 {
			t.Errorf("NewExercise() Instructions length = %d, expected %d", len(exercise.Instructions), 0)
		}
		if len(exercise.Images) != 0 {
			t.Errorf("NewExercise() Images length = %d, expected %d", len(exercise.Images), 0)
		}
	})
}

func TestJsonExerciseUnmarshal(t *testing.T) {
	t.Run("unmarshal valid json", func(t *testing.T) {
		jsonStr := `{
			"id": "exercise-123",
			"name": "Bench Press",
			"force": "push",
			"level": "beginner",
			"mechanic": "compound",
			"equipment": "barbell",
			"primaryMuscles": ["chest", "triceps"],
			"secondaryMuscles": ["shoulders"],
			"instructions": ["Lie on a flat bench.", "Grip the barbell."],
			"category": "strength",
			"images": ["bench1.jpg", "bench2.jpg"]
		}`

		var exercise jsonExercise
		err := json.Unmarshal([]byte(jsonStr), &exercise)
		if err != nil {
			t.Errorf("json.Unmarshal() error = %v", err)
		}
		if exercise.SourceID != "exercise-123" {
			t.Errorf("jsonExercise.SourceID = %q, expected %q", exercise.SourceID, "exercise-123")
		}
		if exercise.Name != "Bench Press" {
			t.Errorf("jsonExercise.Name = %q, expected %q", exercise.Name, "Bench Press")
		}
		if *exercise.Force != "push" {
			t.Errorf("jsonExercise.Force = %q, expected %q", *exercise.Force, "push")
		}
		if exercise.Level != "beginner" {
			t.Errorf("jsonExercise.Level = %q, expected %q", exercise.Level, "beginner")
		}
		if *exercise.Mechanic != "compound" {
			t.Errorf("jsonExercise.Mechanic = %q, expected %q", *exercise.Mechanic, "compound")
		}
		if *exercise.Equipment != "barbell" {
			t.Errorf("jsonExercise.Equipment = %q, expected %q", *exercise.Equipment, "barbell")
		}
		if len(exercise.PrimaryMuscles) != 2 || exercise.PrimaryMuscles[0] != "chest" || exercise.PrimaryMuscles[1] != "triceps" {
			t.Errorf("jsonExercise.PrimaryMuscles = %v, expected %v", exercise.PrimaryMuscles, []string{"chest", "triceps"})
		}
		if len(exercise.SecondaryMuscles) != 1 || exercise.SecondaryMuscles[0] != "shoulders" {
			t.Errorf("jsonExercise.SecondaryMuscles = %v, expected %v", exercise.SecondaryMuscles, []string{"shoulders"})
		}
		if len(exercise.Instructions) != 2 || exercise.Instructions[0] != "Lie on a flat bench." || exercise.Instructions[1] != "Grip the barbell." {
			t.Errorf("jsonExercise.Instructions = %v, expected specific instructions", exercise.Instructions)
		}
		if exercise.Category != "strength" {
			t.Errorf("jsonExercise.Category = %q, expected %q", exercise.Category, "strength")
		}
		if len(exercise.Images) != 2 || exercise.Images[0] != "bench1.jpg" || exercise.Images[1] != "bench2.jpg" {
			t.Errorf("jsonExercise.Images = %v, expected %v", exercise.Images, []string{"bench1.jpg", "bench2.jpg"})
		}
	})

	t.Run("unmarshal json with null fields", func(t *testing.T) {
		jsonStr := `{
			"id": "exercise-456",
			"name": "Pull-up",
			"force": null,
			"level": "intermediate",
			"mechanic": null,
			"equipment": null,
			"primaryMuscles": ["back", "biceps"],
			"secondaryMuscles": [],
			"instructions": ["Grip the bar.", "Pull yourself up."],
			"category": "bodyweight",
			"images": []
		}`

		var exercise jsonExercise
		err := json.Unmarshal([]byte(jsonStr), &exercise)
		if err != nil {
			t.Errorf("json.Unmarshal() error = %v", err)
		}
		if exercise.SourceID != "exercise-456" {
			t.Errorf("jsonExercise.SourceID = %q, expected %q", exercise.SourceID, "exercise-456")
		}
		if exercise.Name != "Pull-up" {
			t.Errorf("jsonExercise.Name = %q, expected %q", exercise.Name, "Pull-up")
		}
		if exercise.Force != nil {
			t.Errorf("jsonExercise.Force should be nil, got %v", exercise.Force)
		}
		if exercise.Level != "intermediate" {
			t.Errorf("jsonExercise.Level = %q, expected %q", exercise.Level, "intermediate")
		}
		if exercise.Mechanic != nil {
			t.Errorf("jsonExercise.Mechanic should be nil, got %v", exercise.Mechanic)
		}
		if exercise.Equipment != nil {
			t.Errorf("jsonExercise.Equipment should be nil, got %v", exercise.Equipment)
		}
		if len(exercise.PrimaryMuscles) != 2 || exercise.PrimaryMuscles[0] != "back" || exercise.PrimaryMuscles[1] != "biceps" {
			t.Errorf("jsonExercise.PrimaryMuscles = %v, expected %v", exercise.PrimaryMuscles, []string{"back", "biceps"})
		}
		if len(exercise.SecondaryMuscles) != 0 {
			t.Errorf("jsonExercise.SecondaryMuscles length = %d, expected %d", len(exercise.SecondaryMuscles), 0)
		}
		if exercise.Category != "bodyweight" {
			t.Errorf("jsonExercise.Category = %q, expected %q", exercise.Category, "bodyweight")
		}
		if len(exercise.Images) != 0 {
			t.Errorf("jsonExercise.Images length = %d, expected %d", len(exercise.Images), 0)
		}
	})
}
