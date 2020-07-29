package main

import (
	"fmt"
)

func main() {
	do()
}

func do() error {
	_, err := fmt.Scanf("failed to do something")
	if err != nil {
		return err // want `error returned from external package is unwrapped`
	}

	return nil
}
