package httpclient

import (
	"io"
	"net/http"
	"strings"

	"github.com/oliviergoulet5/treacle/internal/models"
)

// Execute performs the HTTP request described by request. It returns the
// response status code, headers, body, and any error encountered.
func Execute(request models.ExecuteRequest) (*models.ExecuteRequestResponse, error) {
	req, err := http.NewRequest(request.Method, request.URL, strings.NewReader(request.Body))
	if err != nil {
		return nil, err
	}

	for k, v := range request.Headers {
		req.Header.Set(k, v)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return &models.ExecuteRequestResponse{
		StatusCode: resp.StatusCode,
		Body:       string(respBody),
		Headers:    resp.Header,
	}, nil
}
