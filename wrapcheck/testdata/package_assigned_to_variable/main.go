package main

import (
	"encoding/json"
	"strings"
)

func main() {
	do(json.NewDecoder(strings.NewReader("hello world")))
}

func do(dec *json.Decoder) error {
	var str string
	if err := dec.Decode(&str); err != nil {
		return err // want `error returned from external package is unwrapped`
	}

	return nil
}
