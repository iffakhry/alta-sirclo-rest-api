package delivery

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitUserRoute(e *echo.Echo, DB *gorm.DB) {
	e.GET("/users", CreateGetUsersController(DB))
	e.POST("/users", CreateAddUserController(DB))
	// e.GET("/users/:id", GetUserController)
	// e.PUT("/users/:id", UpdateUserController)
	// e.DELETE("/users/:id", DeleteUserController)
}
