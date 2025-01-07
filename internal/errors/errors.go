package internalerrors

import "errors"

var ErrRecordNotFound = errors.New("record not found error")

var ErrResourceNotFound = errors.New("resource not found error")

var ErrParsingForm = errors.New("form parsing error")
