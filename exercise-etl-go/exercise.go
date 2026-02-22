package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type jsonExercise struct {
	SourceID         string   `json:"id"`
	Name             string   `json:"name"`
	Force            *string  `json:"force"`
	Level            string   `json:"level"`
	Mechanic         *string  `json:"mechanic"`
	Equipment        *string  `json:"equipment"`
	PrimaryMuscles   []string `json:"primaryMuscles"`
	SecondaryMuscles []string `json:"secondaryMuscles"`
	Instructions     []string `json:"instructions"`
	Category         string   `json:"category"`
	Images           []string `json:"images"`
}

type Image struct {
	ImageBlob []byte
	MimeType  string
}

func validateMetadata(metadata jsonExercise) (jsonExercise, error) {
	metadata.SourceID = NormalizeString(metadata.SourceID)
	if metadata.SourceID == "" {
		return metadata, fmt.Errorf("source id is empty")
	}
	metadata.Name = NormalizeString(metadata.Name)
	if metadata.Name == "" {
		return metadata, fmt.Errorf("name is empty")
	}
	metadata.Level = NormalizeString(metadata.Level)
	if metadata.Level == "" {
		return metadata, fmt.Errorf("level is empty")
	}
	if metadata.Force != nil {
		metadata.Force = NormalizeStringPtr(metadata.Force)
		if *metadata.Force == "" {
			return metadata, fmt.Errorf("force is empty string")
		}
	}
	if metadata.Mechanic != nil {
		metadata.Mechanic = NormalizeStringPtr(metadata.Mechanic)
		if *metadata.Mechanic == "" {
			return metadata, fmt.Errorf("mechanic is empty string")
		}
	}
	if metadata.Equipment != nil {
		metadata.Equipment = NormalizeStringPtr(metadata.Equipment)
		if *metadata.Equipment == "" {
			return metadata, fmt.Errorf("equipment is empty string")
		}
	}
	metadata.Category = NormalizeString(metadata.Category)
	if metadata.Category == "" {
		return metadata, fmt.Errorf("category is empty")
	}
	for i := range metadata.PrimaryMuscles {
		metadata.PrimaryMuscles[i] = NormalizeString(metadata.PrimaryMuscles[i])
		if metadata.PrimaryMuscles[i] == "" {
			return metadata, fmt.Errorf("primary muscles[%d] is empty", i)
		}
	}
	for i := range metadata.SecondaryMuscles {
		metadata.SecondaryMuscles[i] = NormalizeString(metadata.SecondaryMuscles[i])
		if metadata.SecondaryMuscles[i] == "" {
			return metadata, fmt.Errorf("secondary muscles[%d] is empty", i)
		}
	}

	filteredInstructions := make([]string, 0, len(metadata.Instructions))
	for _, instr := range metadata.Instructions {
		normalized := NormalizeString(instr)
		if normalized != "" {
			filteredInstructions = append(filteredInstructions, normalized)
		}
	}
	metadata.Instructions = filteredInstructions

	for i := range metadata.Images {
		metadata.Images[i] = strings.TrimSpace(metadata.Images[i])
		if metadata.Images[i] == "" {
			return metadata, fmt.Errorf("images[%d] path is empty", i)
		}
	}
	return metadata, nil
}

type Exercise struct {
	Id               int
	SourceId         string
	Name             string
	Force            *string
	Level            string
	Mechanic         *string
	Equipment        *string
	PrimaryMuscles   []string
	SecondaryMuscles []string
	Instructions     []string
	Category         string
	Images           []Image
}

func NewExercise(
	sourceId string,
	name string,
	force *string,
	level string,
	mechanic *string,
	equipment *string,
	primaryMuscles []string,
	secondaryMuscles []string,
	instructions []string,
	category string,
	images []Image,
) Exercise {
	return Exercise{
		SourceId:         sourceId,
		Name:             NormalizeString(name),
		Force:            NormalizeStringPtr(force),
		Level:            NormalizeString(level),
		Mechanic:         NormalizeStringPtr(mechanic),
		Equipment:        NormalizeStringPtr(equipment),
		PrimaryMuscles:   Map(primaryMuscles, NormalizeString),
		SecondaryMuscles: Map(secondaryMuscles, NormalizeString),
		Instructions:     instructions,
		Category:         NormalizeString(category),
		Images:           images,
	}
}

func NewExerciseFromMetadata(path string) (Exercise, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return Exercise{}, fmt.Errorf("read file %s: %w", path, err)
	}

	var metadata jsonExercise
	if err := json.Unmarshal(data, &metadata); err != nil {
		return Exercise{}, fmt.Errorf("parse json %s: %w", path, err)
	}

	validatedMetadata, err := validateMetadata(metadata)
	if err != nil {
		return Exercise{}, fmt.Errorf("validate metadata %s: %w", path, err)
	}
	metadata = validatedMetadata

	baseDir := filepath.Dir(path)
	images := make([]Image, 0, len(metadata.Images))
	for _, imgPath := range metadata.Images {
		absPath := filepath.Join(baseDir, imgPath)
		imgData, err := os.ReadFile(absPath)
		if err != nil {
			return Exercise{}, fmt.Errorf("read image %s for %s: %w", absPath, path, err)
		}
		mimeType := http.DetectContentType(imgData)
		images = append(images, Image{
			ImageBlob: imgData,
			MimeType:  mimeType,
		})
	}

	return NewExercise(
		metadata.SourceID,
		metadata.Name,
		metadata.Force,
		metadata.Level,
		metadata.Mechanic,
		metadata.Equipment,
		metadata.PrimaryMuscles,
		metadata.SecondaryMuscles,
		metadata.Instructions,
		metadata.Category,
		images,
	), nil
}

func NewExerciseFromFolder(path string) ([]Exercise, error) {
	var exercises []Exercise

	info, err := os.Stat(path)
	if err != nil {
		return nil, err
	}
	if !info.IsDir() {
		return nil, fmt.Errorf("%s is not a directory", path)
	}

	matches, err := filepath.Glob(filepath.Join(path, "*.json"))
	if err != nil {
		return nil, err
	}

	for _, filePath := range matches {
		exercise, err := NewExerciseFromMetadata(filePath)
		if err != nil {
			return nil, err
		}

		exercises = append(exercises, exercise)
	}

	return exercises, nil
}
