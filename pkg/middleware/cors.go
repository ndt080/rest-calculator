package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func CORS() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"POST, GET, PUT, DELETE, OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, User, accept, origin, Cache-Control, X-Requested-With"},
		ExposeHeaders:    []string{"Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, User, accept, origin, Cache-Control, X-Requested-With"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	})
}
