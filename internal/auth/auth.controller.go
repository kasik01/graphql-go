package auth

import (
	"graphql-hasura-demo/internal/base"
	"graphql-hasura-demo/internal/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

type controller struct {
	service *service
}

func (c *controller) login(ctx *gin.Context) {
	var request dto.HasuraRequest[LoginRequest]

	if err := ctx.BindJSON(&request); err != nil {
		ctx.Error(&Errors.InvalidLoginRequest)
		return
	}

	res, err := c.service.login(request.Input.Params)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(base.NewApiMessage(http.StatusOK, res))
}

func newController() *controller {
	service := NewService()
	controller := controller{service}

	return &controller
}
