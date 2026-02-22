package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
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
		return Exercise{}, err
	}

	var metadata jsonExercise
	if err := json.Unmarshal(data, &metadata); err != nil {
		return Exercise{}, err
	}

	baseDir := filepath.Dir(path)
	images := make([]Image, 0, len(metadata.Images))
	for _, imgPath := range metadata.Images {
		absPath := filepath.Join(baseDir, imgPath)
		imgData, err := os.ReadFile(absPath)
		if err != nil {
			return Exercise{}, err
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
