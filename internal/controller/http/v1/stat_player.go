package v1

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/romeros69/basket/internal/entity"
	"github.com/romeros69/basket/internal/usecase"
	"github.com/romeros69/basket/pkg/logger"
)

type statPlayerRoutes struct {
	sp usecase.StatPlayer
	l  logger.Interface
}

func newStatPlayerRoutes(handler *gin.RouterGroup, sp usecase.StatPlayer, l logger.Interface) {
	r := statPlayerRoutes{
		sp: sp,
		l:  l,
	}

	h := handler.Group("/stat_player")
	{
		h.POST("", r.insertPlayer)
		h.GET("/:pid/:mid", r.getPlayerStatsByIDAndMatch)
		h.GET("/goals/:mid", r.getPlayersWithAvgGoalsGreaterThanByMatch)
		h.GET("/all_points/:mid", r.getPlayersWithTotalAvgStatsGreaterThanByMatch)
	}
}

// @Summary Create stat player
// @Tags stat-player
// @Description Create new stat player
// @ID create-stat-player
// @Accept json
// @Produce json
// @Param player body entity.PlayerStat true "Enter new player stat"
// @Success 201 {object} nil
// @Failure 500 {object} errResponse
// @Router /stat_player [post]
func (sr *statPlayerRoutes) insertPlayer(c *gin.Context) {
	var statPlayerParam entity.PlayerStat
	if err := c.ShouldBindJSON(&statPlayerParam); err != nil {
		sr.l.Error(err.Error())
		prepareError(c, err)
		return
	}

	err := sr.sp.InsertPlayerStat(c.Request.Context(), statPlayerParam)
	if err != nil {
		sr.l.Error(err.Error())
		prepareError(c, err)
		return
	}

	c.JSON(http.StatusCreated, nil)
}

// @Summary Get player stats by player id and match id
// @Tags player-stats
// @Description Get player stats by player id and match id
// @ID get-player-stats
// @Produce json
// @Param pid path string true "Enter player id"
// @Param mid path string true "Enter match id"
// @Success 200 {object} []entity.PlayerStat
// @Failure 500 {object} errResponse
// @Router /stat_player/{pid}/{mid} [get]
func (sr *statPlayerRoutes) getPlayerStatsByIDAndMatch(c *gin.Context) {
	// принимаем player id и match id
	playerID := c.Param("pid")
	matchID := c.Param("mid")

	result, err := sr.sp.GetPlayerStatsByIDAndMatch(c.Request.Context(), playerID, matchID)
	if err != nil {
		sr.l.Error(err.Error())
		prepareError(c, err)
		return
	}

	c.JSON(http.StatusOK, result)
}

// @Summary Get players stat avg goals by match id
// @Tags player-stats
// @Description Get players stat avg goals by match id
// @ID get-player-goals
// @Produce json
// @Param goals query string true "Enter avg goals"
// @Param mid path string true "Enter match id"
// @Success 200 {object} []entity.PlayerStat
// @Failure 500 {object} errResponse
// @Router /stat_player/goals/{mid} [get]
func (sr *statPlayerRoutes) getPlayersWithAvgGoalsGreaterThanByMatch(c *gin.Context) {
	// принимаем минимальное число голов (как паармтер) и матч ид
	matchID := c.Param("mid")
	minAVGGoals := c.Query("goals")

	goals, err := strconv.ParseFloat(minAVGGoals, 64)
	if err != nil {
		sr.l.Error(err.Error())
		prepareError(c, err)
		return
	}

	result, err := sr.sp.GetPlayersWithAvgGoalsGreaterThanByMatch(c.Request.Context(), goals, matchID)
	if err != nil {
		sr.l.Error(err.Error())
		prepareError(c, err)
		return
	}

	c.JSON(http.StatusOK, result)

}

// @Summary Get players stat avg points by match id
// @Tags player-stats
// @Description Get players stat avg points by match id
// @ID get-player-points
// @Produce json
// @Param points query string true "Enter avg points"
// @Param mid path string true "Enter match id"
// @Success 200 {object} []entity.PlayerStat
// @Failure 500 {object} errResponse
// @Router /stat_player/all_points/{mid} [get]
func (sr *statPlayerRoutes) getPlayersWithTotalAvgStatsGreaterThanByMatch(c *gin.Context) {
	// принимамем минимальное число всего суммы и матч id
	matchID := c.Param("mid")
	minAllVal := c.Query("points")

	points, err := strconv.ParseFloat(minAllVal, 64)
	if err != nil {
		sr.l.Error(err.Error())
		prepareError(c, err)
		return
	}

	result, err := sr.sp.GetPlayersWithTotalAvgStatsGreaterThanByMatch(c.Request.Context(), points, matchID)
	if err != nil {
		sr.l.Error(err.Error())
		prepareError(c, err)
		return
	}

	c.JSON(http.StatusOK, result)
}
