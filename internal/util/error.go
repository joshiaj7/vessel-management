package util

import "github.com/pkg/errors"

func ErrorWrap(err error) error {
	if err == nil {
		return nil
	}
	return errors.WithStack(err)
}
