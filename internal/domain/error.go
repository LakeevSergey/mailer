package domain

import "errors"

var (
	ErrorEntityNotFound = errors.New("entity not found")
	ErrorDuplicate      = errors.New("duplicate error")

	ErrorTemplateNotFound      = errors.New("template not found")
	ErrorTemplateCodeDuplicate = errors.New("template with same code already exists")
	ErrorEmptyCodeBodyTitle    = errors.New("template cody or body and title should not be empty")
)
