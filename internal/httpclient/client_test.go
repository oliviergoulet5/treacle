package httpclient

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/oliviergoulet5/treacle/internal/models"
)

func TestClient_ExecutesSuccessfulRequest(t *testing.T) {
	var (
		gotMethod string
		gotHeader string
		gotBody   []byte
	)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotHeader = r.Header.Get("X-Test")
		body, _ := io.ReadAll(r.Body)
		gotBody = body

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`${"ok": true}`))
	}))
	defer server.Close()

	resp, err := Execute(models.ExecuteRequest{
		Method: "POST",
		URL:    server.URL,
		Headers: map[string]string{
			"X-Test": "treacle",
		},
		Body: `{"foo": "bar"}`,
	})

	// Assert error
	if err != nil {
		t.Fatalf("Execute() error = %v", err)
	}

	// Assert request behavior
	if gotMethod != "POST" {
		t.Errorf("expected method POST, got %s", gotMethod)
	}

	if gotHeader != "treacle" {
		t.Errorf("expected header X-Test=treacle, got %s", gotHeader)
	}

	if string(gotBody) == "" {
		t.Errorf("expected request body, got empty")
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status 200, got %d", resp.StatusCode)
	}
}
