package apperrors

import "errors"

var (
	ErrPlayerNotFound  = errors.New("player not found")
	ErrInvalidPlayerID = errors.New("invalid player id")
)
