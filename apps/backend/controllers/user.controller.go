package controllers

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	"log"
	"net/http"

	"github.com/ANU7MADHAV/algo-arena/services"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	// "github.com/golang-jwt/jwt/v5"
)

var user services.User

func CreateToken(email string) (string, error) {
	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		log.Fatal(err)
	}

	claims := &jwt.StandardClaims{
		ExpiresAt: 15000,
		Issuer:    "test",
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	tokenString, err := token.SignedString(key)
	fmt.Println("tokensharing", tokenString)

	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func GetAllUsers(c *gin.Context) {
	users, err := user.GetAllUsers()

	if err != nil {
		log.Println(err)
		return
	}
	c.JSON(200, gin.H{"message": users})
}

func CreateUsers(c *gin.Context) {
	var entry services.User
	if err := c.ShouldBindJSON(&entry); err != nil {
		c.Error(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	user, err := user.ChecKUser(entry)

	if err == nil {
		users, err := user.CreateUser(entry)

		if err != nil {
			log.Println(err)
			return
		}
		token, err := CreateToken(users.Email)
		if err != nil {
			log.Println(err)
			return
		}

		fmt.Println("token", token)
		c.JSON(http.StatusOK, token)
	}
	c.JSON(http.StatusFound, "Email already exist")

}

func UpdateUser(c *gin.Context) {
	var entry services.User
	id := c.Param("id")

	if err := c.ShouldBindJSON(&entry); err != nil {
		c.Error(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	users, err := user.UpdateUser(id, entry)

	if err != nil {
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, users)

}
