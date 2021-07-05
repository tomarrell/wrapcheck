package main

import (
	"encoding/json"
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

	_, err = json.Marshal(struct{}{})
	if err != nil {
		return err // want `error returned from external package is unwrapped`
	}

	return errors.New("uh oh")
}
