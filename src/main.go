package main

import (
	"fmt"

	"github.com/gumpen/server-todo/src/handler"
	"github.com/gumpen/server-todo/src/injector"
	"github.com/labstack/echo/v4"
)

func main() {
	fmt.Println("server start")
	messageHandler := injector.InjectMessageHandler()
	echo := echo.New()
	handler.InitRouting(echo, messageHandler)
	echo.Logger.Fatal(echo.Start(":8080"))
}
