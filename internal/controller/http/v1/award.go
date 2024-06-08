package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/romeros69/basket/internal/entity"
	"github.com/romeros69/basket/internal/usecase"
	"github.com/romeros69/basket/pkg/logger"
	"net/http"
	"strconv"
)

type awardRoutes struct {
	a usecase.Award
	l logger.Interface
}

func newAwardRoutes(handler *gin.RouterGroup, a usecase.Award, l logger.Interface) {
	r := awardRoutes{
		a: a,
		l: l,
	}

	h := handler.Group("/award")
	{
		h.POST("", r.createAward)
		h.GET("/:id", r.getAward)
		h.PUT("/:id", r.updateAward)
		h.DELETE("/:id", r.deleteAward)
		h.GET("/list", r.listAwards)
	}
}

type createAwardResp struct {
	AwardID string `json:"awardID"`
}

// @Summary Create award
// @Tags award
// @Description Create new award
// @ID create-award
// @Accept json
// @Produce json
// @Param award body entity.Award true "Enter new award info"
// @Success 201 {object} createAwardResp
// @Failure 500 {object} errResponse
// @Router /award [post]
func (ar *awardRoutes) createAward(c *gin.Context) {
	var awardParam entity.Award
	if err := c.ShouldBindJSON(&awardParam); err != nil {
		ar.l.Error(err.Error())
		prepareError(c, err)
		return
	}

	awardID, err := ar.a.CreateAward(c.Request.Context(), &awardParam)
	if err != nil {
		ar.l.Error(err.Error())
		prepareError(c, err)
		return
	}

	c.JSON(http.StatusCreated, createAwardResp{AwardID: awardID})
}

// @Summary Get award
// @Tags award
// @Description Get award by id
// @ID get-award
// @Produce json
// @Param id path string true "Enter award id"
// @Success 200 {object} entity.Award
// @Failure 400 {object} errResponse
// @Failure 404 {object} errResponse
// @Failure 500 {object} errResponse
// @Router /award/{id} [get]
func (ar *awardRoutes) getAward(c *gin.Context) {
	awardID := c.Param("id")

	award, err := ar.a.GetAward(c.Request.Context(), awardID)
	if err != nil {
		ar.l.Error(err.Error())
		prepareError(c, err)
		return
	}

	c.JSON(http.StatusOK, award)
}

// @Summary Update award
// @Tags award
// @Description Update award by id
// @ID update-award
// @Accept json
// @Produce json
// @Param id path string true "Enter id award"
// @Param award body entity.Award true "Enter new award info for update"
// @Success 200 {object} entity.Award
// @Failure 400 {object} errResponse
// @Failure 404 {object} errResponse
// @Failure 500 {object} errResponse
// @Router /award/{id} [put]
func (ar *awardRoutes) updateAward(c *gin.Context) {
	awardID := c.Param("id")

	var awardParam entity.Award
	if err := c.ShouldBindJSON(&awardParam); err != nil {
		ar.l.Error(err.Error())
		prepareError(c, err)
		return
	}

	newAward, err := ar.a.UpdateAward(c.Request.Context(), awardID, &awardParam)
	if err != nil {
		ar.l.Error(err.Error())
		prepareError(c, err)
		return
	}

	c.JSON(http.StatusOK, newAward)
}

// @Summary Delete award
// @Tags award
// @Description Delete award by id
// @ID delete-award
// @Param id path string true "Enter id award"
// @Success 204 {object} nil
// @Failure 400 {object} errResponse
// @Failure 404 {object} errResponse
// @Failure 500 {object} errResponse
// @Router /award/{id} [delete]
func (ar *awardRoutes) deleteAward(c *gin.Context) {
	awardID := c.Param("id")

	if err := ar.a.DeleteAward(c.Request.Context(), awardID); err != nil {
		ar.l.Error(err.Error())
		prepareError(c, err)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

// @Summary Get award list
// @Tags award
// @Description Get award list
// @ID get-award-list
// @Produce json
// @Param page_size query string false "Enter page size" example="10"
// @Param page_number query string false "Enter page number" example="1"
// @Success 200 {object} entity.Award
// @Failure 400 {object} errResponse
// @Failure 500 {object} errResponse
// @Router /award/list [get]
func (ar *awardRoutes) listAwards(c *gin.Context) {
	var pageSize, pageNumber int64

	pageSize, err := strconv.ParseInt(c.Query("page_size"), 10, 64)
	if err != nil {
		ar.l.Warn("use default page size 10, because: %s", err.Error())
		pageSize = defaultPageSize
	}

	pageNumber, err = strconv.ParseInt(c.Query("page_number"), 10, 64)
	if err != nil {
		ar.l.Warn("use default page number 1, because: %s", err.Error())
		pageNumber = defaultPageNumber
	}

	awards, err := ar.a.GetAwardList(c.Request.Context(), pageSize, pageNumber)
	if err != nil {
		ar.l.Error(err.Error())
		prepareError(c, err)
		return
	}

	c.JSON(http.StatusOK, awards)
}
