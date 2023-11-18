package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"

	"hugobde.dev/amaretti/user"
)

func homeHandleFunc(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "login.html", gin.H{})
}

func loginHandleFunc(ctx *gin.Context) {
	const hmacSecret = "IHateF1SoBad"

	username := ctx.PostForm("username")
	password := ctx.PostForm("password")

	if user.AuthenticateUser(username, password) {
		// create user session
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"uuid": user.UserDB[username].UUID,
		})

		tokenStr, err := token.SignedString([]byte(hmacSecret))

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"token": tokenStr})
	} else {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
	}
}

func main() {
	if err := user.NewUser("user1", "+Pwd123-"); err != nil {
		log.Fatalf("failed to create user: %s", err)
	}
	router := gin.Default()
	router.Static("/static", "./static/")
	router.LoadHTMLGlob("templates/*")
	router.Handle("GET", "/", homeHandleFunc)
	router.Handle("POST", "/login", loginHandleFunc)
	router.Run(":8080")
}
