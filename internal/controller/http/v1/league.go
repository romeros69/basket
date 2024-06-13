package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/romeros69/basket/internal/entity"
	"github.com/romeros69/basket/internal/usecase"
	"github.com/romeros69/basket/pkg/logger"
	"net/http"
	"strconv"
)

type leagueRoutes struct {
	lg usecase.League
	l  logger.Interface
}

func newLeagueRoutes(handler *gin.RouterGroup, lg usecase.League, l logger.Interface) {
	r := leagueRoutes{
		lg: lg,
		l:  l,
	}

	h := handler.Group("/league")
	{
		h.POST("", r.createLeague)
		h.GET("/:id", r.getLeague)
		h.PUT("/:id", r.updateLeague)
		h.DELETE("/:id", r.deleteLeague)
		h.GET("/list", r.listLeagues)
	}
}

type createLeagueResp struct {
	LeagueID string `json:"league_id"`
}

// @Summary Create league
// @Tags league
// @Description Create new league
// @ID create-league
// @Accept json
// @Produce json
// @Param league body entity.League true "Enter new league info"
// @Success 201 {object} createLeagueResp
// @Failure 500 {object} errResponse
// @Router /league [post]
func (lr *leagueRoutes) createLeague(c *gin.Context) {
	var leagueParam entity.League
	if err := c.ShouldBindJSON(&leagueParam); err != nil {
		lr.l.Error(err.Error())
		prepareError(c, err)
		return
	}

	leagueID, err := lr.lg.CreateLeague(c.Request.Context(), &leagueParam)
	if err != nil {
		lr.l.Error(err.Error())
		prepareError(c, err)
		return
	}

	c.JSON(http.StatusCreated, createLeagueResp{LeagueID: leagueID})
}

// @Summary Get league
// @Tags league
// @Description Get league by id
// @ID get-league
// @Produce json
// @Param id path string true "Enter league id"
// @Success 200 {object} entity.League
// @Failure 400 {object} errResponse
// @Failure 404 {object} errResponse
// @Failure 500 {object} errResponse
// @Router /league/{id} [get]
func (lr *leagueRoutes) getLeague(c *gin.Context) {
	leagueID := c.Param("id")

	league, err := lr.lg.GetLeague(c.Request.Context(), leagueID)
	if err != nil {
		lr.l.Error(err.Error())
		prepareError(c, err)
		return
	}

	c.JSON(http.StatusOK, league)
}

// @Summary Update league
// @Tags league
// @Description Update league by id
// @ID update-league
// @Accept json
// @Produce json
// @Param id path string true "Enter id league"
// @Param league body entity.League true "Enter new league info for update"
// @Success 200 {object} entity.League
// @Failure 400 {object} errResponse
// @Failure 404 {object} errResponse
// @Failure 500 {object} errResponse
// @Router /league/{id} [put]
func (lr *leagueRoutes) updateLeague(c *gin.Context) {
	leagueID := c.Param("id")

	var leagueParam entity.League
	if err := c.ShouldBindJSON(&leagueParam); err != nil {
		lr.l.Error(err.Error())
		prepareError(c, err)
		return
	}

	newLeague, err := lr.lg.UpdateLeague(c.Request.Context(), leagueID, &leagueParam)
	if err != nil {
		lr.l.Error(err.Error())
		prepareError(c, err)
		return
	}

	c.JSON(http.StatusOK, newLeague)
}

// @Summary Delete league
// @Tags league
// @Description Delete league by id
// @ID delete-league
// @Param id path string true "Enter id league"
// @Success 204 {object} nil
// @Failure 400 {object} errResponse
// @Failure 404 {object} errResponse
// @Failure 500 {object} errResponse
// @Router /league/{id} [delete]
func (lr *leagueRoutes) deleteLeague(c *gin.Context) {
	leagueID := c.Param("id")

	if err := lr.lg.DeleteLeague(c.Request.Context(), leagueID); err != nil {
		lr.l.Error(err.Error())
		prepareError(c, err)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

// @Summary Get league list
// @Tags league
// @Description Get league list
// @ID get-league-list
// @Produce json
// @Param page_size query string false "Enter page size" example="10"
// @Param page_number query string false "Enter page number" example="1"
// @Success 200 {object} entity.League
// @Failure 400 {object} errResponse
// @Failure 500 {object} errResponse
// @Router /league/list [get]
func (lr *leagueRoutes) listLeagues(c *gin.Context) {
	var pageSize, pageNumber int64

	pageSize, err := strconv.ParseInt(c.Query("page_size"), 10, 64)
	if err != nil {
		lr.l.Warn("use default page size 10, because: %s", err.Error())
		pageSize = defaultPageSize
	}

	pageNumber, err = strconv.ParseInt(c.Query("page_number"), 10, 64)
	if err != nil {
		lr.l.Warn("use default page number 1, because: %s", err.Error())
		pageNumber = defaultPageNumber
	}

	leagues, err := lr.lg.GetLeagueList(c.Request.Context(), pageSize, pageNumber)
	if err != nil {
		lr.l.Error(err.Error())
		prepareError(c, err)
		return
	}

	c.JSON(http.StatusOK, leagues)
}
