package main

import "crypto/aes"

func main() {
	do()
}

func do() (out []byte, err error) {
	block, err := aes.NewCipher([]byte("test_key"))
	if err != nil {
		return nil, err // want `error returned from external package is unwrapped`
	}

	_ = block

	return nil, nil
}
