package delivery

import (
	"net/http"
	"sirclo/restapi/db/gorm/datastore"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func CreateGetUsersController(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var users []datastore.User

		if err := db.Find(&users).Error; err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    400,
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
		c.Bind(&user)

		if err := db.Save(&user).Error; err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    400,
				"status":  "failed",
				"message": "failed to fetch data",
			})
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success create new user",
			"user":    user,
		})
	}
}
