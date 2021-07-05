package main

import (
	"encoding/json"
)

func main() {
	do()
}

func do() error {
	_, err := json.Marshal(struct{}{})
	if err != nil {
		return err
	}

	return nil
}
