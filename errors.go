package supermarket

import "errors"

var ErrInternalServerError = errors.New("internal server error")
var ErrNotFound = errors.New("not found")
var ErrPasswordNotMatch = errors.New("passwords does not match")
