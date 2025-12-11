package storage

import (
	"errors"
	"fmt"
	"os"
)

func ReadFile(name string) ([]byte, error) {
	data, err := os.ReadFile(name)
	if err != nil {
		return nil, errors.New("COULD_NOT_READ_FILE")
	}

	return data, nil
}

func WriteFile(content []byte, name string) error {
	file, err := os.Create(name)

	if err != nil {
		return errors.New("COULD_NOT_CREATE_FILE")
	}

	_, err = file.Write(content)
	defer file.Close()
	
	if err != nil {
		return errors.New("COULD_NOT_WRITE_FILE")
	}

	fmt.Println("File has been created")

	return nil
}