package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

type User struct {
	Id       int    `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type UserData struct {
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

func main() {
	// koneksi ke db
	// <userdb>:<password>@/<db_name>
	db, err := sql.Open("mysql", "root:qwerty123@/db_sirclo_api")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	e := echo.New()
	// routing with query parameter
	// endpoint untuk mendapatkan seluruh data user
	e.GET("/users", func(c echo.Context) error {
		users := []User{}
		result, err := db.Query("select id, name, email, password from users")
		if err != nil {
			fmt.Println(err)
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    400,
				"status":  "failed",
				"message": "failed to fetch data",
			})
		}
		defer result.Close()

		for result.Next() {
			var user User
			err := result.Scan(&user.Id, &user.Name, &user.Email, &user.Password)
			if err != nil {
				return c.JSON(http.StatusBadRequest, FailedResponses(400, "failed to read data"))
			}
			user.Name = "halo " + user.Name
			users = append(users, user)
		}
		return c.JSON(http.StatusOK, SuccessResponses("success to read data", users))
	})
	e.POST("/users", func(c echo.Context) error {
		user := UserData{}
		if err := c.Bind(&user); err != nil {
			fmt.Println(err)
			return c.JSON(http.StatusBadRequest, FailedResponses(400, "failed to bind data"))
		}
		_, err := db.Exec("INSERT INTO users(name, email, password) VALUES(?,?,?)", user.Name, user.Email, user.Password)
		// _, err := db.Exec("UPDATE users SET name=?, email=?, password=? WHERE id=?", user.Name, user.Email, user.Password, id)
		if err != nil {
			return c.JSON(http.StatusBadRequest, FailedResponses(400, "failed to insert data"))
		}
		// defer insert.Close()
		return c.JSON(http.StatusOK, SuccessWithoutDataResponses("success insert data"))
	})
	e.GET("/users/:id", func(c echo.Context) error {
		userid, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, FailedResponses(400, "failed to convert id"))
		}
		result, err := db.Query("SELECT * FROM users WHERE id=?", userid)
		/*
			SQL INJECTION:
				"SELECT * FROM users WHERE id=?", userid
				id=1; DROP TABLE users;
				SELECT * FROM users WHERE id=1; DROP TABLE users;
		*/
		if err != nil {
			return c.JSON(http.StatusBadRequest, FailedResponses(400, "failed to fetch data"))
		}
		defer result.Close()

		if isExist := result.Next(); !isExist {
			return c.JSON(http.StatusBadRequest, FailedResponses(400, "data not found"))
		}

		var user User
		errScan := result.Scan(&user.Id, &user.Name, &user.Email, &user.Password)
		if errScan != nil {
			return c.JSON(http.StatusBadRequest, FailedResponses(400, "failed to read data"))
		}

		return c.JSON(http.StatusOK, SuccessResponses("success to read data", user))

	})
	// e.PUT("/users/:id", UpdateUserController)
	// e.DELETE("/users/:id", DeleteUserController)

	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8080"))
}

func SuccessResponses(message string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusOK,
		"status":  "success",
		"message": message,
		"data":    data,
	}
}

func SuccessWithoutDataResponses(message string) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusOK,
		"status":  "success",
		"message": message,
	}
}

func FailedResponses(code int, message string) map[string]interface{} {
	return map[string]interface{}{
		"code":    code,
		"status":  "failed",
		"message": message,
	}
}
