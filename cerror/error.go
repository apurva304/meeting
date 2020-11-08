package cerror

type Error struct {
	Msg        string
	StatusCode int
}

func New(msg string, statusCode int) error {
	return &Error{
		Msg:        msg,
		StatusCode: statusCode,
	}
}

func (e *Error) Error() string {
	return e.Msg
}

type ErrorWrapper struct {
	Error string `json:"error"`
}
