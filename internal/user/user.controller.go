package user

import (
	"graphql-hasura-demo/internal/base"
	"graphql-hasura-demo/internal/dto/hasura"
	"net/http"

	"github.com/gin-gonic/gin"
)

type controller struct {
	service *Service
}

func (c *controller) Register(ctx *gin.Context) {
	var request hasura.HasuraRequest[RegisterUserRequest]

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.Error(&Errors.InvalidRegisterPayload)
		return
	}

	res, err := c.service.Register(request.Input.Params)
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
