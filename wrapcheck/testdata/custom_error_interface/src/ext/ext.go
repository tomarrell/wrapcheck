package ext

type CustomErr struct{}

func (CustomErr) Error() string {
	return "custom error"
}

func Foo() CustomErr {
	return CustomErr{}
}

func Bar() (int, CustomErr) {
	return 0, CustomErr{}
}

type PtrErr struct{}

func (*PtrErr) Error() string {
	return "pointer error"
}

func Ptr() *PtrErr {
	return &PtrErr{}
}

type CodedErr interface {
	error
	Code() int
}

func Coded() CodedErr {
	return nil
}
