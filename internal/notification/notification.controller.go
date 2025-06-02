package notification

import (
	"graphql-hasura-demo/internal/base"
	"net/http"

	"github.com/gin-gonic/gin"
)

type controller struct {
	service *Service
}

func (c *controller) handleTaskUpdated(ctx *gin.Context) {
	var request TaskUpdatedEvent

	if err := ctx.BindJSON(&request); err != nil {
		ctx.Error(&Errors.InvalidRegisterPayload)
		return
	}

	res, err := c.service.NotifyTaskUpdated(request)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(base.NewApiMessage(http.StatusCreated, res))
}

func newController() *controller {
	service := NewService()
	controller := controller{service}

	return &controller
}
