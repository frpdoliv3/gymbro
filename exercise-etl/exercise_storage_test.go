package main

import (
	"context"
	"database/sql"
	"errors"
	"regexp"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestStorage_HealthCheck(t *testing.T) {
	t.Run("nil database", func(t *testing.T) {
		s := &Storage{db: nil}
		err := s.HealthCheck(context.Background())
		if err == nil {
			t.Error("HealthCheck() expected error for nil database")
		}
		if err.Error() != "database not connected" {
			t.Errorf("HealthCheck() error = %v, expected %q", err, "database not connected")
		}
	})

	t.Run("database ping error", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.MonitorPingsOption(true))
		if err != nil {
			t.Fatalf("failed to create sqlmock: %v", err)
		}
		defer db.Close()

		mock.ExpectPing().WillReturnError(errors.New("connection failed"))

		s := &Storage{db: db}
		err = s.HealthCheck(context.Background())
		if err == nil {
			t.Error("HealthCheck() expected error for ping failure")
		}
		if err.Error() != "connection failed" {
			t.Errorf("HealthCheck() error = %v, expected %q", err, "connection failed")
		}
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("unfulfilled expectations: %v", err)
		}
	})

	t.Run("successful health check", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.MonitorPingsOption(true))
		if err != nil {
			t.Fatalf("failed to create sqlmock: %v", err)
		}
		defer db.Close()

		mock.ExpectPing()

		s := &Storage{db: db}
		err = s.HealthCheck(context.Background())
		if err != nil {
			t.Errorf("HealthCheck() unexpected error: %v", err)
		}
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("unfulfilled expectations: %v", err)
		}
	})
}

func TestStorage_Close(t *testing.T) {
	t.Run("nil database", func(t *testing.T) {
		s := &Storage{db: nil}
		err := s.Close()
		if err != nil {
			t.Errorf("Close() error = %v, expected nil", err)
		}
	})

	t.Run("database close error", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("failed to create sqlmock: %v", err)
		}

		closeErr := errors.New("close failed")
		mock.ExpectClose().WillReturnError(closeErr)

		s := &Storage{db: db}
		err = s.Close()
		if err == nil {
			t.Error("Close() expected error for close failure")
		}
		if err != closeErr {
			t.Errorf("Close() error = %v, expected %v", err, closeErr)
		}
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("unfulfilled expectations: %v", err)
		}
	})

	t.Run("successful close", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("failed to create sqlmock: %v", err)
		}

		mock.ExpectClose()

		s := &Storage{db: db}
		err = s.Close()
		if err != nil {
			t.Errorf("Close() unexpected error: %v", err)
		}
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("unfulfilled expectations: %v", err)
		}
	})
}

func TestStorage_WithTransaction(t *testing.T) {
	ctx := context.Background()

	t.Run("begin transaction error", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("failed to create sqlmock: %v", err)
		}
		defer db.Close()

		mock.ExpectBegin().WillReturnError(errors.New("begin tx failed"))

		s := &Storage{db: db}

		err = s.WithTransaction(ctx, func(tx *sql.Tx) error {
			return nil
		})

		if err == nil {
			t.Error("WithTransaction() expected error for begin transaction failure")
		}
		if err.Error() != "begin transaction: begin tx failed" {
			t.Errorf("WithTransaction() error = %v, expected %q", err, "begin transaction: begin tx failed")
		}
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("unfulfilled expectations: %v", err)
		}
	})

	t.Run("function returns error with rollback failure", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("failed to create sqlmock: %v", err)
		}
		defer db.Close()

		mock.ExpectBegin()
		mock.ExpectRollback().WillReturnError(errors.New("rollback failed"))

		s := &Storage{db: db}

		funcErr := errors.New("function failed")
		err = s.WithTransaction(ctx, func(tx *sql.Tx) error {
			return funcErr
		})

		if err == nil {
			t.Error("WithTransaction() expected error")
		}
		// Check that error contains both function error and rollback error
		if !errors.Is(err, funcErr) {
			t.Errorf("WithTransaction() error should contain %v, got %v", funcErr, err)
		}
		// Check that both error messages are present (order may vary with errors.Join)
		errStr := err.Error()
		if !strings.Contains(errStr, "function failed") || !strings.Contains(errStr, "rollback failed") {
			t.Errorf("WithTransaction() error = %v, expected to contain both 'function failed' and 'rollback failed'", err)
		}
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("unfulfilled expectations: %v", err)
		}
	})

	t.Run("function returns error with successful rollback", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("failed to create sqlmock: %v", err)
		}
		defer db.Close()

		mock.ExpectBegin()
		mock.ExpectRollback()

		s := &Storage{db: db}

		funcErr := errors.New("function failed")
		err = s.WithTransaction(ctx, func(tx *sql.Tx) error {
			return funcErr
		})

		if err == nil {
			t.Error("WithTransaction() expected error")
		}
		if !errors.Is(err, funcErr) {
			t.Errorf("WithTransaction() error = %v, expected %v", err, funcErr)
		}
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("unfulfilled expectations: %v", err)
		}
	})

	t.Run("commit error", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("failed to create sqlmock: %v", err)
		}
		defer db.Close()

		mock.ExpectBegin()
		mock.ExpectCommit().WillReturnError(errors.New("commit failed"))

		s := &Storage{db: db}

		err = s.WithTransaction(ctx, func(tx *sql.Tx) error {
			return nil
		})

		if err == nil {
			t.Error("WithTransaction() expected error for commit failure")
		}
		if err.Error() != "commit transaction: commit failed" {
			t.Errorf("WithTransaction() error = %v, expected %q", err, "commit transaction: commit failed")
		}
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("unfulfilled expectations: %v", err)
		}
	})

	t.Run("successful transaction", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("failed to create sqlmock: %v", err)
		}
		defer db.Close()

		mock.ExpectBegin()
		mock.ExpectCommit()

		s := &Storage{db: db}

		called := false
		err = s.WithTransaction(ctx, func(tx *sql.Tx) error {
			called = true
			return nil
		})

		if err != nil {
			t.Errorf("WithTransaction() unexpected error: %v", err)
		}
		if !called {
			t.Error("WithTransaction() transaction function was not called")
		}
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("unfulfilled expectations: %v", err)
		}
	})
}

func TestStorage_StoreExercise(t *testing.T) {
	ctx := context.Background()

	// Create a minimal exercise for testing
	exercise := Exercise{
		SourceId:         "test-123",
		Name:             "Test Exercise",
		Force:            stringPtr("Push"),
		Level:            "Beginner",
		Mechanic:         stringPtr("Compound"),
		Equipment:        stringPtr("Barbell"),
		PrimaryMuscles:   []string{"Chest", "Triceps"},
		SecondaryMuscles: []string{"Shoulders"},
		Instructions:     []string{"Step 1", "Step 2"},
		Category:         "Strength",
		Images:           []Image{},
	}

	t.Run("transaction error propagates", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("failed to create sqlmock: %v", err)
		}
		defer db.Close()

		mock.ExpectBegin().WillReturnError(errors.New("begin failed"))

		s := &Storage{db: db}

		_, err = s.StoreExercise(ctx, exercise)
		if err == nil {
			t.Error("StoreExercise() expected error for transaction failure")
		}
		if err.Error() != "begin transaction: begin failed" {
			t.Errorf("StoreExercise() error = %v, expected %q", err, "begin transaction: begin failed")
		}
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("unfulfilled expectations: %v", err)
		}
	})

	t.Run("successful store returns ID", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("failed to create sqlmock: %v", err)
		}
		defer db.Close()

		// Begin transaction
		mock.ExpectBegin()

		// Insert exercise
		mock.ExpectExec(regexp.QuoteMeta(`
		INSERT INTO exercises (source_id, name, force, level, mechanic, equipment)
		VALUES (?, ?, ?, ?, ?, ?)
	`)).WithArgs(
			"test-123",
			"Test Exercise",
			"Push",
			"Beginner",
			"Compound",
			"Barbell",
		).WillReturnResult(sqlmock.NewResult(1, 1))

		// Category exists check
		mock.ExpectQuery(regexp.QuoteMeta(`SELECT id FROM categories WHERE name = ?`)).
			WithArgs("Strength").
			WillReturnError(sql.ErrNoRows) // Category doesn't exist

		// Insert category
		mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO categories (name) VALUES (?)`)).
			WithArgs("Strength").
			WillReturnResult(sqlmock.NewResult(10, 1))

		// Link exercise to category
		mock.ExpectExec(regexp.QuoteMeta(`
			INSERT INTO exercise_categories (exercise_id, category_id)
			VALUES (?, ?)
		`)).WithArgs(int64(1), int64(10)).
			WillReturnResult(sqlmock.NewResult(0, 1))

		// Primary muscles - Chest
		mock.ExpectQuery(regexp.QuoteMeta(`SELECT id FROM muscles WHERE name = ?`)).
			WithArgs("Chest").
			WillReturnError(sql.ErrNoRows)

		mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO muscles (name) VALUES (?)`)).
			WithArgs("Chest").
			WillReturnResult(sqlmock.NewResult(20, 1))

		mock.ExpectExec(regexp.QuoteMeta(`
			INSERT INTO exercise_muscles (exercise_id, muscle_id, muscle_type)
			VALUES (?, ?, ?)
		`)).WithArgs(int64(1), int64(20), "primary").
			WillReturnResult(sqlmock.NewResult(0, 1))

		// Primary muscles - Triceps
		mock.ExpectQuery(regexp.QuoteMeta(`SELECT id FROM muscles WHERE name = ?`)).
			WithArgs("Triceps").
			WillReturnError(sql.ErrNoRows)

		mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO muscles (name) VALUES (?)`)).
			WithArgs("Triceps").
			WillReturnResult(sqlmock.NewResult(21, 1))

		mock.ExpectExec(regexp.QuoteMeta(`
			INSERT INTO exercise_muscles (exercise_id, muscle_id, muscle_type)
			VALUES (?, ?, ?)
		`)).WithArgs(int64(1), int64(21), "primary").
			WillReturnResult(sqlmock.NewResult(0, 1))

		// Secondary muscles - Shoulders
		mock.ExpectQuery(regexp.QuoteMeta(`SELECT id FROM muscles WHERE name = ?`)).
			WithArgs("Shoulders").
			WillReturnError(sql.ErrNoRows)

		mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO muscles (name) VALUES (?)`)).
			WithArgs("Shoulders").
			WillReturnResult(sqlmock.NewResult(22, 1))

		mock.ExpectExec(regexp.QuoteMeta(`
			INSERT INTO exercise_muscles (exercise_id, muscle_id, muscle_type)
			VALUES (?, ?, ?)
		`)).WithArgs(int64(1), int64(22), "secondary").
			WillReturnResult(sqlmock.NewResult(0, 1))

		// Instructions
		mock.ExpectExec(regexp.QuoteMeta(`
			INSERT INTO steps (exercise_id, description, step_order)
			VALUES (?, ?, ?)
		`)).WithArgs(int64(1), "Step 1", 1).
			WillReturnResult(sqlmock.NewResult(0, 1))

		mock.ExpectExec(regexp.QuoteMeta(`
			INSERT INTO steps (exercise_id, description, step_order)
			VALUES (?, ?, ?)
		`)).WithArgs(int64(1), "Step 2", 2).
			WillReturnResult(sqlmock.NewResult(0, 1))

		// Commit transaction
		mock.ExpectCommit()

		s := &Storage{db: db}

		id, err := s.StoreExercise(ctx, exercise)
		if err != nil {
			t.Errorf("StoreExercise() unexpected error: %v", err)
		}
		if id != 1 {
			t.Errorf("StoreExercise() returned ID = %d, expected %d", id, 1)
		}
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("unfulfilled expectations: %v", err)
		}
	})

	t.Run("error in storeExerciseTx propagates", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("failed to create sqlmock: %v", err)
		}
		defer db.Close()

		// Begin transaction
		mock.ExpectBegin()

		// Insert exercise fails
		mock.ExpectExec(regexp.QuoteMeta(`
		INSERT INTO exercises (source_id, name, force, level, mechanic, equipment)
		VALUES (?, ?, ?, ?, ?, ?)
	`)).WithArgs(
			"test-123",
			"Test Exercise",
			"Push",
			"Beginner",
			"Compound",
			"Barbell",
		).WillReturnError(errors.New("insert failed"))

		// Rollback
		mock.ExpectRollback()

		s := &Storage{db: db}

		_, err = s.StoreExercise(ctx, exercise)
		if err == nil {
			t.Error("StoreExercise() expected error for insert failure")
		}
		if err.Error() != "insert exercise: insert failed" {
			t.Errorf("StoreExercise() error = %v, expected %q", err, "insert exercise: insert failed")
		}
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("unfulfilled expectations: %v", err)
		}
	})
}

func TestStorage_StoreExercises(t *testing.T) {
	ctx := context.Background()

	exercises := []Exercise{
		{
			SourceId:       "ex1",
			Name:           "Exercise 1",
			Level:          "Beginner",
			Category:       "Strength",
			PrimaryMuscles: []string{"Chest"},
		},
		{
			SourceId:       "ex2",
			Name:           "Exercise 2",
			Level:          "Intermediate",
			Category:       "Strength",
			PrimaryMuscles: []string{"Back"},
		},
	}

	t.Run("empty exercises slice", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("failed to create sqlmock: %v", err)
		}
		defer db.Close()

		// Begin and commit transaction even for empty slice
		mock.ExpectBegin()
		mock.ExpectCommit()

		s := &Storage{db: db}

		err = s.StoreExercises(ctx, []Exercise{})
		if err != nil {
			t.Errorf("StoreExercises() unexpected error for empty slice: %v", err)
		}
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("unfulfilled expectations: %v", err)
		}
	})

	t.Run("transaction error", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("failed to create sqlmock: %v", err)
		}
		defer db.Close()

		mock.ExpectBegin().WillReturnError(errors.New("begin failed"))

		s := &Storage{db: db}

		err = s.StoreExercises(ctx, exercises)
		if err == nil {
			t.Error("StoreExercises() expected error for transaction failure")
		}
		if err.Error() != "begin transaction: begin failed" {
			t.Errorf("StoreExercises() error = %v, expected %q", err, "begin transaction: begin failed")
		}
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("unfulfilled expectations: %v", err)
		}
	})

	t.Run("storeExerciseTx error for first exercise", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("failed to create sqlmock: %v", err)
		}
		defer db.Close()

		// Begin transaction
		mock.ExpectBegin()

		// First exercise insert fails
		mock.ExpectExec(regexp.QuoteMeta(`
		INSERT INTO exercises (source_id, name, force, level, mechanic, equipment)
		VALUES (?, ?, ?, ?, ?, ?)
	`)).WithArgs(
			"ex1",
			"Exercise 1",
			nil, // force
			"Beginner",
			nil, // mechanic
			nil, // equipment
		).WillReturnError(errors.New("insert failed"))

		// Rollback
		mock.ExpectRollback()

		s := &Storage{db: db}

		err = s.StoreExercises(ctx, exercises)
		if err == nil {
			t.Error("StoreExercises() expected error for insert failure")
		}
		if err.Error() != "store exercise ex1: insert exercise: insert failed" {
			t.Errorf("StoreExercises() error = %v, expected %q", err, "store exercise ex1: insert exercise: insert failed")
		}
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("unfulfilled expectations: %v", err)
		}
	})

	t.Run("successful store of multiple exercises", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("failed to create sqlmock: %v", err)
		}
		defer db.Close()

		// Begin transaction
		mock.ExpectBegin()

		// First exercise
		mock.ExpectExec(regexp.QuoteMeta(`
		INSERT INTO exercises (source_id, name, force, level, mechanic, equipment)
		VALUES (?, ?, ?, ?, ?, ?)
	`)).WithArgs(
			"ex1",
			"Exercise 1",
			nil,
			"Beginner",
			nil,
			nil,
		).WillReturnResult(sqlmock.NewResult(1, 1))

		mock.ExpectQuery(regexp.QuoteMeta(`SELECT id FROM categories WHERE name = ?`)).
			WithArgs("Strength").
			WillReturnError(sql.ErrNoRows)

		mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO categories (name) VALUES (?)`)).
			WithArgs("Strength").
			WillReturnResult(sqlmock.NewResult(10, 1))

		mock.ExpectExec(regexp.QuoteMeta(`
			INSERT INTO exercise_categories (exercise_id, category_id)
			VALUES (?, ?)
		`)).WithArgs(int64(1), int64(10)).
			WillReturnResult(sqlmock.NewResult(0, 1))

		mock.ExpectQuery(regexp.QuoteMeta(`SELECT id FROM muscles WHERE name = ?`)).
			WithArgs("Chest").
			WillReturnError(sql.ErrNoRows)

		mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO muscles (name) VALUES (?)`)).
			WithArgs("Chest").
			WillReturnResult(sqlmock.NewResult(20, 1))

		mock.ExpectExec(regexp.QuoteMeta(`
			INSERT INTO exercise_muscles (exercise_id, muscle_id, muscle_type)
			VALUES (?, ?, ?)
		`)).WithArgs(int64(1), int64(20), "primary").
			WillReturnResult(sqlmock.NewResult(0, 1))

		// Second exercise
		mock.ExpectExec(regexp.QuoteMeta(`
		INSERT INTO exercises (source_id, name, force, level, mechanic, equipment)
		VALUES (?, ?, ?, ?, ?, ?)
	`)).WithArgs(
			"ex2",
			"Exercise 2",
			nil,
			"Intermediate",
			nil,
			nil,
		).WillReturnResult(sqlmock.NewResult(2, 1))

		// Category already exists
		mock.ExpectQuery(regexp.QuoteMeta(`SELECT id FROM categories WHERE name = ?`)).
			WithArgs("Strength").
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(10))

		mock.ExpectExec(regexp.QuoteMeta(`
			INSERT INTO exercise_categories (exercise_id, category_id)
			VALUES (?, ?)
		`)).WithArgs(int64(2), int64(10)).
			WillReturnResult(sqlmock.NewResult(0, 1))

		mock.ExpectQuery(regexp.QuoteMeta(`SELECT id FROM muscles WHERE name = ?`)).
			WithArgs("Back").
			WillReturnError(sql.ErrNoRows)

		mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO muscles (name) VALUES (?)`)).
			WithArgs("Back").
			WillReturnResult(sqlmock.NewResult(21, 1))

		mock.ExpectExec(regexp.QuoteMeta(`
			INSERT INTO exercise_muscles (exercise_id, muscle_id, muscle_type)
			VALUES (?, ?, ?)
		`)).WithArgs(int64(2), int64(21), "primary").
			WillReturnResult(sqlmock.NewResult(0, 1))

		// Commit transaction
		mock.ExpectCommit()

		s := &Storage{db: db}

		err = s.StoreExercises(ctx, exercises)
		if err != nil {
			t.Errorf("StoreExercises() unexpected error: %v", err)
		}
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("unfulfilled expectations: %v", err)
		}
	})
}

func TestStorage_StoreExercises_WithExistingEntities(t *testing.T) {
	ctx := context.Background()

	exercise := Exercise{
		SourceId:       "ex-existing",
		Name:           "Exercise Existing",
		Level:          "Beginner",
		Category:       "Strength",
		PrimaryMuscles: []string{"Chest"}, // Muscle already exists
	}

	t.Run("category and muscle already exist", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("failed to create sqlmock: %v", err)
		}
		defer db.Close()

		// Begin transaction
		mock.ExpectBegin()

		// Insert exercise
		mock.ExpectExec(regexp.QuoteMeta(`
		INSERT INTO exercises (source_id, name, force, level, mechanic, equipment)
		VALUES (?, ?, ?, ?, ?, ?)
	`)).WithArgs(
			"ex-existing",
			"Exercise Existing",
			nil,
			"Beginner",
			nil,
			nil,
		).WillReturnResult(sqlmock.NewResult(1, 1))

		// Category already exists
		mock.ExpectQuery(regexp.QuoteMeta(`SELECT id FROM categories WHERE name = ?`)).
			WithArgs("Strength").
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(10))

		// Link exercise to existing category
		mock.ExpectExec(regexp.QuoteMeta(`
			INSERT INTO exercise_categories (exercise_id, category_id)
			VALUES (?, ?)
		`)).WithArgs(int64(1), int64(10)).
			WillReturnResult(sqlmock.NewResult(0, 1))

		// Muscle already exists
		mock.ExpectQuery(regexp.QuoteMeta(`SELECT id FROM muscles WHERE name = ?`)).
			WithArgs("Chest").
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(20))

		// Link exercise to existing muscle
		mock.ExpectExec(regexp.QuoteMeta(`
			INSERT INTO exercise_muscles (exercise_id, muscle_id, muscle_type)
			VALUES (?, ?, ?)
		`)).WithArgs(int64(1), int64(20), "primary").
			WillReturnResult(sqlmock.NewResult(0, 1))

		// Commit transaction
		mock.ExpectCommit()

		s := &Storage{db: db}

		err = s.StoreExercises(ctx, []Exercise{exercise})
		if err != nil {
			t.Errorf("StoreExercises() unexpected error: %v", err)
		}
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("unfulfilled expectations: %v", err)
		}
	})
}
