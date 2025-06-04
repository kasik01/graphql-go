package hasura

type HasuraRequest[T any] struct {
	Action struct {
		Name string `json:"name"`
	} `json:"action"`
	Input struct {
		Params T `json:"params"`
	} `json:"input"`
}
