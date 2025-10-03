package helper

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"testing"
)

func TestUrlGenerationLength(t *testing.T) {
	want := 8
	code := GenerateUrlCode(want)
	if len(code) != want {
		t.Errorf(`Generated code expected length: %q, actual length %q`, want, len(code))
	}
}

func TestUrlGenerationFormat(t *testing.T) {
	code := GenerateUrlCode(8)
	want, _ := regexp.Match("[a-zA-Z0-9]{8}", []byte(code))
	if !want {
		t.Errorf(`Generated code expected to match regex "[a-zA-Z0-9]{8}", actual length result: %q`, code)
	}
}

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
