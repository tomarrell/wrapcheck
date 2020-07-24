package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	do()
}

func do() error {
	_, err := json.Marshal(struct{}{})
	if err != nil {
		return fmt.Errorf("failed to marshal struct", err)
	}

	return nil
}
