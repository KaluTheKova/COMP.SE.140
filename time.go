package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

func main() {

	//clearFileOnStartup("messages.txt")

	listAllFilesInDirectory("/")

	//message := buildTimeStampedMessage("MSG_1", 1, "compse140.i")

	//writeToFile("messages.txt", message)

}

// Write listened messages to file
func writeToFile(filename string, message string) error {
	file, err := os.OpenFile("messages.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(message + "\n")
	if err != nil {
		return err
	}

	// Flush writer
	file.Sync()

	return nil
}

func buildTimeStampedMessage(message string, counter int, topic string) string {
	timestamp := time.Now().Format("2006-01-02T15:04:05.999Z")
	timeStampedMessage := fmt.Sprintf("%v %v %v to %v", timestamp, counter, message, topic)

	return timeStampedMessage
}

// Removes filename
func clearFileOnStartup(filename string) {
	err := os.Remove(filename)
	if err != nil {
		log.Println(err)
	}
}

func listAllFilesInDirectory(path string) {
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return err
		}
		fmt.Printf("dir: %v: name: %s\n", info.IsDir(), path)
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}
}
