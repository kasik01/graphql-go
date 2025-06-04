package notification

import (
	"fmt"
	"graphql-hasura-demo/internal/base"
	"graphql-hasura-demo/internal/dto/hasura"
	"net/http"

	"github.com/gin-gonic/gin"
)

type controller struct {
	service *Service
}

func (c *controller) handleTaskUpdated(ctx *gin.Context) {
	var request hasura.UpdatedEventRequest[Grade]

	// body, _ := io.ReadAll(ctx.Request.Body)
	// fmt.Println(string(body))

	if err := ctx.BindJSON(&request); err != nil {
		ctx.Error(&Errors.InvalidRegisterPayload)
		fmt.Println("abcd : ", err)
		fmt.Println("123abcd : ", err.Error())
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
