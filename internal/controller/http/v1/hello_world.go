package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/romeros69/basket/internal/usecase"
	"github.com/romeros69/basket/pkg/logger"
	"net/http"
)

type helloWorldRoutes struct {
	hw usecase.HelloWorld
	l  logger.Interface
}

func newHelloWorldRoutes(handler *gin.RouterGroup, hw usecase.HelloWorld, l logger.Interface) {
	r := helloWorldRoutes{hw: hw, l: l}

	h := handler.Group("/test-hello")
	{
		h.GET("/hi", r.helloHandler)
	}
}

// @Summary     Hello world
// @Description Print hello world
// @ID          hello-world
// @Tags  	    hello
// @Accept      json
// @Produce     json
// @Success     200 {object} string
// @Failure     500 {object} response
// @Router      /test-hello/hi [get]
func (r *helloWorldRoutes) helloHandler(c *gin.Context) {
	req, err := r.hw.Hello(c.Request.Context())
	if err != nil {
		r.l.Error(err, "http - v1 - hello")
	}

	c.JSON(http.StatusOK, req)
}
