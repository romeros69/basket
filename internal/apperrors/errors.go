package apperrors

import "errors"

var (
	ErrPlayerNotFound          = errors.New("player not found")
	ErrInvalidPlayerID         = errors.New("invalid player id")
	ErrInvalidPlayerPageSize   = errors.New("invalid page size for listing player")
	ErrInvalidPlayerPageNumber = errors.New("invalid page number for listing player")
	ErrAwardNotFound           = errors.New("award not found")
	ErrInvalidAwardID          = errors.New("invalid award id")
	ErrInvalidAwardPageSize    = errors.New("invalid page size for listing award")
	ErrInvalidAwardPageNumber  = errors.New("invalid page number for listing award")
)
