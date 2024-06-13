package v1

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/romeros69/basket/internal/apperrors"
	"net/http"
)

type errResponse struct {
	Error string `json:"error" example:"message"`
}

func errorResponse(c *gin.Context, code int, msg string) {
	c.AbortWithStatusJSON(code, errResponse{msg})
}

func prepareError(c *gin.Context, err error) {
	switch {
	case errors.Is(err, apperrors.ErrPlayerNotFound) ||
		errors.Is(err, apperrors.ErrAwardNotFound) ||
		errors.Is(err, apperrors.ErrGameNotFound):
		errorResponse(c, http.StatusNotFound, err.Error())
	case errors.Is(err, apperrors.ErrInvalidPlayerID) ||
		errors.Is(err, apperrors.ErrInvalidPlayerPageSize) ||
		errors.Is(err, apperrors.ErrInvalidPlayerPageNumber) ||
		errors.Is(err, apperrors.ErrInvalidAwardID) ||
		errors.Is(err, apperrors.ErrInvalidAwardPageNumber) ||
		errors.Is(err, apperrors.ErrInvalidAwardPageSize) ||
		errors.Is(err, apperrors.ErrInvalidGameID) ||
		errors.Is(err, apperrors.ErrInvalidGamePageSize) ||
		errors.Is(err, apperrors.ErrInvalidGamePageNumber):
		errorResponse(c, http.StatusBadRequest, err.Error())
	default:
		errorResponse(c, http.StatusInternalServerError, "internal error")
	}
}
