package events

import (
	"net/http"

	"github.com/fahrurben/geteventgo/common"
	"github.com/fahrurben/geteventgo/users"
	"github.com/gin-gonic/gin"
)

type EventHandler struct {
	Repo    *EventRepository
	Service *EventService
}

func EventEndpoints(router *gin.RouterGroup) {
	db := common.GetDB()
	repo := EventRepository{db: db}
	service := EventService{repo: &repo}
	h := EventHandler{Repo: &repo, Service: &service}

	router.POST("event", users.AuthMiddleware(true), h.EventCreate)
}

func (h *EventHandler) EventCreate(c *gin.Context) {
	eventValidator := EventValidator{}

	if err := c.ShouldBindJSON(&eventValidator); err != nil {
		c.JSON(http.StatusBadRequest, common.NewError("validation", err))
		return
	}

	event := eventValidator.toModel()
	event.Owner = c.MustGet("my_user_model").(users.UserModel)

	err := h.Service.Create(c, &event)

	if err != nil {
		c.JSON(http.StatusBadRequest, common.NewError("database", err))
		return
	}

	c.JSON(http.StatusOK, event)
}
