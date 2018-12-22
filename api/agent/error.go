package agent

import "github.com/fnproject/fn/api/models"

// FuncError is an error that is the function's fault, that uses the
// models.APIError but distinguishes fault to function specific errors
type FuncError interface {
	models.APIError

	// hidden method needed for duck typing
	userError()
}

type concFuncError struct {
	models.APIError
}

func (c concFuncError) userError() {}

// NewFuncError returns a FuncError
func NewFuncError(err models.APIError) error { return concFuncError{err} }

// IsFuncError checks if err is of type FuncError
func IsFuncError(err error) bool { _, ok := err.(FuncError); return ok }
