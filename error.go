package talkback

import "errors"

var (
	ErrInvalidField   = errors.New("invalid field")
	ErrInvalidOp      = errors.New("invalid op")
	ErrInvalidPreload = errors.New("invalid preload")
)
