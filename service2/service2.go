package main

import (
	"fmt"
	"net/http"
)

// Starts Service2.
// Receives Request2 from Service1.
// Composes response from request1 remote adress and prints it to std.out.
func main() {
	//log.Printf("Service2 starting.")
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8002", nil)
}

func hello(writer http.ResponseWriter, req *http.Request) {
	service2Adrress := req.Host
	request2Adrress := req.RemoteAddr

	fmt.Println("Hello from " + request2Adrress + "\nto " + service2Adrress)
}
