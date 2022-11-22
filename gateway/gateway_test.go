package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// https://apitest.dev/

// Needs to have httpserv running
func TestGetMessagesAPICall(t *testing.T) {
	mockFilecontents := []byte("2022-11-22T11:01:32.149Z 1 MSG_{1} to compse140.o\n2022-11-22T11:01:35.155Z 2 Got MSG_{2} to compse140.i\n2022-11-22T11:01:38.156Z 3 Got MSG_{3} to compse140.i")
	//mockFilecontents := []byte("2022-11-22T11:01:32.149Z 1 MSG_{1} to compse140.o" + "2022-11-22T11:01:35.155Z 2 Got MSG_{2} to compse140.i" + "2022-11-22T11:01:38.156Z 3 Got MSG_{3} to compse140.i")

	router := gin.Default()
	router.GET("/messages", getMessages)
	req, _ := http.NewRequest("GET", "/messages", nil)
	writer := httptest.NewRecorder()
	router.ServeHTTP(writer, req)

	responseData, _ := ioutil.ReadAll(writer.Body)
	assert.Equal(t, string(mockFilecontents), string(responseData), "Get test file contents from httpserv")
}
