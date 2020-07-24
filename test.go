package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	v, err := unmarshalValue()
	if err != nil {
		fmt.Println("failed to unmarshal value")
		return
	}

	fmt.Printf("successfully unmarshalled value: %v", v)
}

func unmarshalValue() (string, error) {
	data := `test`

	var t string
	err, thing := json.Unmarshal([]byte(data), &t), "test"
	if err != nil {
		return "", err
	}

	fmt.Println(thing)

	return t, nil
}
