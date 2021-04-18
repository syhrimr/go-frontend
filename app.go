package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var host = "http://skilvul-gc.com"
var backendHost = "http://skilvul-gc.com:7070"
var groupchatBackendHost = "http://skilvul-gc.com:8080"

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.GET("/register", renderRegisterPage)
	r.GET("/login", renderLoginPage)
	r.GET("/groupchat/list", renderGetJoinedGroupchats)
	r.Static("/css", "./css")
	r.Run(":80")
}

func renderLoginPage(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"login.html",
		gin.H{
			"title":        "Login",
			"host":         host,
			"account_host": backendHost,
			"gc_host":      groupchatBackendHost,
		},
	)
}

func renderRegisterPage(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", gin.H{
		"title":        "Register",
		"host":         host,
		"account_host": backendHost,
		"gc_host":      groupchatBackendHost,
	})
}

func renderGetJoinedGroupchats(c *gin.Context) {
	c.HTML(http.StatusOK, "groupchatlist.html", gin.H{
		"title":        "Groupchats",
		"host":         host,
		"account_host": backendHost,
		"gc_host":      groupchatBackendHost,
	})
}
