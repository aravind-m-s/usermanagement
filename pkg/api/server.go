package http

import (
	"usermanagement/pkg/api/handler"
	"usermanagement/pkg/config"

	"github.com/gin-gonic/gin"
)

type ServerHTTP struct {
	engine *gin.Engine
}

func NewServerHTTP(userHandler *handler.UserHandler) *ServerHTTP {
	engine := gin.New()

	engine.Use(gin.Logger())

	api := engine.Group("/api")

	api.GET("users", userHandler.GetAllUsers)

	api.GET("users/:id", userHandler.GetSingleUser)

	api.POST("users", userHandler.CreateUser)

	api.DELETE("users/:id", userHandler.DeleteUser)

	api.PUT("users/:id", userHandler.UpdateUser)

	return &ServerHTTP{engine}

}

func (sh *ServerHTTP) Start(cnf config.Config) {
	sh.engine.Run(cnf.Port)
}
