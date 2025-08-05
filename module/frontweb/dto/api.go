package dto

type ApiResponse struct {
	Status     string      `json:"status"`
	Data       interface{} `json:"data"`
	Message    string      `json:"message"`
	StatusCode int         `json:"status_code"`
}
