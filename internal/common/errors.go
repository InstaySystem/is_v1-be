package common

import "errors"

var (
	ErrUsernameAlreadyExists = errors.New("username đã tồn tại")
	
	ErrEmailAlreadyExists = errors.New("email đã tồn tại")
)