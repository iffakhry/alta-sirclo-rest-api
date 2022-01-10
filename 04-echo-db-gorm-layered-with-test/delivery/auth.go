package delivery

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func CreateToken(jwtSecret, userName string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userName"] = userName
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix() //Token expires after 1 hour
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwtSecret))
}

func GetUserName(jwtSecret string, e echo.Context) (string, error) {
	user := e.Get("user").(*jwt.Token)
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		userName := claims["userName"].(string)
		if userName == "" {
			return userName, fmt.Errorf("empty username")
		}
		return userName, nil
	}
	return "", fmt.Errorf("invalid user")
}
