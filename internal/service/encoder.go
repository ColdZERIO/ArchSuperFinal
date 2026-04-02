package services

import (
	"Go/internal/models"
	"encoding/json"
	"os"
)

func decoderJson(filename string, tasks []models.Task) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	enc := json.NewEncoder(file)
	enc.SetIndent(" ", "	")

	err = enc.Encode(tasks)
	return err
}