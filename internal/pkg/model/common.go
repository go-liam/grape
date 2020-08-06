package models

// APIResponse :
type APIResponse struct {
	Code    int         `json:"code" binding:"required"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}
