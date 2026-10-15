package main

import "example.com/ignorelocal/local"

func main() {
	_ = do()
}

func do() error {
	if err := local.Error(); err != nil {
		return err
	}

	return nil
}
