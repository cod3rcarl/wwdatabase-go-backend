package errors

import "errors"

var (
	ErrNoChampionsReturned = errors.New("no champions returned from query")
	ErrInvalidNullTime     = errors.New("timestamp provided is not valid")
)
