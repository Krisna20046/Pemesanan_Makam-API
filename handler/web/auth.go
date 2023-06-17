package web

import (
	"embed"
	"net/http"
	"path"
	"text/template"

	"github.com/Krisna20046/client"
	"github.com/Krisna20046/model"
	"github.com/Krisna20046/service"

	"github.com/gin-gonic/gin"
)

type AuthWeb interface {
	Login(c *gin.Context)
	LoginProcess(c *gin.Context)
	Register(c *gin.Context)
	RegisterProcess(c *gin.Context)
	Logout(c *gin.Context)
}

type authWeb struct {
	userClient     client.UserClient
	sessionService service.SessionService
	embed          embed.FS
}

func NewAuthWeb(userClient client.UserClient, sessionService service.SessionService, embed embed.FS) *authWeb {
	return &authWeb{userClient, sessionService, embed}
}

func (a *authWeb) Login(c *gin.Context) {
	var filepath = path.Join("views", "auth", "login.html")
	var header = path.Join("views", "general", "header.html")


	var tmpl, err = template.ParseFS(a.embed, filepath, header)
	if err != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(c.Writer, nil)
	if err != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
	}
}

func (a *authWeb) LoginProcess(c *gin.Context) {
	username := c.Request.FormValue("username")
	password := c.Request.FormValue("password")
	var user model.User

	status, err := a.userClient.Login(username, password)
	if err != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	session, err := a.sessionService.GetSessionByUsername(username)
	if err != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	if status == 200 {
		http.SetCookie(c.Writer, &http.Cookie{
			Name:   "session_token",
			Value:  session.Token,
			Path:   "/",
			MaxAge: 31536000,
			Domain: "",
		})
		if user.Role == "admin" {
			c.Redirect(http.StatusSeeOther, "/client/admin")
		}
		if user.Role == "user" {
			c.Redirect(http.StatusSeeOther, "/client/user")
		}
	} else {
		c.Redirect(http.StatusSeeOther, "/client/login")
	}
}

func (a *authWeb) Register(c *gin.Context) {
	var filepath = path.Join("views", "auth", "register.html")
	var header = path.Join("views", "general", "header.html")


	var tmpl, err = template.ParseFS(a.embed, filepath, header)
	if err != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(c.Writer, nil)
	if err != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
	}
}

func (a *authWeb) RegisterProcess(c *gin.Context) {
	nama := c.Request.FormValue("nama")
	email := c.Request.FormValue("email")
	username := c.Request.FormValue("username")
	password := c.Request.FormValue("password")
	no_hp := c.Request.FormValue("no_hp")
	jenis_kelamin := c.Request.FormValue("jenis_kelamin")
	alamat := c.Request.FormValue("alamat")

	status, err := a.userClient.Register(nama, email, username, password, no_hp, jenis_kelamin, alamat)
	if err != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	if status == 201 {
		c.Redirect(http.StatusSeeOther, "/client/login")
	} else {
		c.Writer.WriteHeader(http.StatusInternalServerError)
	}
}

func (a *authWeb) Logout(c *gin.Context) {
	c.SetCookie("session_token", "", -1, "/", "", false, false)
	c.Redirect(http.StatusSeeOther, "/")
}
