package exceptions

import "errors"

var (
	ErrNotFound = errors.New("Your requested Item is not found")
	ErrConflict = errors.New("Your Item already exist")
)
