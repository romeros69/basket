package v1

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/romeros69/basket/internal/apperrors"
	"github.com/romeros69/basket/internal/entity"
	"github.com/romeros69/basket/internal/usecase"
	"github.com/romeros69/basket/pkg/logger"
	"net/http"
)

type playerRoutes struct {
	p usecase.Player
	l logger.Interface
}

func newPlayerRoutes(handler *gin.RouterGroup, p usecase.Player, l logger.Interface) {
	r := playerRoutes{
		p: p,
		l: l,
	}

	h := handler.Group("/player")
	{
		h.POST("", r.createPlayer)
		h.GET("/:id", r.getPlayer)
	}
}

func (pr *playerRoutes) createPlayer(c *gin.Context) {
	var playerParam entity.Player
	if err := c.ShouldBindJSON(&playerParam); err != nil {
		pr.l.Error(err.Error())
		prepareError(c, err)
		return
	}

	playerID, err := pr.p.CreatePlayer(c.Request.Context(), &playerParam)
	if err != nil {
		pr.l.Error(err.Error())
		prepareError(c, err)
		return
	}

	c.JSON(http.StatusCreated, playerID)
}

func (pr *playerRoutes) getPlayer(c *gin.Context) {
	playerID := c.Param("id")

	player, err := pr.p.GetPlayer(c.Request.Context(), playerID)
	if err != nil {
		pr.l.Error(err.Error())
		prepareError(c, err)
		return
	}

	c.JSON(http.StatusOK, player)
}

func prepareError(c *gin.Context, err error) {
	switch {
	case errors.Is(err, apperrors.ErrPlayerNotFound):
		errorResponse(c, http.StatusNotFound, apperrors.ErrPlayerNotFound.Error())
	case errors.Is(err, apperrors.ErrInvalidPlayerID):
		errorResponse(c, http.StatusBadRequest, apperrors.ErrInvalidPlayerID.Error())
	default:
		errorResponse(c, http.StatusInternalServerError, "internal error")
	}
}
