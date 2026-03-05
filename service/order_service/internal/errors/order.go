package errors

import "errors"

var (
	ErrOrderNotFound = errors.New("order not found")
	ErrOrderAlreadyPaid = errors.New("order already paid")
	ErrOrderExpired = errors.New("order expired")
	ErrOrderFailed = errors.New("order failed")
	ErrOrderPending = errors.New("order pending")
)