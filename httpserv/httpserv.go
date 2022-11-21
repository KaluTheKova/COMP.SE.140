package main

import (
	"log"
	"net/http"
	"os"
)

var filename string = "messages.txt"

// When requested, returns content of the file created by OBSE
func main() {
	log.Printf("httpserv starting")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8081", nil)
}

// Responds to HTTP GET <host>:8080 with readFile
func handler(writer http.ResponseWriter, req *http.Request) {
	fileContents := readFile()

	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/text")
	writer.Write([]byte(fileContents))
	return
}

// readFile reads file written by OBSE
func readFile() []byte {
	fileContents, err := os.ReadFile(filename)
	if err != nil {
		log.Println(err)
	}
	return fileContents
}
