package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://chenyitian.gitbooks.io/gin-web-framework/content/docs/7.html

func TestGetMessages(t *testing.T) {
	var expectedFileContents string = "2022-11-22T11:01:32.149Z 1 MSG_{1} to compse140.o\n2022-11-22T11:01:35.155Z 2 Got MSG_{2} to compse140.i\n2022-11-22T11:01:38.156Z 3 Got MSG_{3} to compse140.i\n"
	var mockFilecontents string = "2022-11-22T11:01:32.149Z 1 MSG_{1} to compse140.o\n2022-11-22T11:01:35.155Z 2 Got MSG_{2} to compse140.i\n2022-11-22T11:01:38.156Z 3 Got MSG_{3} to compse140.i"

	testServer := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(writer, mockFilecontents)
	}))
	defer testServer.Close()

	mockServerURL := testServer.URL

	// Create test client
	testClient := NewCustomTestClient()
	defer testClient.CloseIdleConnections()

	resp := testClient.GetMessages(mockServerURL)

	assert.Equal(t, expectedFileContents, resp)
}

func TestPutState(t *testing.T) {
	cases := []struct {
		input    []string
		response []string
		expected []string
	}{
		{
			input:    []string{"INIT"},
			response: []string{"ORIG service set to initial state"},
			expected: []string{"ORIG service set to initial state"},
		},
		{
			input:    []string{"PAUSED"},
			response: []string{"ORIG service paused"},
			expected: []string{"ORIG service paused"},
		},
		{
			input:    []string{"RUNNING"},
			response: []string{"ORIG service running"},
			expected: []string{"ORIG service running"},
		},
		{
			input:    []string{"SHUTDOWN"},
			response: []string{"ORIG service shutting down"},
			expected: []string{"ORIG service shutting down"},
		},
	}

	// Create test client
	testClient := NewCustomTestClient()
	defer testClient.CloseIdleConnections()

	// Loop through all inputs and expected data
	for _, testCase := range cases {
		testServer := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, req *http.Request) {
			fmt.Fprintln(writer, testCase.response)
		}))
		defer testServer.Close()

		mockServerURL := testServer.URL
		resp := testClient.PutState(mockServerURL, fmt.Sprint(testCase.input))

		assert.Equal(t, fmt.Sprint(testCase.expected), strings.ReplaceAll(resp, "\n", ""))
	}
}

func TestGetState(t *testing.T) {
	var expectedState string = "PAUSED\n"
	var mockState string = "PAUSED"

	testServer := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(writer, mockState)
	}))
	defer testServer.Close()

	mockServerURL := testServer.URL

	// Create test client
	testClient := NewCustomTestClient()
	defer testClient.CloseIdleConnections()

	resp := testClient.GetMessages(mockServerURL)

	assert.Equal(t, expectedState, resp)
}

func TestGetRunLog(t *testing.T) {
	var expectedFileContents string = "2020-11-01T06:35:01.373Z: INIT\n2020-11-01T06.35:01.380Z: RUNNING\n2020-11-01T06:40:01.373Z: PAUSED\n2020-11-01T06:40:01.373Z: RUNNING\n"
	var mockFilecontents string = "2020-11-01T06:35:01.373Z: INIT\n2020-11-01T06.35:01.380Z: RUNNING\n2020-11-01T06:40:01.373Z: PAUSED\n2020-11-01T06:40:01.373Z: RUNNING"

	testServer := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(writer, mockFilecontents)
	}))
	defer testServer.Close()

	mockServerURL := testServer.URL

	// Create test client
	testClient := NewCustomTestClient()
	defer testClient.CloseIdleConnections()

	resp := testClient.GetMessages(mockServerURL)

	assert.Equal(t, expectedFileContents, resp)
}
