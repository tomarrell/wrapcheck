package local

import "errors"

func Error() error {
	return errors.New("failed")
}
