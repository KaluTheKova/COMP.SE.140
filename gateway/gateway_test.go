package main

import (
	"net/http"
	"testing"

	"github.com/steinfletcher/apitest"
)

// https://apitest.dev/

func testGetMessages(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		msg := `{"message": "hello"}`
		_, _ = w.Write([]byte(msg))
		w.WriteHeader(http.StatusOK)
	}

	// Test API gateway's response to calls to it
	// Example: Get /messages returns mock messages
	apitest.New(). // configuration
			HandlerFunc(handler).
			Get("/message"). // request
			Expect(t).       // expectations
			Body(`{"message": "hello"}`).
			Status(http.StatusOK).
			End()
}
