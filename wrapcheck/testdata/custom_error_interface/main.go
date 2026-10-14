package main

import "ext"

func main() {}

// Custom error types returned through an error result are reported.

func structVar() error {
	err := ext.Foo()
	return err // want "error returned from external package is unwrapped"
}

func structCall() error {
	return ext.Foo() // want "error returned from external package is unwrapped"
}

func pointerCall() error {
	return ext.Ptr() // want "error returned from external package is unwrapped"
}

func pointerVar() error {
	err := ext.Ptr()
	return err // want "error returned from external package is unwrapped"
}

func tupleCall() (int, error) {
	return ext.Bar() // want "error returned from external package is unwrapped"
}

func tupleVar() (int, error) {
	n, err := ext.Bar()
	return n, err // want "error returned from external package is unwrapped"
}

func embeddedInterface() error {
	return ext.Coded() // want "error returned from external package is unwrapped"
}

func anyResult() any {
	return ext.Foo() // want "error returned from external package is unwrapped"
}

func funcLit() ext.CustomErr {
	f := func() error {
		err := ext.Foo()
		return err // want "error returned from external package is unwrapped"
	}
	_ = f

	return ext.Foo()
}

// Results which a wrapped error cannot be assigned to are not reported.

func concreteCall() ext.CustomErr {
	return ext.Foo()
}

func concreteVar() ext.CustomErr {
	err := ext.Foo()
	return err
}

func concretePointer() *ext.PtrErr {
	return ext.Ptr()
}

func concreteTupleCall() (int, ext.CustomErr) {
	return ext.Bar()
}

func concreteTupleVar() (int, ext.CustomErr) {
	n, err := ext.Bar()
	return n, err
}

func embeddedInterfaceResult() ext.CodedErr {
	return ext.Coded()
}

func funcLitConcrete() error {
	f := func() ext.CustomErr {
		err := ext.Foo()
		return err
	}
	_ = f

	return nil
}
