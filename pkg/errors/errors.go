package errors

import (
	"errors"
	"fmt"
)

const (
	ErrNotMatched = Error("not matched") // parser didn't match input
)

var ErrNotSupported = NewFatalError(Error("not supported"))

type (
	ParserError interface {
		error
		IsFatal() bool
	}
	Error      string
	fatalError struct {
		error
	}

	wrapErr struct {
		error
		cause error
	}
)

func (e fatalError) IsFatal() bool {
	return true
}

func (e fatalError) Error() string {
	return e.error.Error()
}

func (e fatalError) Wrap(cause error) error {
	return wrapErr{error: e, cause: cause}
}

func (e fatalError) Unwrap() error {
	return e.error
}

func (e fatalError) Is(err error) bool {
	return errors.Is(e.error, err)
}

func (e Error) IsFatal() bool {
	return false
}

func (e Error) Error() string {
	return string(e)
}

func (e Error) Wrap(cause error) error {
	return wrapErr{error: e, cause: cause}
}

func (e wrapErr) Error() string {
	return fmt.Sprintf("%s: %s", e.error, e.cause)
}

func (e wrapErr) Unwrap() error {
	return e.cause
}

func (e wrapErr) Is(err error) bool {
	return errors.Is(e.error, err) || errors.Is(e.cause, err)
}

func (e wrapErr) IsFatal() bool {
	return IsFatal(e.error) || IsFatal(e.cause)
}

func NewFatalError(err error) error {
	return fatalError{err}
}

func IsFatal(err error) bool {
	e, ok := err.(ParserError)
	return ok && e.IsFatal()
}
