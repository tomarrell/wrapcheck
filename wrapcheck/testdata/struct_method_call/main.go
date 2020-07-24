package main

import "errors"

func main() {
	do()
}

func do() error {
	ss := someStruct{}
	err := ss.someMethod()
	if err != nil {
		return err
	}

	return nil
}

// Struct with method
type someStruct struct{}

func (s *someStruct) someMethod() error {
	return errors.New("failed")
}
