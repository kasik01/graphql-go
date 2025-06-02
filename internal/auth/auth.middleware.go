package auth

import (
	"strings"

	"github.com/gin-gonic/gin"
)

type Middleware struct {
	service *service
}

func toPayload(rawPayload any) (*TokenPayload, error) {
	payloadMap, ok := rawPayload.(map[string]interface{})
	if !ok {
		return nil, &Errors.InvalidToken
	}

	var payload TokenPayload

	if id, ok := payloadMap["Id"].(string); ok {
		payload.Id = id
	} else {
		return nil, &Errors.InvalidToken
	}

	return &payload, nil
}

func (m *Middleware) JwtAuthMiddleware(c *gin.Context) {
	authHeader := c.Request.Header.Get("Authorization")
	splitedStrings := strings.Split(authHeader, " ")

	if len(splitedStrings) != 2 {
		c.Error(&Errors.UnAuthorized)
		c.Abort()
		return
	}

	authToken := splitedStrings[1]
	payload, err := m.service.verifyToken(authToken)
	if err != nil {
		c.Error(err)
		c.Abort()
		return
	}

	c.Set("x-user-id", payload.Id)
	c.Next()
}

func NewMiddleware() *Middleware {
	service := NewService()
	return &Middleware{service}
}
