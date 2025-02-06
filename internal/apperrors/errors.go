package apperrors

import "errors"

var ErrRecordNotFound = errors.New("record not found")

var ErrResourceNotFound = errors.New("resource not found")

var ErrParsingForm = errors.New("parse form error")
