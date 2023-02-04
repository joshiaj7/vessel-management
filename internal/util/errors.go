package util

import "github.com/pkg/errors"

type ErrorTracer struct {
	Message string
	Err     error
}

func ErrorWrap(err error) error {
	if err == nil {
		return nil
	}
	return errors.WithStack(err)
}

func NewTracer(message string) *ErrorTracer {
	return &ErrorTracer{
		Message: message,
	}
}

func TracerFromError(err error) *ErrorTracer {
	tracer := NewTracer(err.Error())
	tracer.Err = err
	_, ok := err.(StackTracer)
	if !ok {
		tracer.Err = errors.WithStack(err)
	}
	return tracer
}

type StackTracer interface {
	StackTrace() errors.StackTrace
}

func (e *ErrorTracer) Error() string {
	return e.Message
}

func (e *ErrorTracer) Unwrap() error {
	return e.Err
}

func (e *ErrorTracer) Wrap(err error) *ErrorTracer {
	e.Err = err
	_, ok := err.(StackTracer)
	if !ok {
		e.Err = errors.WithStack(err)
	}

	return e
}

func (e *ErrorTracer) StackTrace() errors.StackTrace {
	err := e.Unwrap()
	errWithStack, ok := err.(StackTracer)
	if ok {
		return errWithStack.StackTrace()
	}
	return nil
}
