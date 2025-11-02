package types

type APIResponse struct {
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}
