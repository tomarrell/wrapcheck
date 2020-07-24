package main

import (
	"encoding/json"
)

func main() {
	do()
}

func do() (string, error) {
	var t string
	err, _ := json.Unmarshal([]byte("test"), &t), "test"
	if err != nil {
		return "", err // want "error returned from external package is unwrapped"
	}

	return t, nil
}
