package jsonstorage

import (
	"encoding/json"
	"os"
)

func ReadData(file string) ([]byte, error) {
	return os.ReadFile(file) // should opt out of using ReadFile to read file in faster
}

func LoadJson(data *[]byte, template interface{}) error {
	return json.Unmarshal(*data, template)
}
