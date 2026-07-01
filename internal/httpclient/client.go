package httpclient

import (
	"io"
	"net/http"

	"github.com/oliviergoulet5/treacle/internal/models"
)

func Execute(method, url string, headers map[string]string) (*models.ExecuteRequestResponse, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return &models.ExecuteRequestResponse{
		StatusCode: resp.StatusCode,
		Body:       string(body),
		Headers:    resp.Header,
	}, nil
}
