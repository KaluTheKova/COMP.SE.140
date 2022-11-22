package main

import (
	"net/http"
	"testing"

	"github.com/steinfletcher/apitest"
)

// https://apitest.dev/

func testGetMessagesAPICall(t *testing.T) {
	var messagesArray [3]string
	messagesArray[0] = "2022-11-11T18:01:38.011Z 1 MSG_1 to compse140.i"
	messagesArray[1] = "2022-11-11T18:11:58.08Z 1 MSG_1 to compse140.i"
	messagesArray[2] = "2022-11-11T18:12:39.646Z 1 MSG_1 to compse140.i"

	messages := []byte("2022-11-11T18:01:38.011Z 1 MSG_1 to compse140.i\n2022-11-11T18:11:58.08Z 1 MSG_1 to compse140.i\n2022-11-11T18:12:39.646Z 1 MSG_1 to compse140.i")

	handler := func(w http.ResponseWriter, r *http.Request) {
		msg := `{"message": "hello"}`
		_, _ = w.Write([]byte(msg))
		w.WriteHeader(http.StatusOK)
	}

	// Test API gateway's response to calls to it
	// Example: Get /messages returns mock messages
	apitest.New(). // configuration
			HandlerFunc(GetMessagesAPICall()).
			Get("/message"). // request
			Expect(t).       // expectations
			Body(`{"message": "hello"}`).
			Status(http.StatusOK).
			End()
}
