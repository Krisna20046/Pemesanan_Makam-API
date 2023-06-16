package api

import (
	"github.com/Krisna20046/model"
	"github.com/Krisna20046/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserAPI interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
	GetUsersByRole(c *gin.Context)
}

type userAPI struct {
	userService service.UserService
}

func NewUserAPI(userService service.UserService) *userAPI {
	return &userAPI{userService}
}

func (u *userAPI) Register(c *gin.Context) {
	var user model.UserRegister

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, model.NewErrorResponse("invalid decode json"))
		return
	}

	if user.Email == "" || user.Password == "" || user.Username == "" {
		c.JSON(http.StatusBadRequest, model.NewErrorResponse("register data is empty"))
		return
	}

	if user.Role == "" {
		user.Role = "user"
	}

	var recordUser = model.User{
		Nama:         user.Nama,
		Role:         user.Role,
		Username:     user.Username,
		Email:        user.Email,
		Password:     user.Password,
		NoHP:         user.NoHP,
		JenisKelamin: user.JenisKelamin,
		Alamat:       user.Alamat,
	}

	recordUser, err := u.userService.Register(&recordUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.NewErrorResponse("error internal server"))
		return
	}

	c.JSON(http.StatusCreated, model.NewSuccessResponse("register success"))
}

func (u *userAPI) Login(c *gin.Context) {
	var user model.UserLogin

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, model.NewErrorResponse("invalid decode json"))
		return
	}

	if user.Username == "" || user.Password == "" {
		c.JSON(http.StatusBadRequest, model.NewErrorResponse("username or password is empty"))
		return
	}

	token, err := u.userService.Login(&model.User{
		Username: user.Username,
		Password: user.Password,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.NewErrorResponse("error internal server"))
		return
	}

	c.SetCookie("session_token", *token, 3600, "/", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"user_id": user.Username,
		"message": "login success",
	})
}

func (u *userAPI) GetUsersByRole(c *gin.Context) {
	// Pengecekan apakah pengguna yang melakukan permintaan memiliki role admin
	if c.GetString("role") != "admin" {
		c.JSON(http.StatusUnauthorized, model.NewErrorResponse("Only admin can perform this action"))
		return
	}

	role := c.Param("role")
	users, err := u.userService.GetUsersByRole(role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.NewErrorResponse("error internal server"))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}
