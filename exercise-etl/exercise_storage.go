package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

type Storage struct {
	db *sql.DB
}

func OpenStorage(dbPath string) (*Storage, error) {
	info, err := os.Stat(dbPath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("database file does not exist: %s", dbPath)
		}
		return nil, fmt.Errorf("check database file: %w", err)
	}
	if !info.Mode().IsRegular() {
		return nil, fmt.Errorf("database path is not a regular file: %s", dbPath)
	}

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, fmt.Errorf("open database: %w", err)
	}

	if _, err := db.Exec(`PRAGMA foreign_keys = ON;`); err != nil {
		db.Close()
		return nil, fmt.Errorf("set database pragmas: %w", err)
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("ping database: %w", err)
	}

	return &Storage{db: db}, nil
}

func CreateStorage(dbPath, schemaPath string) (*Storage, error) {
	if err := os.MkdirAll(filepath.Dir(dbPath), 0755); err != nil {
		return nil, fmt.Errorf("create database directory: %w", err)
	}

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, fmt.Errorf("open database: %w", err)
	}

	schemaSQL, err := os.ReadFile(schemaPath)
	if err != nil {
		db.Close()
		return nil, fmt.Errorf("read schema file %q: %w", schemaPath, err)
	}

	if _, err := db.Exec(string(schemaSQL)); err != nil {
		db.Close()
		return nil, fmt.Errorf("execute schema: %w", err)
	}

	if _, err := db.Exec(`PRAGMA foreign_keys = ON; PRAGMA busy_timeout = 5000;`); err != nil {
		db.Close()
		return nil, fmt.Errorf("set database pragmas: %w", err)
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("ping database: %w", err)
	}

	return &Storage{db: db}, nil
}

func (s *Storage) Close() error {
	if s.db == nil {
		return nil
	}
	return s.db.Close()
}

func (s *Storage) StoreExercise(ctx context.Context, exercise Exercise) (int64, error) {
	var exerciseID int64
	err := s.WithTransaction(ctx, func(tx *sql.Tx) error {
		id, err := s.storeExerciseTx(ctx, tx, exercise)
		if err != nil {
			return err
		}
		exerciseID = id
		return nil
	})
	if err != nil {
		return 0, err
	}
	return exerciseID, nil
}

func (s *Storage) StoreExercises(ctx context.Context, exercises []Exercise) error {
	return s.WithTransaction(ctx, func(tx *sql.Tx) error {
		for _, exercise := range exercises {
			if _, err := s.storeExerciseTx(ctx, tx, exercise); err != nil {
				return fmt.Errorf("store exercise %s: %w", exercise.SourceId, err)
			}
		}
		return nil
	})
}

func (s *Storage) HealthCheck(ctx context.Context) error {
	if s.db == nil {
		return fmt.Errorf("database not connected")
	}
	return s.db.PingContext(ctx)
}

func (s *Storage) WithTransaction(ctx context.Context, fn func(tx *sql.Tx) error) error {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("begin transaction: %w", err)
	}

	if err := fn(tx); err != nil {
		if rbErr := tx.Rollback(); rbErr != nil && rbErr != sql.ErrTxDone {
			return errors.Join(err, fmt.Errorf("rollback failed: %w", rbErr))
		}
		return err
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("commit transaction: %w", err)
	}

	return nil
}

func (s *Storage) getOrCreateMuscle(ctx context.Context, tx *sql.Tx, name string) (int64, error) {

	var id int64
	err := tx.QueryRowContext(ctx, `SELECT id FROM muscles WHERE name = ?`, name).Scan(&id)
	if err == nil {
		return id, nil
	}
	if err != sql.ErrNoRows {
		return 0, fmt.Errorf("query muscle: %w", err)
	}

	result, err := tx.ExecContext(ctx, `INSERT INTO muscles (name) VALUES (?)`, name)
	if err != nil {
		return 0, fmt.Errorf("insert muscle: %w", err)
	}
	return result.LastInsertId()
}

func (s *Storage) getOrCreateCategory(ctx context.Context, tx *sql.Tx, name string) (int64, error) {

	var id int64
	err := tx.QueryRowContext(ctx, `SELECT id FROM categories WHERE name = ?`, name).Scan(&id)
	if err == nil {
		return id, nil
	}
	if err != sql.ErrNoRows {
		return 0, fmt.Errorf("query category: %w", err)
	}

	result, err := tx.ExecContext(ctx, `INSERT INTO categories (name) VALUES (?)`, name)
	if err != nil {
		return 0, fmt.Errorf("insert category: %w", err)
	}
	return result.LastInsertId()
}

func (s *Storage) storeExerciseTx(ctx context.Context, tx *sql.Tx, exercise Exercise) (int64, error) {
	result, err := tx.ExecContext(ctx, `
		INSERT INTO exercises (source_id, name, force, level, mechanic, equipment)
		VALUES (?, ?, ?, ?, ?, ?)
	`,
		exercise.SourceId,
		exercise.Name,
		exercise.Force,
		exercise.Level,
		exercise.Mechanic,
		exercise.Equipment,
	)
	if err != nil {
		return 0, fmt.Errorf("insert exercise: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("get last insert id: %w", err)
	}

	categoryID, err := s.getOrCreateCategory(ctx, tx, exercise.Category)
	if err != nil {
		return 0, fmt.Errorf("insert category: %w", err)
	}
	_, err = tx.ExecContext(ctx, `
		INSERT INTO exercise_categories (exercise_id, category_id)
		VALUES (?, ?)
	`, id, categoryID)
	if err != nil {
		return 0, fmt.Errorf("link exercise to category: %w", err)
	}

	for _, muscle := range exercise.PrimaryMuscles {
		muscleID, err := s.getOrCreateMuscle(ctx, tx, muscle)
		if err != nil {
			return 0, fmt.Errorf("insert primary muscle %s: %w", muscle, err)
		}
		_, err = tx.ExecContext(ctx, `
			INSERT INTO exercise_muscles (exercise_id, muscle_id, muscle_type)
			VALUES (?, ?, ?)
		`, id, muscleID, "primary")
		if err != nil {
			return 0, fmt.Errorf("link exercise to primary muscle: %w", err)
		}
	}

	for _, muscle := range exercise.SecondaryMuscles {
		muscleID, err := s.getOrCreateMuscle(ctx, tx, muscle)
		if err != nil {
			return 0, fmt.Errorf("insert secondary muscle %s: %w", muscle, err)
		}
		_, err = tx.ExecContext(ctx, `
			INSERT INTO exercise_muscles (exercise_id, muscle_id, muscle_type)
			VALUES (?, ?, ?)
		`, id, muscleID, "secondary")
		if err != nil {
			return 0, fmt.Errorf("link exercise to secondary muscle: %w", err)
		}
	}

	for i, instruction := range exercise.Instructions {
		_, err := tx.ExecContext(ctx, `
			INSERT INTO steps (exercise_id, description, step_order)
			VALUES (?, ?, ?)
		`, id, instruction, i+1)
		if err != nil {
			return 0, fmt.Errorf("insert instruction %d: %w", i, err)
		}
	}

	for i, img := range exercise.Images {
		_, err := tx.ExecContext(ctx, `
			INSERT INTO exercise_images (exercise_id, image_order, image_blob, mime_type)
			VALUES (?, ?, ?, ?)
		`, id, i, img.ImageBlob, img.MimeType)
		if err != nil {
			return 0, fmt.Errorf("insert image %d: %w", i, err)
		}
	}

	return id, nil
}
