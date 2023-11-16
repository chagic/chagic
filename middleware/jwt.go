package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func Jwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := validateToken(c)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"Unauthorized": "Authentication required"})
			fmt.Println(err)
			c.Abort()
			return
		}
		c.Next()
	}
}

func validateToken(c *gin.Context) error {
	token, err := getToken(c)
	if err != nil {
		return err
	}
	decode, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		fmt.Println(decode)
		c.Set("username", decode["user"].(map[string]interface{})["username"])
		c.Set("userId", decode["user"].(map[string]interface{})["id"])
		return nil
	}

	return errors.New("Invalid token provided")
}

func getToken(c *gin.Context) (*jwt.Token, error) {

	tokenString := getTokenFromRequest(c)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	return token, err
}

func getTokenFromRequest(c *gin.Context) string {
	bearerToken := c.Request.Header.Get("Authorization")
	if bearerToken == "" {
		bearerToken = c.Request.Header.Get("Sec-WebSocket-Protocol")
	}

	splitToken := strings.Split(bearerToken, " ")
	if len(splitToken) == 2 {
		return splitToken[1]
	}
	return splitToken[0]
}
