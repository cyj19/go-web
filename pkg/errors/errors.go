package errors

type fundamental struct {
	code int
	msg  string
}

func (f *fundamental) Error() string {
	return f.msg
}

func New(message string) error {
	return &fundamental{
		msg: message,
	}
}
