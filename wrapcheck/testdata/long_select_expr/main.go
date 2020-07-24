package main

import (
	"encoding/json"
	"strings"
)

type thing struct {
	thing2 struct {
		thing3 struct {
			decoder *json.Decoder
		}
	}
}

func main() {
	t := &thing{}
	t.thing2.thing3.decoder = json.NewDecoder(strings.NewReader("hello world"))

	do(t)
}

func do(thing *thing) error {
	var str string
	if err := thing.thing2.thing3.decoder.Decode(&str); err != nil {
		return err // want `error returned from external package is unwrapped`
	}

	return nil
}
