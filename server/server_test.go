package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"testing"
	"time"
)

func TestServerSuccessfulConnection(t *testing.T) {
	// Start the server in a goroutine
	go main()
	time.Sleep(1 * time.Second) // Give the server time to start

	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		t.Fatalf("Failed to connect to server: %v", err)
	}
	defer conn.Close()

	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)

	challenge, err := reader.ReadString('\n')
	if err != nil {
		t.Fatalf("Failed to read challenge: %v", err)
	}
	challenge = strings.TrimSpace(challenge)

	nonce := calculateValidNonce(challenge)

	_, err = writer.WriteString(nonce + "\n")
	if err != nil {
		t.Fatalf("Failed to send nonce: %v", err)
	}
	writer.Flush()

	response, err := reader.ReadString('\n')
	if err != nil {
		t.Fatalf("Failed to read response: %v", err)
	}
	response = strings.TrimSpace(response)

	for _, quote := range quotes {
		if response == quote {
			return // Test passes
		}
	}

	t.Errorf("Response was not a valid quote: %s", response)
}

func calculateValidNonce(challenge string) string {
	for i := 0; ; i++ {
		nonce := fmt.Sprintf("%d", i)
		if isValidProof(challenge, nonce) {
			return nonce
		}
	}
}
