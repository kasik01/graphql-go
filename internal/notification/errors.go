package notification

import (
	"graphql-hasura-demo/internal/base"
	"net/http"
)

var Errors = struct {
	InvalidRegisterPayload base.ApiError
}{
	InvalidRegisterPayload: base.ApiError{
		Status:  http.StatusBadRequest,
		Message: "invalid payload",
	},
}
