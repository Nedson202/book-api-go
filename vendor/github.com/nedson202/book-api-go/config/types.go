package config

// Payload structure for error responses
type Payload struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
}
