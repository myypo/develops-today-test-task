package data

import (
	"fmt"
	errVerb "sca/internal/error/verbose"
	"slices"
)

type DataError interface {
	errVerb.VerboseError
	NotFound() bool
	Conflict() bool
	BadRequest() bool
}

type dataError struct {
	userErr error
	techErr error
	errType DataErrorType
}

func (e *dataError) NotFound() bool {
	return e.errType == NotFound
}

func (e *dataError) Conflict() bool {
	return e.errType == Conflict
}

func (e *dataError) BadRequest() bool {
	return e.errType == BadRequest
}

func (e *dataError) Error() string {
	return e.userErr.Error()
}

func (e *dataError) Verbose() error {
	return e.techErr
}

type DataErrorType int

const (
	NotFound DataErrorType = iota
	Conflict
	BadRequest
	Internal
)

func newErrNotFound(modName string, techErr error) DataError {
	return &dataError{
		userErr: errNotFound(modName),
		techErr: techErr,
		errType: NotFound,
	}
}

func newErrConflict(modName, fieldName string, techErr error) DataError {
	return &dataError{
		userErr: errConflict(modName, fieldName),
		techErr: techErr,
		errType: Conflict,
	}
}

func newErrBadRequest(modName string, techErr error) DataError {
	return &dataError{
		userErr: errBadRequest(modName),
		techErr: techErr,
		errType: BadRequest,
	}
}

func newErrInternal(modName string, techErr error) DataError {
	return &dataError{
		userErr: errInternal(modName),
		techErr: techErr,
		errType: Internal,
	}
}

func errBadRequest(modName string) error {
	return fmt.Errorf("the %s related request is invalid", modName)
}

func errNotFound(modName string) error {
	return fmt.Errorf("the requested %s does not exist", modName)
}

func errConflict(modName, field string) error {
	return fmt.Errorf("the provided %s data is using duplicate value for %s", modName, field)
}

func errInternal(modName string) error {
	return fmt.Errorf("unexpected error occured when performing operation on %s", modName)
}

func NewErr(
	errTyp DataErrorType,
	techErr error,
	modName string,
	fieldName string,
	expTypes ...DataErrorType,
) DataError {
	if !slices.Contains(expTypes, errTyp) {
		return newErrInternal(modName, techErr)
	}

	switch errTyp {
	case Conflict:
		return newErrConflict(modName, fieldName, techErr)
	case NotFound:
		return newErrNotFound(modName, techErr)
	case BadRequest:
		return newErrBadRequest(modName, techErr)
	default:
		return newErrInternal(modName, techErr)
	}
}
