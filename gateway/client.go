package main

import (
	"crypto/tls"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

// CustomClient is a customized *http.Client
type CustomClient struct {
	*http.Client
}

// NewCustomClient creates a new client with custom settings.
func NewCustomClient() *CustomClient {
	// Skips x509 error
	customTransport := http.DefaultTransport.(*http.Transport).Clone()
	customTransport.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	CustomClient := &CustomClient{&http.Client{Timeout: 120 * time.Second, Transport: customTransport}}

	return CustomClient
}

// NewCustomTestClient creates CustomClient for unit testing.
func NewCustomTestClient() *CustomClient {
	customTransport := http.DefaultTransport.(*http.Transport).Clone()
	customTransport.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	CustomClient := &CustomClient{&http.Client{Timeout: 40 * time.Second, Transport: customTransport}}

	return CustomClient
}

// GetMessages gets messages from url and reads them into a string
func (c CustomClient) GetMessages(url string) string {

	resp, err := c.Get(url)
	if err != nil {
		log.Panic(err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Panic(err)
	}
	return string(respBody)
}

// PutState sends payload to given address. INIT, PAUSED, RUNNING, SHUTDOWN.
func (c CustomClient) PutState(url string, payload string) string {

	req, err := http.NewRequest(http.MethodPut, url, strings.NewReader(payload))
	if err != nil {
		log.Panic(err)
	}

	resp, err := c.Do(req)
	if err != nil {
		log.Panic(err)
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Panic(err)
	}
	return string(respBody)
}
