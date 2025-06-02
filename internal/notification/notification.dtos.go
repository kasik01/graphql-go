package notification

type (
	Task struct {
		Id          string `json:"id"`
		Description string `json:"description"`
		IsCompleted bool   `json:"is_completed"`
		UserId      string `json:"user_id"`
	}
	TaskUpdatedEvent struct {
		Event struct {
			SessionVariables struct {
				UserId string `json:"x-hasura-user-id"`
			} `json:"session_variables"`
			Data struct {
				Old Task `json:"old"`
				New Task `json:"new"`
			} `json:"data"`
		} `json:"event"`
	}

	TaskUpdatedEventResponse struct {
		Status  string `json:"status"`
		Message string `json:"message"`
	}
)
