package main

import "encoding/json"

func main() {
	do()
}

func do() (err error) {
	err = json.Unmarshal([]byte(""), nil)
	return // TODO want `error returned from external package is unwrapped`
}
