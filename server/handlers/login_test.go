package handlers_test

import (
	"bytes"

	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Broklam/wg/handlers"
)

type Response struct {
	Expiry string `json:"expiry"`
	Status string `json:"status"`
	Token  string `json:"token"`
}

func TestLoginHandler(t *testing.T) {
	requestBody := map[string]string{
		"login":    "1",
		"password": "1",
	}
	body, err := json.Marshal(requestBody)
	if err != nil {
		t.Fatalf("Error marshaling request body: %v", err)
	}
	req, err := http.NewRequest("POST", "/login", bytes.NewReader(body))
	if err != nil {
		t.Fatalf("Error creating request: %v", err)
	}
	rr := httptest.NewRecorder()

	handlers.LoginHandler(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, rr.Code)
	}

	var respBody Response
	if err := json.Unmarshal(rr.Body.Bytes(), &respBody); err != nil {
		t.Fatalf("Error unmarshaling response body: %v", err)
	}

	if respBody.Expiry == "" || respBody.Status == "" || respBody.Token == "" {
		t.Errorf("Response fields are empty")
	}

}
