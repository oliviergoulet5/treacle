package models

type ExecuteRequest struct {
	Method  string            `json:"method"`
	URL     string            `json:"url"`
	Headers map[string]string `json:"headers"`
}

type ExecuteRequestResponse struct {
	StatusCode int                 `json:"statusCode"`
	Body       string              `json:"body"`
	Headers    map[string][]string `json:"headers"`
}
