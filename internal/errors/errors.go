package internalerrors

import "errors"

var ErrRecordNotFound = errors.New("failed to find record")

var ErrResourceNotFound = errors.New("failed to find resource")

var ErrParsingForm = errors.New("failed to parse form")
