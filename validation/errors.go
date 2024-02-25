package validation

import "errors"

var (
	ErrCaptionLength = errors.New("0-1024 characters after entities parsing")
)
