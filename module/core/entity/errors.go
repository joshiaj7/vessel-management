package entity

import (
	"errors"
	"fmt"
	"net/http"
)

var (
	ErrorShipmentNotFound = NewError("Shipment is not found", http.StatusNotFound)

	ErrorVesselNotFound   = NewError("Vessel is not found", http.StatusNotFound)
	ErrorVesselDuplicated = NewError("Shipment has duplication", http.StatusBadRequest)
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
