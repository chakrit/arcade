package arcade

import "strings"

var (
	ErrPrecondition = NewError("precondition", "precondition error")
	ErrUnknown      = NewError("unknown", "unknown error")

	ErrIO        = NewError("io", "I/O error")
	ErrIntegrity = NewError("integrity", "integrity check failed")
)

type Error struct {
	code    string
	message string
	cause   error
}

func NewError(code, message string) *Error {
	return &Error{
		code:    code,
		message: message,
	}
}

// code needs to be updated with message because it's rare that we only code
// but not the message and it's likely the developer forgets to do it leading
// to inconsistent error object.
func (err *Error) WithCode(code, message string) *Error {
	clone := *err
	clone.code, clone.message = code, message
	return &clone
}

func (err *Error) WithMessage(message string) *Error {
	clone := *err
	clone.message = message
	return &clone
}

func (err *Error) WithCause(cause error) *Error {
	clone := *err
	clone.cause = cause
	return &clone
}

func (err *Error) Code() string {
	if err == nil {
		return ""
	} else {
		return err.code
	}
}

func (err *Error) Message() string {
	if err == nil {
		return ""
	} else {
		return err.message
	}
}

func (err *Error) Cause() error {
	if err == nil {
		return nil
	} else {
		return err.cause
	}
}

func (err *Error) String() string {
	return err.Error()
}

func (err *Error) Error() string {
	if err == nil {
		return ""
	}

	builder := strings.Builder{}
	if err.code != "" {
		builder.WriteRune('(')
		builder.WriteString(err.code)
		builder.WriteRune(')')
		builder.WriteRune(' ')
	}

	builder.WriteString(err.message)
	if err.cause != nil {
		builder.WriteString(" caused by ")
		builder.WriteString(err.cause.Error())
	}

	return builder.String()
}
