package errors

import "errors"

var (
	ErrTrainNotFound = errors.New("train not found")
	ErrTrainCodeAlreadyExists = errors.New("train code already exists")
)