package ext

type CustomErr struct{}

func (CustomErr) Error() string {
	return "custom error"
}

func Foo() CustomErr {
	return CustomErr{}
}
