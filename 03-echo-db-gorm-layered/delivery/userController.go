package delivery

import (
	"fmt"
	"net/http"
	"sirclo/restapi/db/gorm/datastore"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func CreateLoginController(db *gorm.DB, jwtSecret string) echo.HandlerFunc {
	return func(c echo.Context) error {
		// validasi user
		var identity datastore.Identity
		if err := c.Bind(&identity); err != nil {
			fmt.Println(err)
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    http.StatusBadRequest,
				"status":  "unauthorized",
				"message": "invalid request",
			})
		}
		user, err := datastore.GetUserByIdentity(db, identity)
		if err != nil {
			fmt.Println(err)
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"code":    http.StatusUnauthorized,
				"status":  "unauthorized",
				"message": "unauthorized access",
			})
		}
		// membuat token
		token, err := CreateToken(jwtSecret, user.Name)
		if err != nil {
			fmt.Println(err)
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    http.StatusInternalServerError,
				"status":  "internal server error",
				"message": "cannot create token",
			})
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"token": token,
			"name":  user.Name,
		})
	}
}

func CreateGetUsersController(db *gorm.DB, jwtSecret string) echo.HandlerFunc {
	return func(c echo.Context) error {
		currentUserName, err := GetUserName(jwtSecret, c)
		if err != nil {
			fmt.Println(err)
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"code":    http.StatusUnauthorized,
				"status":  "unauthorized",
				"message": "unauthorized access",
			})
		}
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
			"message":     "success get all users",
			"users":       users,
			"currentUser": currentUserName,
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
