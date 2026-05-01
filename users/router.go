package users

import (
	"errors"
	"net/http"

	"github.com/fahrurben/geteventgo/common"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Handler struct {
	Repo    *UserRepository
	Service *UserService
}

func UserEndpoints(router *gin.RouterGroup) {
	db := common.GetDB()
	repo := UserRepository{db: db}
	service := UserService{repo: &repo}
	h := Handler{Repo: &repo, Service: &service}

	router.POST("register", h.UserRegister)
	router.POST("login", h.UserLogin)
}

func (h *Handler) UserRegister(c *gin.Context) {
	registerValidator := RegisterValidator{}

	if err := c.ShouldBindJSON(&registerValidator); err != nil {
		c.JSON(http.StatusBadRequest, common.NewError("validation", err))
		return
	}

	err := h.Service.Register(c, registerValidator)

	if err != nil {
		c.JSON(http.StatusBadRequest, common.NewError("database", err))
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func (h *Handler) UserLogin(c *gin.Context) {
	loginValidator := LoginValidator{}

	if err := c.ShouldBindJSON(&loginValidator); err != nil {
		c.JSON(http.StatusBadRequest, common.NewError("validation", err))
		return
	}

	userModel, err := h.Service.Login(c, loginValidator)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) || errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			c.JSON(http.StatusBadRequest, common.NewError("authentication", errors.New("Wrong username or password")))
			return
		}

		c.JSON(http.StatusBadRequest, common.NewError("database", err))
		return
	}

	userSerializer := UserSerializer{Model: userModel}

	c.JSON(http.StatusOK, userSerializer.Response())
}
