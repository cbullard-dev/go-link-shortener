package storage

import (
	"encoding/json"
	"os"
)

func LoadData(fileName string, urlMap map[string]string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}

	defer file.Close()

	decoder := json.NewDecoder(file)
	return decoder.Decode(&urlMap)
}

func SaveData(fileName string, urlMap map[string]string) error {
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}

	defer file.Close()

	encoder := json.NewEncoder(file)
	return encoder.Encode(urlMap)
}
