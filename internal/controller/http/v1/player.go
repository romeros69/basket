package v1

import (
	"github.com/gin-gonic/gin"
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
	}
}

func (pr *playerRoutes) createPlayer(c *gin.Context) {
	var playerParam entity.Player
	if err := c.ShouldBindJSON(&playerParam); err != nil {
		pr.l.Error("error parse param", err.Error())
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	playerID, err := pr.p.CreatePlayer(c.Request.Context(), &playerParam)
	if err != nil {
		pr.l.Error("create player", err)
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, playerID)
}
