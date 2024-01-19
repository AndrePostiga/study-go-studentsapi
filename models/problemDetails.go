package models

type ProblemDetails struct {
	Type   string   `json:"type"`
	Title  string   `json:"title"`
	Status int      `json:"status"`
	Detail []string `json:"detail"`
}
