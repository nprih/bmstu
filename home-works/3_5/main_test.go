package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/hello", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(HelloHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("неправильный код ответа: получили %v, ожидали %v", status, http.StatusOK)
	}

	expectedContentType := "application/json"
	if contentType := rr.Header().Get("Content-Type"); contentType != expectedContentType {
		t.Errorf("неправильный Content-Type: получили %v, ожидали %v", contentType, expectedContentType)
	}

	var response map[string]string
	if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
		t.Errorf("не удалось распарсить JSON: %v", err)
	}

	expectedMessage := "hello, world!"
	if message, ok := response["message"]; !ok {
		t.Error("в ответе отсутствует поле 'message'")
	} else if message != expectedMessage {
		t.Errorf("неправильное сообщение: получили %v, ожидали %v", message, expectedMessage)
	}
}
