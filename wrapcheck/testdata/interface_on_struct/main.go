package main

type errorer interface {
	Decode(v interface{}) error
	decode(v interface{}) error
}

type foo struct {
	bar errorer
}

func main() {
	do(foo{})
	doInternal(foo{})
}

func do(f foo) error {
	var str string
	err := f.bar.Decode(&str)
	if err != nil {
		return err // want `error returned from interface method should be wrapped`
	}

	return nil
}

func doInternal(f foo) error {
	var str string
	err := f.bar.decode(&str)
	if err != nil {
		return err // unexported methods are validated at their implementation
	}

	return nil
}
