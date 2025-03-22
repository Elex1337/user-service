package route

import (
	"github.com/Elex1337/user-service/api/handler"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func Routes(e *echo.Echo, userHandler handler.UserHandler) {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	users := e.Group("/users")

	users.POST("", userHandler.CreateUser)
	users.GET("/:id", userHandler.GetUserByID)
	users.PUT("", userHandler.UpdateUser)
	users.DELETE("/:id", userHandler.DeleteUser)

	e.GET("/health", func(c echo.Context) error {
		return c.JSON(200, map[string]string{
			"status": "ok",
		})
	})
	e.GET("/swagger/*", echoSwagger.WrapHandler)
}
