package storager

import "errors"

var (
	ErrorEntityNotFound = errors.New("entity not found")
	ErrorDuplicate      = errors.New("duplicate error")
)
