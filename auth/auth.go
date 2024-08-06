package auth

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

// type User struct {
// 	Username string `json:"username"`
// 	Password string `json:"password"`
// }

var jwtKey = []byte("my_secret_key")

// func JWTMiddleware(c echo.Context) error {

// 	tokenString := c.Request().Header.Get("Authorization")

// 	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
// 		}
// 		return jwtKey, nil
// 	})

// 	if err != nil || !token.Valid {
// 		return c.JSON(http.StatusUnauthorized, err.Error())
// 	}

// 	claims := token.Claims.(jwt.MapClaims)
// 	c.Set("User", claims["User"])
// 	c.Set("Role", claims["Role"])

// 	return nil
// }

// Middleware for checking if the user is an admin
func AdminMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")[7:]
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			return c.JSON(http.StatusUnauthorized, err.Error())
		}

		claims := token.Claims.(jwt.MapClaims)
		if claims["Role"] != "admin" {
			return echo.ErrUnauthorized
		}

		return next(c)
	}
}

// Middleware for checking if the user is a regular user
func UserMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")[7:]
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			return c.JSON(http.StatusUnauthorized, err.Error())
		}

		claims := token.Claims.(jwt.MapClaims)
		role := claims["Role"].(string)
		if role != "user" && role != "admin" {
			return echo.ErrUnauthorized
		}
		return next(c)
	}
}

func GenerateJWT(username string, role string) (string, error) {
	claims := &jwt.MapClaims{
		"ExpiresAt": time.Now().Add(24 * time.Hour).Unix(),
		"User":      username,
		"Role":      role,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func Restricted(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	username := claims["sub"].(string)
	return c.JSON(http.StatusOK, echo.Map{"message": "Welcome " + username})
}
