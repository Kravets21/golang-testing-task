package errors

import "github.com/pkg/errors"

var (
	ErrRequestEmptyBody = func() error {
		return errors.New("request body is empty")
	}
	ErrReadAll = func(err error) error {
		return errors.Wrap(err, "readAll failed")
	}
	ErrJSONUnmarshal = func(err error) error {
		return errors.Wrap(err, "json unmarshal failed")
	}
	ErrJSONMarshal = func(err error) error {
		return errors.Wrap(err, "json marshal failed")
	}
)
