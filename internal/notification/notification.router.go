package notification

import (
	"github.com/gin-gonic/gin"
)

type router struct {
	controller *controller
	engine     *gin.Engine
}

func (r *router) Init() {
	group := r.engine.Group("/notifications")
	group.POST("/", r.controller.handleTaskUpdated)
}

func NewRouter(e *gin.Engine) *router {
	controller := newController()
	router := router{controller, e}

	return &router
}
