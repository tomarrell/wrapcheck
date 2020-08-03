package main

import (
	"encoding/json"
	"strings"
)

type errorer interface {
	Decode(v interface{}) error
}

type foo struct {
	bar errorer
}

func main() {
	d := json.NewDecoder(strings.NewReader("hello world"))
	do(foo{d})
}

func do(f foo) error {
	var str string
	err := f.bar.Decode(&str)
	if err != nil {
		return err // want `error returned from interface method should be wrapped`
	}

	return nil
}
