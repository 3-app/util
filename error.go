package util

type Err struct {
	Msg string
}

func (e *Err) Error() string  {
	return e.Msg
}

func NewErr(str string) *Err  {
	return &Err{Msg: str}
}