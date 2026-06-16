package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStatusHandler(t *testing.T) {
	type want struct {
		code        int
		response    string
		contentType string
	}
	tests := []struct {
		name string
		want want
	}{
		{
			name: "Positive test 1 (200 OK)",
			want: want{
				code:        200,
				response:    `{"status": "ok"}`,
				contentType: "application/json",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodGet, "/status", nil)
			w := httptest.NewRecorder()
			StatusHandler(w, request)
			res := w.Result()
			assert.Equal(t, tt.want.code, res.StatusCode)
		})
	}
}
