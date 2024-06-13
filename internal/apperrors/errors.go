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
	ErrGameNotFound            = errors.New("game not found")
	ErrInvalidGameID           = errors.New("invalid game id")
	ErrInvalidGamePageSize     = errors.New("invalid page size for listing game")
	ErrInvalidGamePageNumber   = errors.New("invalid page number for listing game")
	ErrLeagueNotFound          = errors.New("league not found")
	ErrInvalidLeagueID         = errors.New("invalid league id")
	ErrInvalidLeaguePageSize   = errors.New("invalid page size fir listing league")
	ErrInvalidLeaguePageNumber = errors.New("invalid page number for listing league")
)
