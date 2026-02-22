package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	flag.Usage = func() {
		fmt.Println("Usage:", filepath.Base(os.Args[0]), "<exercise-folder>")
		fmt.Println()
		fmt.Println("Extracts exercises from a folder, transforms into SQL rows and loads them into a SQLite database")
		fmt.Println()
		fmt.Println("The exercise folder must come from the free exercise db repository (https://github.com/yuhonas/free-exercise-db)")
	}
	flag.Parse()

	args := flag.Args()
	if len(args) != 1 {
		flag.Usage()
		os.Exit(1)
	}

	folder_path := args[0]
	info, err := os.Stat(folder_path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: exercise path \"%s\" does not exist.\n", folder_path)
		os.Exit(1)
	}

	if !info.IsDir() {
		fmt.Fprintf(os.Stderr, "Error: exercise path \"%s\" is not a directory", folder_path)
		os.Exit(1)
	}

	// Load all exercises from the folder
	exercises, err := NewExerciseFromFolder(folder_path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: failed to load exercises from %s: %v\n", folder_path, err)
		os.Exit(1)
	}

	fmt.Printf("Loaded %d exercises from %s\n", len(exercises), folder_path)

	// Print each exercise for verification
	for i, exercise := range exercises {
		fmt.Printf("\nExercise %d:\n", i+1)
		fmt.Printf("  SourceID: %s\n", exercise.SourceId)
		fmt.Printf("  Name: %s\n", exercise.Name)
		fmt.Printf("  Level: %s\n", exercise.Level)
		fmt.Printf("  Categories: %s\n", exercise.Category)
		fmt.Printf("  Images: %d images\n", len(exercise.Images))
	}

}
