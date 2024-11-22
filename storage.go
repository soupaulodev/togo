package main

import (
	"encoding/json"
	"os"
)

// Represents a storage instance
type Storage[T any] struct {
	FileName string
}

// NewStorage creates a new storage instance
func NewStorage[T any](fileName string) *Storage[T] {
	return &Storage[T]{
		FileName: fileName,
	}
}

// Save marshals the data into JSON and writes it to the file
func (s *Storage[T]) Save(data T) error {
	fileData, err := json.MarshalIndent(data, "", "    ")

	if err != nil {
		return err
	}

	return os.WriteFile(s.FileName, fileData, 0644)
}

// Load reads the data from the file and unmarshals it into the provided data structure
func (s *Storage[T]) Load(data *T) error {
	fileData, err := os.ReadFile(s.FileName)

	if err != nil {
		return err
	}
	
	return json.Unmarshal(fileData, data)
}

