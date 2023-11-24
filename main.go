package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"

	"hugobde.dev/amaretti/user"
)

var HMAC_SECRET = []byte("IHateF1SoBad")

const TOKEN_COOKIE_NAME = "authToken"

func homePageGet(ctx *gin.Context) {
	fmt.Println(ctx.Keys)
	if session, ok := ctx.Keys["session"].(*user.Session); !ok {
		ctx.Redirect(http.StatusSeeOther, "/login")
	} else {
		ctx.HTML(http.StatusOK, "home.html", gin.H{
			"Username": session.User.Username,
		})
	}
}

func loginPageGet(ctx *gin.Context) {
	if _, ok := ctx.Keys["session"].(*user.Session); !ok {
		ctx.HTML(http.StatusOK, "login.html", gin.H{})
	} else {
		ctx.Redirect(http.StatusSeeOther, "/")
	}
}

func loginFormPost(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")

	if user.AuthenticateUser(username, password) {
		sessionUUID, err := user.NewSession(user.UserDB[username])

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// create user session
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"uuid": sessionUUID,
		})

		tokenStr, err := token.SignedString([]byte(HMAC_SECRET))

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.SetCookie(TOKEN_COOKIE_NAME, tokenStr, 3600, "/", "localhost", true, true)
		ctx.Redirect(http.StatusSeeOther, "/")
	} else {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
	}
}

func userSessionMidWare(ctx *gin.Context) {
	var sessionTokenStr string
	var sessionToken *jwt.Token
	var tokenClaims jwt.MapClaims
	var sessionUUID uuid.UUID
	var session *user.Session
	var ok bool
	var err error

	if sessionTokenStr, err = ctx.Cookie(TOKEN_COOKIE_NAME); err != nil {
		ctx.Next()
		return
	}

	if sessionToken, err = jwt.Parse(sessionTokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return HMAC_SECRET, nil
	}); err != nil {
		ctx.Next()
		return
	}

	if tokenClaims, ok = sessionToken.Claims.(jwt.MapClaims); !ok {
		ctx.Next()
		return
	}

	if _, ok = tokenClaims["uuid"]; !ok {
		ctx.Next()
		return
	}

	if sessionUUID, err = uuid.Parse(tokenClaims["uuid"].(string)); err != nil {
		ctx.Next()
		return
	}

	if session, ok = user.SessionDB[sessionUUID]; !ok {
		ctx.Next()
		return
	}

	ctx.Set("session", session)
	ctx.Next()
}

func main() {
	if err := user.NewUser("user1", "+Pwd123-"); err != nil {
		log.Fatalf("failed to create user: %s", err)
	}

	router := gin.Default()
	router.Static("/static", "./static/")
	router.LoadHTMLGlob("templates/*")

	router.Use(userSessionMidWare)

	router.GET("/", homePageGet)

	router.GET("/login", loginPageGet)
	router.POST("/login", loginFormPost)

	router.Run(":8080")
}
