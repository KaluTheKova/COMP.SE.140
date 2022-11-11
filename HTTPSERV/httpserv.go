package main

import (
	"log"
	"net/http"
)

// When requested, returns content of the file created by OBSE
func main() {
	log.Printf("HTTPSERV starting")
	http.HandleFunc("/", get)
	http.ListenAndServe(":8080", nil)
}

// readFile reads file written by OBSE
func readFile() string {
	// 1. When receiving HTTP GET <host>:8080
	// 2. Open file (or current copy of it. Possibly mounted/shared folder with OBSE?)
	// 3. Read contents of file.
	// 4. Return contents of file.
	fileContents := "lol"
	return fileContents
}

// Responds to HTTP GET <host>:8080 with readFile
func get(writer http.ResponseWriter, req *http.Request) {

	response := readFile()

	_, err := http.Get(rabbitMQAddress)
	if err != nil {
		log.Fatal(err)
	}
}
