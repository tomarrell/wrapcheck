package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

func main() {
	do()
}

func do() error {
	_, err := json.Marshal(struct{}{})
	if err != nil {
		return fmt.Errorf("failed to marshal struct", err)
	}

	// This should not error, as it's a call into a local function which returns
	// an error
	err = returnsError()
	if err != nil {
		return err
	}

	// This should also not error, as it's a call into the errors pkg
	err = errors.New("failed")
	if err != nil {
		return err
	}

	// This should error
	_, err = json.Marshal(struct{}{})
	if err != nil {
		return err // want `error returned from external package is unwrapped`
	}

	return nil
}

func returnsError() error {
	return errors.New("failed")
}
