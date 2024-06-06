package middleware

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"gtihub.com/Jram-IR/go-auth-jwt/config"
	"gtihub.com/Jram-IR/go-auth-jwt/models"
)

func CheckAuth(c *gin.Context) {
	//get the cookie
	fmt.Println("middleware")
	tokenString, err := c.Cookie("Authorization")

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	//parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("error in the signing method used %v", token.Header["alg"])
		}

		return []byte(os.Getenv("SECRET")), nil
	})

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	//get the claims
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		//fetch the user
		var user models.User
		config.DB.Find(&user, claims["userId"])

		c.Set("user", user)

	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}
