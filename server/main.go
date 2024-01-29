package main

import (
	"bufio"
	_ "bufio"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	_ "fmt"
	"math/rand"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

func main() {
	_, cancel := context.WithCancel(context.Background())
	defer cancel()
	ln, err := net.Listen("tcp", ":8080")
	if err {
		fmt.Println("Can't start server")
		cancel()
	}

	defer ln.Close()
	go func() {
		for {
			conn, err := ln.Accept()
			if err != nil {
				continue
			}
			go handleConnection(conn)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	<-sigChan
	fmt.Println("\nReceived shutdown signal, closing server...")

	cancel()

	fmt.Println("Server successfully shutdown")
}
func handleConnection(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)

	// Send POW challenge to the client
	_, err := writer.WriteString(Challenge + "\n")
	if err != nil {
		fmt.Println("Error sending challenge:", err)
		return
	}
	writer.Flush()

	// Read client response (nonce)
	nonce, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading nonce:", err)
		return
	}
	nonce = strings.TrimSpace(nonce)

	// Verify the POW
	if isValidProof(Challenge, nonce) {
		// Send a quote if POW is valid
		rand.Seed(time.Now().UnixNano())
		selectedQuote := quotes[rand.Intn(len(quotes))]
		_, err := writer.WriteString(selectedQuote + "\n")
		if err != nil {
			fmt.Println("Error sending quote:", err)
			return
		}
		writer.Flush()
	} else {
		_, err := writer.WriteString("Invalid POW.\n")
		if err != nil {
			fmt.Println("Error sending invalid message:", err)
			return
		}
		writer.Flush()
	}
}

const (
	Challenge  = "SolveThisChallenge"
	Difficulty = 4
)

func isValidProof(challenge, nonce string) bool {
	hash := sha256.Sum256([]byte(challenge + nonce))
	hashStr := hex.EncodeToString(hash[:])
	return strings.HasPrefix(hashStr, strings.Repeat("0", Difficulty))
}

var quotes = []string{
	"Do not dwell in the past, do not dream of the future, concentrate the mind on the present moment.",
	"Life is really simple, but we insist on making it complicated.",
	"The only true wisdom is in knowing you know nothing.",
}
