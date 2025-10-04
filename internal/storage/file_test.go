package storage

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestSaveData(t *testing.T) {
	saveFileName := filepath.Join(t.TempDir(), "testSaveData.json")
	testMap := map[string]string{"unique_key_one": "string_value_one", "unique_key_two": "string_value_two"}

	SaveData(saveFileName, testMap)

	_, err := os.Stat(saveFileName)

	if os.IsNotExist(err) {
		t.Errorf(`The save file does not exist: %q`, err)
	}

	data, err := os.ReadFile(saveFileName)

	if err != nil {
		t.Errorf(`Error reading the file: %q`, err)
	}
	loadedData := make(map[string]string)
	json.Unmarshal([]byte(data), &loadedData)
	if !reflect.DeepEqual(loadedData, testMap) {
		t.Errorf(`Input data: %q, does not match the output data: %q`, testMap, loadedData)
	}
}

func TestLoadData(t *testing.T) {
	loadFileName := "../../testdata/testDatabaseLoad.json"
	testMap := map[string]string{"unique_key_one": "string_value_one", "unique_key_two": "string_value_two"}
	loadMap := make(map[string]string)
	err := LoadData(loadFileName, loadMap)
	if err != nil {
		t.Errorf(`Loading the data returned error: %q`, err)
	}

	if !reflect.DeepEqual(testMap, loadMap) {
		t.Errorf(`Test data: %q, does not match the loaded data: %q`, testMap, loadMap)
	}
}
