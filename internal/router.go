package internal

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "rest-calculator/docs"
)

func (handler *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.Use(handler.preflight)

	root := router.Group("/")
	{
		root.GET("/health", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{"status": "alive"})
		})

		root.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		root.POST("/execute", handler.ExecuteOperation)
	}

	return router
}

func (handler *Handler) preflight(ctx *gin.Context) {
	ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, User, accept, origin, Cache-Control, X-Requested-With")
	ctx.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, OPTIONS")
	if ctx.Request.Method == "OPTIONS" {
		ctx.AbortWithStatus(204)
	}
	ctx.Next()
}
