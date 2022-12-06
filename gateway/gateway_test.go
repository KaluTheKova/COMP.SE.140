package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
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
	testServer := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(writer, "Received PUT")
	}))
	defer testServer.Close()

	mockServerURL := testServer.URL

	// Create test client
	testClient := NewCustomTestClient()
	defer testClient.CloseIdleConnections()

	resp := testClient.PutState(mockServerURL, "INIT")

	log.Println("DEBUG:", resp)

}

/* func TestGetMessagesFromHttpserv(t *testing.T) {
	mockResponse := []byte("2022-11-22T11:01:32.149Z 1 MSG_{1} to compse140.o\n2022-11-22T11:01:35.155Z 2 Got MSG_{2} to compse140.i\n2022-11-22T11:01:38.156Z 3 Got MSG_{3} to compse140.i")

	// Create a response recorder so you can inspect the response
	writer := httptest.NewRecorder()

	// Router
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	// Create the mock request you'd like to test. Make sure the second argument
	// here is the same as one of the routes you defined in the router setup block!
	req, _ := http.NewRequest("GET", "/messages", nil)
	router.ServeHTTP(writer, req)

	assert.Equal(t, 200, writer.Code)
	assert.Equal(t, string(mockResponse), writer.Body.String())
} */

// Needs to have httpserv running
// func TestGetMessages(t *testing.T) {
// 	mockFilecontents := []byte("2022-11-22T11:01:32.149Z 1 MSG_{1} to compse140.o\n2022-11-22T11:01:35.155Z 2 Got MSG_{2} to compse140.i\n2022-11-22T11:01:38.156Z 3 Got MSG_{3} to compse140.i")
// 	//mockFilecontents := []byte("2022-11-22T11:01:32.149Z 1 MSG_{1} to compse140.o" + "2022-11-22T11:01:35.155Z 2 Got MSG_{2} to compse140.i" + "2022-11-22T11:01:38.156Z 3 Got MSG_{3} to compse140.i")

// 	router := gin.Default()
// 	router.GET("/messages", getMessages)
// 	req, _ := http.NewRequest("GET", "/messages", nil)
// 	writer := httptest.NewRecorder()
// 	router.ServeHTTP(writer, req)

// 	responseData, _ := ioutil.ReadAll(writer.Body)
// 	assert.Equal(t, string(mockFilecontents), string(responseData), "Get test file contents from httpserv")
// }

// Testissä Lähetä request curl -v -X GET localhost:8083/messages
