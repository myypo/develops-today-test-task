package http

import (
	"errors"
	"fmt"
	"net/http"
)

var (
	errBadRequest = errors.New("Bad request input. Please, provide valid data")

	errUnauthorized = errors.New("Access denied. Please, authorize")

	errConflict = errors.New("Please, try providing alternative input")

	errNotFound = errors.New(
		"Sorry, but it seems like what you have requested does not exist. Please, try providing different input",
	)

	errInternal = errors.New(
		"We are sorry, but something went wrong. Please, try again later",
	)

	errForbidden = errors.New("Sorry, but the action you have tried to take is forbidden")
)

type ErrorHttp interface {
	error
	Code() int
}

type errorHttp struct {
	msg  error
	code int
}

func (e errorHttp) Error() string {
	return e.msg.Error()
}

func (e errorHttp) Code() int {
	return e.code
}

func NewErrConflict(errToWrap error) ErrorHttp {
	return errorHttp{
		msg:  fmt.Errorf("%w: %w", errConflict, errToWrap),
		code: http.StatusConflict,
	}
}

func NewErrUnauthorized(errToWrap error) ErrorHttp {
	return errorHttp{
		msg:  fmt.Errorf("%w: %w", errUnauthorized, errToWrap),
		code: http.StatusUnauthorized,
	}
}

func NewErrForbidden(errToWrap error) ErrorHttp {
	return errorHttp{
		msg:  fmt.Errorf("%w: %w", errForbidden, errToWrap),
		code: http.StatusForbidden,
	}
}

func NewErrBadRequest(errToWrap error) ErrorHttp {
	return errorHttp{
		msg:  fmt.Errorf("%w: %w", errBadRequest, errToWrap),
		code: http.StatusBadRequest,
	}
}

func NewErrInternal(errToWrap error) ErrorHttp {
	return errorHttp{
		msg:  fmt.Errorf("%w: %w", errInternal, errToWrap),
		code: http.StatusInternalServerError,
	}
}

func NewErrNotFound(errToWrap error) ErrorHttp {
	return errorHttp{
		msg:  fmt.Errorf("%w: %w", errNotFound, errToWrap),
		code: http.StatusNotFound,
	}
}
