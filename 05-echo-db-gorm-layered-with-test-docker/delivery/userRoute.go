package delivery

import (
	"fmt"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

func LogElapsedTime(handler echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		startTime := time.Now()
		err := handler(c)
		elapsed := time.Since(startTime)
		fmt.Println(elapsed)
		return err
	}
}

func InitUserRoute(e *echo.Echo, DB *gorm.DB, jwtSecret string) {
	e.Pre(middleware.RemoveTrailingSlash(), middleware.Logger())
	e.POST("/login", CreateLoginController(DB, jwtSecret))
	e.GET("/users", CreateGetUsersController(DB, jwtSecret), middleware.JWT([]byte(jwtSecret)))
	e.GET("/users/:id", CreateGetUserByIdController(DB, jwtSecret), middleware.JWT([]byte(jwtSecret)))
	e.POST("/users", CreateAddUserController(DB))

	// e.Use(
	// 	middleware.Logger(),
	// 	middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
	// 		// echo dono:password | base64
	// 		if username == "dono" && password == "password" {
	// 			return true, nil
	// 		}
	// 		return false, nil
	// 	}),
	// )

	// e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
	// 	Format: `[${time_rfc3339}] ${status} ${method} ${host}${path} ${latency_human}` + "\n",
	// }))

	// e.GET("/users", CreateGetUsersController(DB), middleware.Logger())
	// e.POST("/users", CreateAddUserController(DB), LogElapsedTime)
	// e.GET("/users/:id", GetUserController)
	// e.PUT("/users/:id", UpdateUserController)
	// e.DELETE("/users/:id", DeleteUserController)
}
