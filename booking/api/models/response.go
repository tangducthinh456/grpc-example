package models

type Response struct {
	StatusCode int         `json:"status_code" binding:"required"`
	Message    []string    `json:"message,omitempty"`
	Payload    interface{} `json:"payload,omitempty"`
}
