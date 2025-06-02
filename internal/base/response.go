package base

type ApiReponse struct {
	Data  *any    `json:"data"`
	Error *string `json:"message"`
}

func NewApiMessage(status int, data any) (int, *ApiReponse) {
	return status, &ApiReponse{&data, nil}
}
