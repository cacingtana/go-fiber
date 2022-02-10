package business

import "errors"

var (
	ErrInternalServerError = errors.New("Internal Server Error")

	ErrHasBeenModified = errors.New("Data has been modified")

	ErrNotFound = errors.New("Data was not found")

	ErrInvalidSpec = errors.New("Data is not valid")
)
