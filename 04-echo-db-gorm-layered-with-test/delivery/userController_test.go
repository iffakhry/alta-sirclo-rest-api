package delivery

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"sirclo/restapi/db/gorm/datastore"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func InitEchoTestAPI() (*echo.Echo, *gorm.DB) {
	jwtSecret := os.Getenv("JWT_SECRET")
	connectionString := os.Getenv("DB_CONNECTION_STRING")
	db, err := datastore.InitDB(connectionString)
	if err != nil {
		panic(err)
	}
	if err = datastore.InitialMigration(db); err != nil {
		panic(err)
	}

	e := echo.New()
	InitUserRoute(e, db, jwtSecret)
	return e, db
}

func TestLoginSuccess(t *testing.T) {
	var testCases = struct {
		name       string
		path       string
		expectCode int
	}{

		name:       "success login",
		path:       "/login",
		expectCode: http.StatusOK,
	}

	e, db := InitEchoTestAPI()
	requestBody, _ := json.Marshal(map[string]string{
		"name":     "alta1",
		"password": "12345",
	})
	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
	res := httptest.NewRecorder()
	req.Header.Set("Content-Type", "application/json")
	context := e.NewContext(req, res)
	context.SetPath("/login")

	type Responses struct {
		Token string `json="token"`
		Name  string `json="name"`
	}
	if assert.NoError(t, CreateLoginController(db, "rahasia")(context)) {
		bodyResponses := res.Body.String()
		var response Responses

		err := json.Unmarshal([]byte(bodyResponses), &response)
		if err != nil {
			assert.Error(t, err, "error")
		}

		assert.Equal(t, testCases.expectCode, res.Code)
		assert.Equal(t, "alta1", response.Name)
	}

}

func TestLoginFailedBind(t *testing.T) {
	var testCases = struct {
		name       string
		path       string
		expectCode int
		status     string
	}{

		name:       "failed bind data",
		path:       "/login",
		expectCode: http.StatusBadRequest,
		status:     "unauthorized",
	}

	e, db := InitEchoTestAPI()
	requestBody, _ := json.Marshal(map[string]interface{}{
		"name":     "alta1",
		"password": 12345,
	})
	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
	res := httptest.NewRecorder()
	req.Header.Set("Content-Type", "application/json")
	context := e.NewContext(req, res)
	context.SetPath("/login")

	type Responses struct {
		Code    string `json="code"`
		Status  string `json="status"`
		Message string `json="message"`
	}
	if assert.NoError(t, CreateLoginController(db, "rahasia")(context)) {
		bodyResponses := res.Body.String()
		var response Responses

		err := json.Unmarshal([]byte(bodyResponses), &response)
		if err != nil {
			assert.Error(t, err, "error")
		}

		assert.Equal(t, testCases.expectCode, res.Code)
		assert.Equal(t, testCases.status, response.Status)
		assert.Equal(t, "invalid request", response.Message)
	}

}

func TestGetUsersController(t *testing.T) {
	t.Run("success get all users", func(t *testing.T) {

		e, db := InitEchoTestAPI()
		// requestBody, _ := json.Marshal(map[string]interface{}{
		// 	"name":     "alta1",
		// 	"password": 12345,
		// })

		token, err := CreateToken(os.Getenv("JWT_SECRET"), "alta")
		fmt.Println(token)
		if err != nil {
			panic(err)
		}
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/users")

		// jika menggunakan param:
		// context.SetParamNames("id")
		// context.SetParamValues("1")

		type Responses struct {
			Message     string `json="message"`
			Users       string `json="users"`
			CurrentUser string `json="currentUser"`
		}
		if assert.NoError(t, middleware.JWT([]byte("rahasia"))(CreateGetUsersController(db, "rahasia"))(context)) {
			bodyResponses := res.Body.String()
			var response Responses

			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}

			assert.Equal(t, http.StatusOK, res.Code)
			assert.Equal(t, "success get all users", response.Message)
			assert.Equal(t, "alta", response.CurrentUser)
		}

		// if assert.NoError(t, CreateGetUsersController(db, os.Getenv("JWT_SECRET"))(context)) {
		// 	bodyResponses := res.Body.String()
		// 	var response Responses

		// 	err := json.Unmarshal([]byte(bodyResponses), &response)
		// 	if err != nil {
		// 		assert.Error(t, err, "error")
		// 	}

		// 	assert.Equal(t, http.StatusOK, res.Code)
		// 	assert.Equal(t, "success get all users", response.Message)
		// }

	})
}
