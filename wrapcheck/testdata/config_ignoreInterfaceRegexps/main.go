package main

import (
	"encoding/json"
	"strings"
)

type errorer interface {
	Decode(v interface{}) error
}

func main() {
	d := json.NewDecoder(strings.NewReader("hello world"))
	do(d)
}

func do(fn errorer) error {
	var str string
	err := fn.Decode(&str)
	if err != nil {
		return err // errorer interface ignored as per `ignoreInterfaceRegexps`
	}

	return nil
}
