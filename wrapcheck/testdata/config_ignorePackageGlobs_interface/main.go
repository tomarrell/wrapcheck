package main

import (
	"encoding/json"
	"time"
)

func main() {
	do(&time.Time{})
}

func do(fn json.Unmarshaler) error {
	err := fn.UnmarshalJSON([]byte{})
	if err != nil {
		return err
	}

	return nil
}
