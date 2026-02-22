package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	var override bool
	var schemaPath string

	flag.BoolVar(&override, "override", false, "Override existing database file")
	flag.BoolVar(&override, "o", false, "Override existing database file")
	flag.StringVar(&schemaPath, "schema", "db/create.sql", "Path to SQL schema file")
	flag.StringVar(&schemaPath, "s", "db/create.sql", "Path to SQL schema file")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [OPTIONS] <exercise-folder> <exercise-db>\n", filepath.Base(os.Args[0]))
		fmt.Fprintln(os.Stderr, "\nExtracts exercises from a folder, transforms into SQL rows and loads them into a SQLite database")
		fmt.Fprintln(os.Stderr, "\nThe exercise folder must come from the free exercise db repository (https://github.com/yuhonas/free-exercise-db)")
		fmt.Fprintln(os.Stderr, "\nOptions:")
		fmt.Fprintln(os.Stderr, "  -o, --override    Override existing database file")
		fmt.Fprintln(os.Stderr, "  -s, --schema      Path to SQL schema file (default: db/create.sql)")
	}

	flag.Parse()

	args := flag.Args()
	if len(args) != 2 {
		flag.Usage()
		os.Exit(1)
	}

	folderPath := args[0]
	dbPath := args[1]

	if info, err := os.Stat(folderPath); err != nil {
		fmt.Fprintf(os.Stderr, "Error: exercise folder \"%s\" does not exist.\n", folderPath)
		os.Exit(1)
	} else if !info.IsDir() {
		fmt.Fprintf(os.Stderr, "Error: exercise path \"%s\" is not a directory\n", folderPath)
		os.Exit(1)
	}

	if info, err := os.Stat(dbPath); err == nil && info.IsDir() {
		fmt.Fprintf(os.Stderr, "Error: database path \"%s\" is a directory, not a file\n", dbPath)
		os.Exit(1)
	}

	if _, err := os.Stat(dbPath); err == nil {
		if !override {
			fmt.Fprintf(os.Stderr, "Error: database file \"%s\" already exists. Use --override flag to delete and recreate.\n", dbPath)
			os.Exit(1)
		}

		if err := os.Remove(dbPath); err != nil {
			fmt.Fprintf(os.Stderr, "Error: failed to delete existing database file: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Deleted existing database file: %s\n", dbPath)
	} else if !os.IsNotExist(err) {
		fmt.Fprintf(os.Stderr, "Error: cannot access database file \"%s\": %v\n", dbPath, err)
		os.Exit(1)
	}

	fmt.Printf("Creating database: %s\n", dbPath)
	storage, err := CreateStorage(dbPath, schemaPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: failed to create database: %v\n", err)
		os.Exit(1)
	}
	defer storage.Close()

	fmt.Printf("Loading exercises from: %s\n", folderPath)
	exercises, err := NewExerciseFromFolder(folderPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: failed to load exercises: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Loaded %d exercises\n", len(exercises))

	ctx := context.Background()
	if err := storage.StoreExercises(ctx, exercises); err != nil {
		fmt.Fprintf(os.Stderr, "Error: failed to store exercises: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Successfully stored all exercises in database")
}
