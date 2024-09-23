package utils

import (
	"encoding/json"
	"os"
)

func ReadJSONFile(filePath string) map[string]string {
	defaultData := make(map[string]string)
	data, err := os.ReadFile(filePath)
	if err != nil {
		return defaultData
	}

	var result map[string]string
	err = json.Unmarshal(data, &result)
	if err != nil {
		return defaultData
	}
	return result
}
