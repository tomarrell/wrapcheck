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
	// Internal function
	if err := fn(); err != nil {
		return err // want `package-internal error should be wrapped`
	}
	if err := fn(); err != nil {
		return fmt.Errorf("wrap: %w", err)
	}

	// Internal struct method
	ss := someStruct{}
	if err := ss.someMethod(); err != nil {
		return err // want `package-internal error should be wrapped`
	}
	if err := ss.someMethod(); err != nil {
		return fmt.Errorf("wrap: %w", err)
	}

	// Interface method
	var si someInterface = &someInterfaceImpl{}
	if err := si.SomeMethod(); err != nil {
		return err // want `error returned from interface method should be wrapped`
	}
	if err := si.SomeMethod(); err != nil {
		return fmt.Errorf("wrap: %w", err)
	}

	// External function
	if _, err := json.Marshal(struct{}{}); err != nil {
		return err // want `error returned from external package is unwrapped`
	}
	if _, err := json.Marshal(struct{}{}); err != nil {
		return fmt.Errorf("wrap: %w", err)
	}

	// Ignore sigs
	if err := wrap(errors.New("error")); err != nil {
		return err
	}

	// Extra ignore sigs
	if err := wrapError(errors.New("error")); err != nil {
		return err
	}

	// Ignore sig regexps
	if err := newError(errors.New("error")); err != nil {
		return err
	}

	return nil
}

func fn() error {
	return errors.New("error")
}

type someStruct struct{}

func (s *someStruct) someMethod() error {
	return errors.New("error")
}

type someInterface interface {
	SomeMethod() error
}

type someInterfaceImpl struct {
	someInterface
}

func wrap(err error) error {
	return fmt.Errorf("wrap: %w", err)
}

func wrapError(err error) error {
	return fmt.Errorf("wrap: %w", err)
}

func newError(err error) error {
	return fmt.Errorf("new error: %w", err)
}
