package file

import (
	"errors"
	"os"
	"strings"
)

func ReadFile(name string) ([]byte, error) {
	if !strings.HasSuffix(name, ".json") {
		return nil, errors.New("NOT_JSON_FILE")
	}

	data, err := os.ReadFile(name)
	if err != nil {
		return nil, errors.New("COULD_NOT_READ_FILE")
	}

	return data, nil
}