package supermarket

import "errors"

var (
	ErrInternalServerError = errors.New("internal server error")
	ErrNotFound            = errors.New("not found")
	ErrPasswordNotMatch    = errors.New("passwords does not match")
	ErrTokenExpired        = errors.New("token expired")
	ErrTokenIatClaim       = errors.New("token must have iat claim")
	ErrTokenSubjectClaim   = errors.New("token must have sub claim")
)
