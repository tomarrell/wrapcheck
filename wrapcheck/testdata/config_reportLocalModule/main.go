package main

import "example.com/reportlocal/local"

func main() {
	_ = do()
}

func do() error {
	if err := local.Error(); err != nil {
		return err
	}

	return nil
}
