package entity

import (
	"errors"
	"fmt"
	"net/http"
)

var (
	ErrorParamType = NewError("Wrong param type", http.StatusUnprocessableEntity)

	ErrorVoyageNotFound = NewError("Voyage is not found", http.StatusNotFound)

	ErrorVesselNotFound   = NewError("Vessel is not found", http.StatusNotFound)
	ErrorVesselDuplicated = NewError("Vessel has duplication", http.StatusBadRequest)
)

type RequestError struct {
	StatusCode int
	Err        error
}

func (r RequestError) Error() string {
	return fmt.Sprintf("status %d: err %v", r.StatusCode, r.Err)
}

func NewError(message string, code int) error {
	return RequestError{
		StatusCode: code,
		Err:        errors.New(message),
	}
}
