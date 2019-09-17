package server

import "errors"

var (
	ErrNotExists = errors.New("object not exists")
	ErrExists    = errors.New("object already exists")
	ErrDenied    = errors.New("access denied")
)
