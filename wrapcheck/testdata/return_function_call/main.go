package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	_, err := do()
	fmt.Println(err)
}

func do() ([]byte, error) {
	return json.Marshal(struct{}{}) // want `error returned from external package is unwrapped`
}
