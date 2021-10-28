package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

type errorer interface {
	Decode(v interface{}) error
}

func main() {
	do1()

	d := json.NewDecoder(strings.NewReader("hello world"))
	do2(d)

	do3(5)
	do4(5)
}

func do1() error {
	_, err := json.Marshal(struct{}{}) // external package ignored by package+method-name regexp
	if err != nil {
		return err
	}

	return nil
}

func do2(fn errorer) error {
	var str string
	err := fn.Decode(&str) // interface ignored by regexp
	if err != nil {
		return err
	}

	return nil
}

func do3(i int) error {
	if i%2 == 0 {
		return errors.NewMod2Error(fmt.Sprintf("%d is an even number", i)) // external package ignored by method-name regexp
	}

	return nil
}

func do4(i int) error {
	if i%3 == 0 {
		return errors.NewMod3Error(fmt.Sprintf("%d is divisible by 3", i)) // external package ignored by method-name regexp
	}

	return nil
}
