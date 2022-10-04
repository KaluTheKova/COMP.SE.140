package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var service2Address string = os.Getenv("SERVICE2_ADDRESS")

// Starts Service1.
// Receives Request1 from outside the docker network.
// Sends Request2 to Service2.
// Composes response from request1 remote adress and prints it to std.out.
func main() {
	//log.Printf("Service1 starting at 8001. Access me via \"curl localhost:8001\"")
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8001", nil)
}

func hello(writer http.ResponseWriter, req *http.Request) {
	service1Adrress := req.Host
	request1Adrress := req.RemoteAddr

	fmt.Println("Hello from " + request1Adrress + "\nto " + service1Adrress)

	_, err := http.Get(service2Address)
	if err != nil {
		log.Fatal(err)
	}
}
