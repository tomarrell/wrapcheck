package main

import (
	"encoding/json"
	"errors"
)

func main() {
	do()
}

func do() error {
	// no issue with function in 'extraIgnoreSigs'
	_, err := json.Marshal(struct{}{})
	if err != nil {
		return err
	}

	// expect issue for function that is not ignored
	res := struct{}{}
	if err := json.Unmarshal([]byte("{}"), &res); err != nil {
		return err // want `error returned from external package is unwrapped`
	}

	// no issue with function in 'ignoreSigs'
	return errors.New("Some error")
}
