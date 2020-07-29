package main

import "errors"

func main() {
	do()
}

func do() error {
	err := errors.New("failed to do something")
	if err != nil {
		return err
	}

	return nil
}
