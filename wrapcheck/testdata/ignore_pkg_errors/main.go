package main

import (
	"fmt"

	"github.com/pkg/errors"
)

func main() {
	do()
}

func do() error {
	err := fmt.Errorf("failed to do something")
	if err != nil {
		return errors.Wrap(err, "uh oh")
	}

	if err != nil {
		return errors.Wrapf(err, "uh oh")
	}

	if err != nil {
		return errors.WithMessage(err, "uh oh")
	}

	if err != nil {
		return errors.WithMessagef(err, "uh %s", "oh")
	}

	if err != nil {
		return errors.WithStack(err)
	}

	return errors.New("uh oh")
}
