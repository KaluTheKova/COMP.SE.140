package main

import (
	"log"
	"time"
)

func main() {
	timestamp := time.Now().Format("2006-01-02T15:04:05.999Z")
	log.Printf(timestamp)
}
