package errors

import "errors"

var (
	ErrTrainSeatNotFound    = errors.New("train seat not found")
	ErrScheduleSeatNotFound = errors.New("schedule seat not found")
	ErrSeatAlreadyBooked    = errors.New("seat already booked")
)
