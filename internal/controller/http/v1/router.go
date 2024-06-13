package v1

import (
	"github.com/gin-gonic/gin"
	_ "github.com/romeros69/basket/docs"
	"github.com/romeros69/basket/internal/usecase"
	"github.com/romeros69/basket/pkg/logger"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// NewRouter
// Swagger spec:
// @title       Basket LAB
// @description Basket no-sql lab
// @version     1.0
// @host        localhost:8080
// @schemes 	http
// @BasePath    /v1
func NewRouter(handler *gin.Engine, p usecase.Player, a usecase.Award, g usecase.Game, l logger.Interface) {
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	swaggerHandler := ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, "DISABLE_SWAGGER_HTTP_HANDLER")
	handler.GET("/swagger/*any", swaggerHandler)

	h := handler.Group("/v1")
	{
		newPlayerRoutes(h, p, l)
		newAwardRoutes(h, a, l)
		newGameRoutes(h, g, l)
	}
}
