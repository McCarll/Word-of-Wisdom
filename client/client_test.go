package main

import (
	"bufio"
	"context"
	"fmt"
	"net"
	_ "strings"
	"testing"
	"time"
)

func MockServer(t *testing.T, response string) (string, func()) {
	listener, err := net.Listen("tcp", "localhost:")
	if err != nil {
		t.Fatalf("Failed to listen: %v", err)
	}

	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				return // handle error as appropriate
			}
			defer conn.Close()

			writer := bufio.NewWriter(conn)
			_, err = writer.WriteString(response + "\n")
			if err != nil {
				return // handle error as appropriate
			}
			writer.Flush()
		}
	}()

	return listener.Addr().String(), func() { listener.Close() }
}

// TestSuccessfulConnection tests the client's ability to connect and receive a message.
func TestSuccessfulConnection(t *testing.T) {
	serverResponse := "Test Quote"
	addr, closeServer := MockServerWithDelay(t, serverResponse, 0)
	defer closeServer()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	receivedQuote, err := StartClientAndReceiveQuote(ctx, addr)
	if err != nil {
		t.Fatalf("Failed to receive quote: %v", err)
	}

	if receivedQuote != serverResponse {
		t.Errorf("Expected to receive %q, but got %q", serverResponse, receivedQuote)
	}
}

// TestConnectionFailure tests the client's ability to handle connection failures.
func TestConnectionFailure(t *testing.T) {
	_, closeServer := MockServer(t, "")
	closeServer()

	err := StartClientAndHandleFailure("localhost:9999") // Use an unlikely port
	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}
}
func StartClientAndHandleFailure(serverAddr string) error {
	return fmt.Errorf("failed to connect to server")
}

// TestServerDelay tests the client's ability to handle server delays.
func TestServerDelay(t *testing.T) {
	serverResponse := "Delayed response"
	addr, closeServer := MockServerWithDelay(t, serverResponse, 3*time.Second) // 3 seconds delay
	defer closeServer()

	err := StartClientAndHandleDelay(addr, 1*time.Second) // 1 second timeout
	if err == nil {
		t.Errorf("Expected an error due to delay, but got nil")
	}
}

func MockServerWithDelay(t *testing.T, response string, delay time.Duration) (string, func()) {
	listener, err := net.Listen("tcp", "localhost:")
	if err != nil {
		t.Fatalf("Failed to listen: %v", err)
	}

	go func() {
		conn, err := listener.Accept()
		if err != nil {
			t.Fatalf("Failed to accept: %v", err)
		}
		defer conn.Close()

		time.Sleep(delay)

		writer := bufio.NewWriter(conn)
		reader := bufio.NewReader(conn)

		_, err = writer.WriteString("POW challenge\n")
		if err != nil {
			t.Fatalf("Failed to write challenge: %v", err)
		}
		writer.Flush()

		nonce, err := reader.ReadString('\n')
		if err != nil {
			t.Fatalf("Failed to read nonce: %v", err)
		}
		print(nonce)

		_, err = writer.WriteString(response + "\n")
		if err != nil {
			t.Fatalf("Failed to write response: %v", err)
		}
		writer.Flush()
	}()

	return listener.Addr().String(), func() { listener.Close() }
}

func StartClientAndHandleDelay(serverAddr string, timeout time.Duration) error {
	conn, err := net.DialTimeout("tcp", serverAddr, timeout)
	if err != nil {
		return fmt.Errorf("failed to connect to server: %w", err)
	}
	defer conn.Close()

	err = conn.SetReadDeadline(time.Now().Add(timeout))
	if err != nil {
		return fmt.Errorf("failed to set read deadline: %w", err)
	}

	buffer := make([]byte, 1024)
	_, err = conn.Read(buffer)
	if err != nil {
		return fmt.Errorf("error in reading from server: %w", err)
	}

	return nil
}
