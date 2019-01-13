package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func ReadConfigFile(fileName string) (map[string]int, error) {
	fmt.Printf("Reading config file %v ...\n", fileName)

	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var result map[string]int
	if err := json.Unmarshal([]byte(bytes), &result); err != nil {
		return nil, err
	}

	return result, nil
}
