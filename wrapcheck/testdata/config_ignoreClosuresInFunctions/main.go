package main

import (
	"os"
)

func main() {
	do()
}

func do() error {
	if err := closure(func() error {
		if _, err := os.Open("nonexistent"); err != nil {
			return err // This should be ignored due to ignoreClosuresInFunctions: "closure("
		}
		return nil
	}); err != nil {
		return err
	}

	if err := transaction(func() error {
		if _, err := os.Open("nonexistent"); err != nil {
			return err // want "error returned from external package is unwrapped"
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}

func closure(fn func() error) error {
	return fn()
}

func transaction(fn func() error) error {
	return fn()
}
