package router

import (
	handler "repocon/interface/api"
	"repocon/interface/middleware"

	"github.com/labstack/echo/v4"
)

func NewRouter(uh handler.IUserHandler) *echo.Echo {
	e := echo.New()
	api := e.Group("/api")
	api.Use(middleware.AuthMiddleware)
	api.POST("/user", uh.UpsertUser)
	api.GET("/users/:id", uh.GetUser)
	return e
}
