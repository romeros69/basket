package v1

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/romeros69/basket/internal/apperrors"
	"github.com/romeros69/basket/internal/entity"
	"github.com/romeros69/basket/internal/usecase"
	"github.com/romeros69/basket/pkg/logger"
)

const (
	defaultPageSize   = 10
	defaultPageNumber = 1
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
		h.PUT("/:id", r.updatePlayer)
		h.DELETE("/:id", r.deletePlayer)
		h.GET("/list", r.listPlayers)
	}
}

type createPlayerResp struct {
	PlayerID string `json:"playerID"`
}

// @Summary Create player
// @Tags player
// @Description Create new player
// @ID create-player
// @Accept json
// @Produce json
// @Param player body entity.Player true "Enter new player info"
// @Success 201 {object} createPlayerResp
// @Failure 500 {object} errResponse
// @Router /player [post]
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

	c.JSON(http.StatusCreated, createPlayerResp{PlayerID: playerID})
}

// @Summary Get player
// @Tags player
// @Description Get player by id
// @ID get-player
// @Produce json
// @Param id path string true "id пользователя"
// @Success 200 {object} entity.Player
// @Failure 400 {object} errResponse
// @Failure 404 {object} errResponse
// @Failure 500 {object} errResponse
// @Router /player/{id} [get]
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

// @Summary Update player
// @Tags player
// @Description Update player by id
// @ID update-player
// @Accept json
// @Produce json
// @Param id path string true "Enter id player"
// @Param player body entity.Player true "Enter new player info for update"
// @Success 200 {object} entity.Player
// @Failure 400 {object} errResponse
// @Failure 404 {object} errResponse
// @Failure 500 {object} errResponse
// @Router /player/{id} [put]
func (pr *playerRoutes) updatePlayer(c *gin.Context) {
	playerID := c.Param("id")

	var playerParam entity.Player
	if err := c.ShouldBindJSON(&playerParam); err != nil {
		pr.l.Error(err.Error())
		prepareError(c, err)
		return
	}

	newPlayer, err := pr.p.UpdatePlayer(c.Request.Context(), playerID, &playerParam)
	if err != nil {
		pr.l.Error(err.Error())
		prepareError(c, err)
		return
	}

	c.JSON(http.StatusOK, newPlayer)
}

// @Summary Delete player
// @Tags player
// @Description Delete player by id
// @ID delete-player
// @Param id path string true "Enter id player"
// @Success 204 {object} nil
// @Failure 400 {object} errResponse
// @Failure 404 {object} errResponse
// @Failure 500 {object} errResponse
// @Router /player/{id} [delete]
func (pr *playerRoutes) deletePlayer(c *gin.Context) {
	playerID := c.Param("id")

	if err := pr.p.DeletePlayer(c.Request.Context(), playerID); err != nil {
		pr.l.Error(err.Error())
		prepareError(c, err)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

// @Summary Get player list
// @Tags player
// @Description Get player list
// @ID get-player-list
// @Produce json
// @Param page_size query string false "Enter page size" example="10"
// @Param page_number query string false "Enter page number" example="1"
// @Success 200 {object} entity.Player
// @Failure 400 {object} errResponse
// @Failure 500 {object} errResponse
// @Router /player/list [get]
func (pr *playerRoutes) listPlayers(c *gin.Context) {
	var pageSize, pageNumber int64

	pageSize, err := strconv.ParseInt(c.Query("page_size"), 10, 64)
	if err != nil {
		pr.l.Warn("use default page size 10, because: %s", err.Error())
		pageSize = defaultPageSize
	}

	pageNumber, err = strconv.ParseInt(c.Query("page_number"), 10, 64)
	if err != nil {
		pr.l.Warn("use default page number 1, because: %s", err.Error())
		pageNumber = defaultPageNumber
	}

	players, err := pr.p.GetPlayerList(c.Request.Context(), pageSize, pageNumber)
	if err != nil {
		pr.l.Error(err.Error())
		prepareError(c, err)
		return
	}

	c.JSON(http.StatusOK, players)
}

func prepareError(c *gin.Context, err error) {
	switch {
	case errors.Is(err, apperrors.ErrPlayerNotFound):
		errorResponse(c, http.StatusNotFound, err.Error())
	case errors.Is(err, apperrors.ErrInvalidPlayerID) ||
		errors.Is(err, apperrors.ErrInvalidPlayerPageSize) ||
		errors.Is(err, apperrors.ErrInvalidPlayerPageNumber):
		errorResponse(c, http.StatusBadRequest, err.Error())
	default:
		errorResponse(c, http.StatusInternalServerError, "internal error")
	}
}
