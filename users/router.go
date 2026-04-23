package users

import (
	"net/http"

	"github.com/fahrurben/geteventgo/common"
	"github.com/gin-gonic/gin"
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
