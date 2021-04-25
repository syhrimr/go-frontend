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
<<<<<<< HEAD
	r.GET("/groupchat/chat", renderGetJoinedGroupchats)
	r.GET("/groupchat/list", renderRoomList)
=======
	r.GET("/groupchat/list", renderGetJoinedGroupchats)
	r.GET("/groupchat/explore", renderExplorePage)
>>>>>>> 4650c47088f1c618aeb548ad17da31239d20ec71
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
	c.HTML(
		http.StatusOK,
		"register.html",
		gin.H{
			"title":        "Register",
			"host":         host,
			"account_host": backendHost,
			"gc_host":      groupchatBackendHost,
		})
}

func renderGetJoinedGroupchats(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"groupchatlist.html",
		gin.H{
			"title":        "Groupchats",
			"host":         host,
			"account_host": backendHost,
			"gc_host":      groupchatBackendHost,
		})
}

func renderRoomList(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"roomlist.html",
		gin.H{
			"title":        "Roomlist",
			"host":         host,
			"account_host": backendHost,
			"gc_host":      groupchatBackendHost,
		},
	)
}

func renderExplorePage(c *gin.Context)  {
	c.HTML(http.StatusOK, "explore.html", gin.H{
		"title":	"Explore Group",
		"host":		host,
		"account_host": backendHost,
		"gc_host":      groupchatBackendHost,

	})
}
