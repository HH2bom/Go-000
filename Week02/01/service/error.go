package service

import (
	"fmt"
)

const (
	notFoundCode = "0001"
	xxxxxxxxCode = "0000"
)

type Err struct {
	code string
	info string
	err  error
}

func (e *Err) Cause() error {
	return e.err
}

func (e *Err) Unwrap() error {
	return e.err
}

func (e *Err) Error() string {
	// return fmt.Sprintf("failed by code: %v ,because %v, sub err stack trace: %+v", e.code, e.info, e.err)
	return fmt.Sprintf("failed by code: %v ,because %v, sub err stack trace: %v", e.code, e.info, e.err)
}
