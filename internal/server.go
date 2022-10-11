package internal

import (
	"github.com/braintree/manners"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"rest-calculator/pkg/logger"
	"rest-calculator/pkg/middleware"
	"time"
)

type Server struct {
	instance *manners.GracefulServer
}

func (server *Server) Run() error {
	logger.InitStdoutLogger()
	router := Inject()
	server.instance = manners.NewWithServer(&http.Server{
		Addr:           ":8000",
		Handler:        router,
		MaxHeaderBytes: 1048576, // 1 << 20
		ReadTimeout:    30 * time.Second,
		WriteTimeout:   30 * time.Second,
	})

	log.Infoln("Server is running")
	return server.instance.ListenAndServe()
}

func (server *Server) Shutdown() {
	log.Infoln("Shutting down")
	server.instance.Close()
}

func Inject() *gin.Engine {
	handler := NewHandler()

	router := handler.InitRoutes()
	router.Use(logger.GetLoggerHandlerFunc())
	router.Use(middleware.CORS())

	return router
}
