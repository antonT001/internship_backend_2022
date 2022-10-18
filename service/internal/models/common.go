package models

type Out struct {
	Success bool    `json:"success"`
	Error   *string `json:"error,omitempty"`
}
