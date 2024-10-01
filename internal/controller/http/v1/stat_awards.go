package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/romeros69/basket/internal/entity"
	"github.com/romeros69/basket/internal/usecase"
	"github.com/romeros69/basket/pkg/logger"
)

type statAwardsRoutes struct {
	sa usecase.StatAwards
	l  logger.Interface
}

func newStatAwardsRoutes(handler *gin.RouterGroup, sa usecase.StatAwards, l logger.Interface) {
	r := statAwardsRoutes{
		sa: sa,
		l:  l,
	}

	h := handler.Group("/stat_awards")
	{
		h.POST("", r.createRecord)
		h.GET("/tournament/:id", r.viewPlayersAndRewardsInTournament)
		h.GET("/match/:id", r.viewPlayersAndRewardsInMatch)
		h.GET("/player/:id", r.ViewRewardsForPlayer)
		h.GET("/reward/:id", r.ViewWhoGotSpecificReward)
	}
}

// @Summary Create record
// @Tags record
// @Description Create new record
// @ID create-record
// @Accept json
// @Produce json
// @Param player body entity.RewardStat true "Enter new record info"
// @Success 201 {object} nil
// @Failure 500 {object} errResponse
// @Router /stat_awards [post]
func (sr *statAwardsRoutes) createRecord(c *gin.Context) {
	var statParam entity.RewardStat
	if err := c.ShouldBindJSON(&statParam); err != nil {
		sr.l.Error(err.Error())
		prepareError(c, err)
		return
	}

	err := sr.sa.CreateRecord(c.Request.Context(), statParam)
	if err != nil {
		sr.l.Error(err.Error())
		prepareError(c, err)
		return
	}

	c.JSON(http.StatusCreated, nil)
}

// @Summary Get stat by tournament
// @Tags stat
// @Description Get stat by tournament id
// @ID get-stat-tournament
// @Produce json
// @Param id path string true "Enter tournament id"
// @Success 200 {object} []entity.RewardStat
// @Failure 500 {object} errResponse
// @Router /stat_awards/tournament/{id} [get]
func (sr *statAwardsRoutes) viewPlayersAndRewardsInTournament(c *gin.Context) {
	tournamentID := c.Param("id")

	result, err := sr.sa.ViewPlayersAndRewardsInTournament(c.Request.Context(), tournamentID)
	if err != nil {
		sr.l.Error(err.Error())
		prepareError(c, err)
		return
	}

	c.JSON(http.StatusOK, result)
}

// @Summary Get stat by match
// @Tags stat
// @Description Get stat by match id
// @ID get-stat-match
// @Produce json
// @Param id path string true "Enter match id"
// @Success 200 {object} []entity.RewardStat
// @Failure 500 {object} errResponse
// @Router /stat_awards/match/{id} [get]
func (sr *statAwardsRoutes) viewPlayersAndRewardsInMatch(c *gin.Context) {
	matchID := c.Param("id")

	result, err := sr.sa.ViewPlayersAndRewardsInMatch(c.Request.Context(), matchID)
	if err != nil {
		sr.l.Error(err.Error())
		prepareError(c, err)
		return
	}

	c.JSON(http.StatusOK, result)
}

// @Summary Get stat by player
// @Tags stat
// @Description Get stat by player id
// @ID get-stat-player
// @Produce json
// @Param id path string true "Enter player id"
// @Success 200 {object} []entity.RewardStat
// @Failure 500 {object} errResponse
// @Router /stat_awards/player/{id} [get]
func (sr *statAwardsRoutes) ViewRewardsForPlayer(c *gin.Context) {
	playerID := c.Param("id")

	result, err := sr.sa.ViewRewardsForPlayer(c.Request.Context(), playerID)
	if err != nil {
		sr.l.Error(err.Error())
		prepareError(c, err)
		return
	}

	c.JSON(http.StatusOK, result)
}

// @Summary Get stat by reward
// @Tags stat
// @Description Get stat by reward id
// @ID get-stat-reward
// @Produce json
// @Param id path string true "Enter reward id"
// @Success 200 {object} []entity.RewardStat
// @Failure 500 {object} errResponse
// @Router /stat_awards/reward/{id} [get]
func (sr *statAwardsRoutes) ViewWhoGotSpecificReward(c *gin.Context) {
	rewardID := c.Param("id")

	result, err := sr.sa.ViewWhoGotSpecificReward(c.Request.Context(), rewardID)
	if err != nil {
		sr.l.Error(err.Error())
		prepareError(c, err)
		return
	}

	c.JSON(http.StatusOK, result)
}
