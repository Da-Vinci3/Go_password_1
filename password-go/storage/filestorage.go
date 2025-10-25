package storage

import (
	"encoding/json"
	"fmt"
	"os"
)

func ReadPasswordFile(filename string) ([]map[string]string, error) {
	var passwords []map[string]string
	content, err := os.ReadFile(filename)
	if err != nil && !os.IsNotExist(err) {
		return nil, fmt.Errorf("error reading file: %v", err)
	}

	if len(content) > 0 {
		err = json.Unmarshal(content, &passwords)
		if err != nil {
			return nil, fmt.Errorf("error parsing JSON: %v", err)
		}
	}

	return passwords, nil
}

func SavePasswords(filename string, passwords []map[string]string) error {
	updatedJSON, err := json.MarshalIndent(passwords, "", "    ")
	if err != nil {
		return fmt.Errorf("error creating JSON: %v", err)
	}

	err = os.WriteFile(filename, updatedJSON, 0644)
	if err != nil {
		return fmt.Errorf("error writing to file: %v", err)
	}

	return nil
}
