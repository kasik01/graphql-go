package hasura

type (
	UpdatedEventRequest[T any] struct {
		Event struct {
			Data struct {
				New T `json:"new"`
				Old T `json:"old"`
			} `json:"data"`
			SessionVariables struct {
				// UserId string `json:"x-hasura-user-id"`
			} `json:"session_variables"`
		} `json:"event"`
	}
)

type UpdatedEventResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
