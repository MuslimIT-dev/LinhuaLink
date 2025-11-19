package model

type ErrorResponseCode401 struct {
	Error string `json:"error" example:"StatusUnauthorized"`
}
type ErrorResponseCode400 struct {
	Error string `json:"error" example:"StatusBadRequest"`
}
type ErrorResponseCode500 struct {
	Error string `json:"error" example:"StatusInternalServerError"`
}
