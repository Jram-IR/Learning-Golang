package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gtihub.com/Jram-IR/go-auth-jwt/config"
	"gtihub.com/Jram-IR/go-auth-jwt/models"
)

func RegisterUser(c *gin.Context) {
	// get the user email and password from the req body
	var userData struct {
		Email    string
		Password string
	}

	if err := c.Bind(&userData); err != nil {
		c.String(http.StatusBadRequest, "could not process post body")
		return
	}
	//hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(userData.Password), 10)
	if err != nil {
		c.String(http.StatusBadRequest, "could not hash password")
		return
	}
	// create the user
	user := models.User{Email: userData.Email, Password: string(hash)}
	result := config.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "could not create user"})
		return
	}
	// respond
	c.IndentedJSON(http.StatusOK, user)

}

func LoginUser(c *gin.Context) {
	// get the user email and password from the req body
	var userData struct {
		Email    string
		Password string
	}

	if err := c.Bind(&userData); err != nil {
		c.String(http.StatusBadRequest, "could not process post body")
		return
	}
	//find with email the user
	var user models.User
	result := config.DB.Find(&user, "email = ?", userData.Email)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "could not find user"})
		return
	}

	//compare the password given and hashed one
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userData.Password))
	if err != nil {
		c.String(http.StatusBadRequest, "please check email or password")
		return
	}
	// create the JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": user.ID,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		c.String(http.StatusBadRequest, "error signing the token")
		return
	}

	//create the cookie
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24, "", "", false, true)

	//respond
	c.IndentedJSON(http.StatusOK, tokenString)

}

func ValidateUser(c *gin.Context) {
	user, _ := c.Get("user")
	c.IndentedJSON(http.StatusAccepted, user)
}
