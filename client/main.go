package main

import (
	"bufio"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

const (
	ServerAddress = "localhost:8080"
	Difficulty    = 4
)

func solvePOW(challenge string) string {
	var nonce int
	for {
		testNonce := fmt.Sprintf("%x", nonce)
		hash := sha256.Sum256([]byte(challenge + testNonce))
		hashStr := hex.EncodeToString(hash[:])

		if strings.HasPrefix(hashStr, strings.Repeat("0", Difficulty)) {
			return testNonce
		}
		nonce++
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		quote, err := StartClientAndReceiveQuote(ctx, ServerAddress)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println("Received from server:", quote)
	}()
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	<-sigChan
	fmt.Println("\nReceived shutdown signal, closing client...")

	cancel()

	fmt.Println("Client successfully shutdown")
}
func StartClientAndReceiveQuote(ctx context.Context, serverAddr string) (string, error) {
	conn, err := net.Dial("tcp", serverAddr)
	if err != nil {
		return "", fmt.Errorf("failed to connect to server: %w", err)
	}
	defer conn.Close()

	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)

	// Receive POW challenge from server
	challenge, err := reader.ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("failed to read challenge: %w", err)
	}
	challenge = strings.TrimSpace(challenge)

	nonce := solvePOW(challenge)

	_, err = writer.WriteString(nonce + "\n")
	if err != nil {
		return "", fmt.Errorf("failed to send nonce: %w", err)
	}
	writer.Flush()

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	default:
		quote, err := reader.ReadString('\n')
		if err != nil {
			return "", fmt.Errorf("failed to read quote: %w", err)
		}
		return strings.TrimSpace(quote), nil
	}
}
