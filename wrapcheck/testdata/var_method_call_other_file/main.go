package main

func main() {
	do()
}

func do() error {
	return GlobalErr // TODO want `error returned from external package is unwrapped`
}
