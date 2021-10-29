package errors

type Mod2Error struct {
	msg string
}

func NewMod2Error(msg string) error {
	return &Mod2Error{msg}
}

func (e Mod2Error) Error() string {
	return e.msg
}

type Mod3Error struct {
	msg string
}

func NewMod3Error(msg string) error {
	return &Mod3Error{msg}
}

func (e Mod3Error) Error() string {
	return e.msg
}
