package auth

import (
	"graphql-hasura-demo/internal/base"
	"net/http"
)

var Errors = struct {
	UnAuthorized        base.ApiError
	InvalidToken        base.ApiError
	IncorrectPassword   base.ApiError
	InvalidLoginRequest base.ApiError
}{
	UnAuthorized: base.ApiError{
		Status:  http.StatusUnauthorized,
		Message: "unauthorized",
	},
	InvalidToken: base.ApiError{
		Status:  http.StatusUnauthorized,
		Message: "invalid token",
	},
	IncorrectPassword: base.ApiError{
		Status:  http.StatusBadRequest,
		Message: "wrong password",
	},
	InvalidLoginRequest: base.ApiError{
		Status:  http.StatusBadRequest,
		Message: "invalid login request",
	},
}
