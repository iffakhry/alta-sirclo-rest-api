package delivery

import (
	"fmt"
	"net/http"
	"sirclo/restapi/db/gorm/datastore"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func CreateGetUsersController(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		users, err := datastore.GetUsers(db)
		if err != nil {
			fmt.Println(err)
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    http.StatusInternalServerError,
				"status":  "failed",
				"message": "failed to fetch data",
			})
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success get all users",
			"users":   users,
		})
	}
}

func CreateAddUserController(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := datastore.User{}
		if err := c.Bind(&user); err != nil {
			fmt.Println(err)
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    http.StatusBadRequest,
				"status":  "failed",
				"message": "Bad request",
			})
		}
		if err := datastore.AddUser(db, &user); err != nil {
			fmt.Println(err)
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    http.StatusInternalServerError,
				"status":  "failed",
				"message": "failed to save data",
			})
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success create new user",
			"user":    user,
		})
	}
}
