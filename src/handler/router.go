package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitRouting(echo *echo.Echo, messageHandler MessageHandler) {
	echo.Use(middleware.Logger())
	echo.GET("/messages", messageHandler.Index())
	echo.POST("/messages", messageHandler.Create())
}
