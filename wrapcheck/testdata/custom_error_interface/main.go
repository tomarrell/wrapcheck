package main

import "ext"

func bar() error {
	err := ext.Foo()
	return err // want "error returned from external package is unwrapped"
}
