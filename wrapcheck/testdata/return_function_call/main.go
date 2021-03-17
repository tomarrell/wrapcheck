package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

func main() {
	_, err := do()
	fmt.Println(err)

	doNoErr()
	doThroughInt()
	doThroughIntWithoutWrap()
	doThroughIntWithWrap()
}

func do() ([]byte, error) {
	return json.Marshal(struct{}{}) // want `error returned from external package is unwrapped`
}

func doNoErr() bool {
	return strings.HasPrefix("hello world", "hello")
}

func doThroughInt() (bool, bool, bool, bool, error) {
	return testInt(impl{}).MultipleReturn() // want `error returned from interface method should be wrapped`
}

func doThroughIntWithoutWrap() error {
	return testInt(impl{}).ErrorReturn() // want `error returned from interface method should be wrapped`
}

func doThroughIntWithWrap() error {
	return fmt.Errorf("failed: %v", testInt(impl{}).ErrorReturn())
}

type testInt interface {
	MultipleReturn() (bool, bool, bool, bool, error)
	ErrorReturn() error
}

type impl struct{}

func (_ impl) MultipleReturn() (bool, bool, bool, bool, error) {
	return true, true, true, true, errors.New("uh oh")
}

func (_ impl) ErrorReturn() error {
	return errors.New("uh oh")
}
