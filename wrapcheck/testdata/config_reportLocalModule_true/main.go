package main

import "example.com/reportlocal/local"

func main() {
	_ = do()
}

func do() error {
	if err := local.Error(); err != nil {
		return err // want `error returned from external package is unwrapped`
	}

	return nil
}
