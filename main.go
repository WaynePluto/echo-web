package main

import (
	"echo-web/custom"
	_ "echo-web/docs"
	"echo-web/modules/user"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host 127.0.0.1:8000

func main() {
	e := echo.New()

	custom.InitValidator(e)

	e.Use(custom.InitCtx)
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	user.SetRouter(e)

	e.Logger.SetLevel(2)

	e.Logger.Fatal(e.Start(":8000"))

}
