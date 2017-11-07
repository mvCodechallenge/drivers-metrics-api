package controllers

const JSON_DATA = "json"

/**
	Type for API common JSON response
 */
type APIResponse struct {
	Error string `json:"error"`
	Data  interface{} `json:"data"`
}
