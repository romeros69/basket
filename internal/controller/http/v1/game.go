package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/romeros69/basket/internal/entity"
	"github.com/romeros69/basket/internal/usecase"
	"github.com/romeros69/basket/pkg/logger"
	"net/http"
	"strconv"
)

type gameRoutes struct {
	g usecase.Game
	l logger.Interface
}

func newGameRoutes(handler *gin.RouterGroup, g usecase.Game, l logger.Interface) {
	r := gameRoutes{
		g: g,
		l: l,
	}

	h := handler.Group("/game")
	{
		h.POST("", r.createGame)
		h.GET("/:id", r.getGame)
		h.PUT("/:id", r.updateGame)
		h.DELETE("/:id", r.deleteGame)
		h.GET("/list", r.listGames)
	}
}

type createGameResp struct {
	GameID string `json:"game_id"`
}

// @Summary Create game
// @Tags game
// @Description Create new game
// @ID create-game
// @Accept json
// @Produce json
// @Param game body entity.Game true "Enter new game info"
// @Success 201 {object} createGameResp
// @Failure 500 {object} errResponse
// @Router /game [post]
func (gr *gameRoutes) createGame(c *gin.Context) {
	var gameParam entity.Game
	if err := c.ShouldBindJSON(&gameParam); err != nil {
		gr.l.Error(err.Error())
		prepareError(c, err)
		return
	}

	gameID, err := gr.g.CreateGame(c.Request.Context(), &gameParam)
	if err != nil {
		gr.l.Error(err.Error())
		prepareError(c, err)
		return
	}

	c.JSON(http.StatusCreated, createGameResp{GameID: gameID})
}

// @Summary Get game
// @Tags game
// @Description Get game by id
// @ID get-game
// @Produce json
// @Param id path string true "Enter game id"
// @Success 200 {object} entity.Game
// @Failure 400 {object} errResponse
// @Failure 404 {object} errResponse
// @Failure 500 {object} errResponse
// @Router /game/{id} [get]
func (gr *gameRoutes) getGame(c *gin.Context) {
	gameID := c.Param("id")

	game, err := gr.g.GetGame(c.Request.Context(), gameID)
	if err != nil {
		gr.l.Error(err.Error())
		prepareError(c, err)
		return
	}

	c.JSON(http.StatusOK, game)
}

// @Summary Update game
// @Tags game
// @Description Update game by id
// @ID update-game
// @Accept json
// @Produce json
// @Param id path string true "Enter id game"
// @Param award body entity.Game true "Enter new game info for update"
// @Success 200 {object} entity.Game
// @Failure 400 {object} errResponse
// @Failure 404 {object} errResponse
// @Failure 500 {object} errResponse
// @Router /game/{id} [put]
func (gr *gameRoutes) updateGame(c *gin.Context) {
	gameID := c.Param("id")

	var gameParam entity.Game
	if err := c.ShouldBindJSON(&gameParam); err != nil {
		gr.l.Error(err.Error())
		prepareError(c, err)
		return
	}

	newGame, err := gr.g.UpdateGame(c.Request.Context(), gameID, &gameParam)
	if err != nil {
		gr.l.Error(err.Error())
		prepareError(c, err)
		return
	}

	c.JSON(http.StatusOK, newGame)
}

// @Summary Delete game
// @Tags game
// @Description Delete game by id
// @ID delete-game
// @Param id path string true "Enter id game"
// @Success 204 {object} nil
// @Failure 400 {object} errResponse
// @Failure 404 {object} errResponse
// @Failure 500 {object} errResponse
// @Router /game/{id} [delete]
func (gr *gameRoutes) deleteGame(c *gin.Context) {
	gameID := c.Param("id")

	if err := gr.g.DeleteGame(c.Request.Context(), gameID); err != nil {
		gr.l.Error(err.Error())
		prepareError(c, err)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

// @Summary Get game list
// @Tags game
// @Description Get game list
// @ID get-game-list
// @Produce json
// @Param page_size query string false "Enter page size" example="10"
// @Param page_number query string false "Enter page number" example="1"
// @Success 200 {object} entity.Game
// @Failure 400 {object} errResponse
// @Failure 500 {object} errResponse
// @Router /game/list [get]
func (gr *gameRoutes) listGames(c *gin.Context) {
	var pageSize, pageNumber int64

	pageSize, err := strconv.ParseInt(c.Query("page_size"), 10, 64)
	if err != nil {
		gr.l.Warn("use default page size 10, because: %s", err.Error())
		pageSize = defaultPageSize
	}

	pageNumber, err = strconv.ParseInt(c.Query("page_number"), 10, 64)
	if err != nil {
		gr.l.Warn("use default page number 1, because: %s", err.Error())
		pageNumber = defaultPageNumber
	}

	games, err := gr.g.GetGameList(c.Request.Context(), pageSize, pageNumber)
	if err != nil {
		gr.l.Error(err.Error())
		prepareError(c, err)
		return
	}

	c.JSON(http.StatusOK, games)
}
