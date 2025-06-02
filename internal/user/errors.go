package user

import (
	"graphql-hasura-demo/internal/base"
	"net/http"
)

var Errors = struct {
	NotFound               base.ApiError
	InvalidRegisterPayload base.ApiError
	DuplicatedUsername     base.ApiError
}{
	NotFound: base.ApiError{
		Status:  http.StatusNotFound,
		Message: "user not found",
	},
	InvalidRegisterPayload: base.ApiError{
		Status:  http.StatusBadRequest,
		Message: "invalid register payload",
	},
	DuplicatedUsername: base.ApiError{
		Status:  http.StatusBadRequest,
		Message: "username duplicated",
	},
}
