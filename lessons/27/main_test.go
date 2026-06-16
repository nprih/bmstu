package main

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/ozontech/cute"
	"github.com/ozontech/cute/asserts/json"
)

func TestStatusHandler(t *testing.T) {
	cute.NewTestBuilder().
		Title("Simple test 1").
		Description("This is simple test for 200 ok and json").
		Create().
		RequestBuilder(
			cute.WithURI("http://localhost:8080/status"),
			cute.WithMethod(http.MethodGet),
		).
		ExpectExecuteTimeout(5*time.Second).
		ExpectStatus(http.StatusOK).
		AssertBody(
			json.Equal("$.status", "ok"),
		).
		ExecuteTest(context.Background(), t)
}

func TestIndexPageNotFound(t *testing.T) {
	cute.NewTestBuilder().
		Title("Simple test 2").
		Description("This is simple test for 404 on Index Page").
		Create().
		RequestBuilder(
			cute.WithURI("http://localhost:8080/"),
			cute.WithMethod(http.MethodGet),
		).
		ExpectExecuteTimeout(2*time.Second).
		ExpectStatus(http.StatusNotFound).
		ExecuteTest(context.Background(), t)
}
