package errors

import "errors"

var (
	// ErrNotFound - key not found on storage
	ErrNotFound = errors.New("not found")
)
