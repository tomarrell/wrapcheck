package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func main() {
	var b bytes.Buffer

	test := func() error {
		return json.NewEncoder(&b).Encode("test")
	}

	fmt.Println(test())
}
