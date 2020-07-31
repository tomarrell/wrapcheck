package main

import "encoding/json"

var (
	_, GlobalErr = json.Marshal(struct{}{})
)
