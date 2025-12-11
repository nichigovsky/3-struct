package storage

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Storage struct {}

func ReadFile(name string) (*Storage, error) {
	file, err := os.ReadFile(name + ".json")
	if err != nil {
		return nil, errors.New("COULD_NOT_READ_FILE")
	}

	var storage Storage
	err = json.Unmarshal(file, &storage)

	if err != nil {
		return nil, errors.New("COULD_NOT_READ_JSON")
	}

	return &storage, nil
}

func WriteFile(storage *Storage, name string) error {
	file, err := os.Create(name + ".json")

	if err != nil {
		return errors.New("COULD_NOT_CREATE_FILE")
	}

	data, err := json.Marshal(storage)

	if err != nil {
		return errors.New("COULD_NOT_CREATE_JSON")
	}

	_, err = file.Write(data)
	defer file.Close()
	
	if err != nil {
		return errors.New("COULD_NOT_WRITE_FILE")
	}

	fmt.Println("JSON file has been created")

	return nil
}