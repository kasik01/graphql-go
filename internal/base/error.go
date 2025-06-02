package base

import (
	"log"
	"net/http"
	"runtime"

	"github.com/gin-gonic/gin"
)

type ApiError struct {
	Status  int
	Message string
}

func (e *ApiError) Error() string {
	return e.Message
}

func logError(err error) {
	trace := make([]byte, 1024)
	runtime.Stack(trace, true)
	log.Printf("ERROR: %s\n%s", err, trace)
}

func NewErrorMessgae(err error) (int, *ApiReponse) {
	var apiError *ApiError

	// This the best way to log?
	logError(err)

	switch e := err.(type) {
	case *ApiError:
		apiError = e
	default:
		apiError = &ApiError{http.StatusInternalServerError, "Internal Server Error"}
	}

	msg := apiError.Error()
	return apiError.Status, &ApiReponse{nil, &msg}
}

func ErrorHandler(c *gin.Context) {
	c.Next()

	if len(c.Errors) == 0 {
		return
	}

	// log.Println(c.Errors[len(c.Errors) - 1].Err.Error())
	c.JSON(NewErrorMessgae(c.Errors[len(c.Errors)-1].Err))
}
