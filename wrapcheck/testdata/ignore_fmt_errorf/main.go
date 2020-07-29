package main

import "fmt"

func main() {
	do()
}

func do() error {
	err := fmt.Errorf("failed to do something")
	if err != nil {
		return err
	}

	return nil
}
