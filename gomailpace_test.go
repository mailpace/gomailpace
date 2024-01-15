package gomailpace

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSendSuccess(t *testing.T) {
	// Create a fake Mailpace API server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check request method and path
		if r.Method != "POST" || r.URL.Path != "/send" {
			t.Errorf("Expected POST request to /send, got %s request to %s", r.Method, r.URL.Path)
		}

		// Read request body
		var payload Payload
		err := json.NewDecoder(r.Body).Decode(&payload)
		if err != nil {
			t.Errorf("Failed to decode request body: %v", err)
		}

		// Check some fields in the payload
		if payload.From != "service@example.com" {
			t.Errorf("Expected 'From' field to be 'service@example.com', got %s", payload.From)
		}

		// Respond with a success status
		w.WriteHeader(http.StatusOK)
	}))

	// Close the server when done
	defer server.Close()

	// Create a Mailpace client with the fake server URL
	client := NewClient("fake_token", server.URL+"/send")

	// Create a payload for testing
	payload := Payload{
		From:     "service@example.com",
		To:       "user@example.com",
		Subject:  "Test",
		TextBody: "Hello, this is a test!",
	}

	// Send the email using the client
	err := client.Send(context.Background(), payload)
	if err != nil {
		t.Errorf("Failed to send email: %v", err)
	}
}

func TestSendFailure(t *testing.T) {
	// Create a fake Mailpace API server that returns an error
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Respond with an error status
		w.WriteHeader(http.StatusInternalServerError)
	}))

	// Close the server when done
	defer server.Close()

	// Create a Mailpace client with the fake server URL
	client := NewClient("fake_token", server.URL+"/send")

	// Create a payload for testing
	payload := Payload{
		From:     "service@example.com",
		To:       "user@example.com",
		Subject:  "Test",
		TextBody: "Hello, this is a test!",
	}

	// Send the email using the client
	err := client.Send(context.Background(), payload)
	if err == nil {
		t.Error("Expected an error, but got nil")
	}
}

func TestNewClientWithoutURL(t *testing.T) {
	token := "your_token"

	client := NewClient(token)

	// Check if the token is set correctly
	if client.Token != token {
		t.Errorf("Expected token %s, got %s", token, client.Token)
	}

	// Check if the URL is set to the default value
	expectedURL := "https://app.mailpace.com/api/v1/send"
	if client.URL != expectedURL {
		t.Errorf("Expected URL %s, got %s", expectedURL, client.URL)
	}
}

func TestSendWithTimeout(t *testing.T) {
	// Create a fake Mailpace API server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check request method and path
		if r.Method != "POST" || r.URL.Path != "/send" {
			t.Errorf("Expected POST request to /send, got %s request to %s", r.Method, r.URL.Path)
		}

		// Read request body
		var payload Payload
		err := json.NewDecoder(r.Body).Decode(&payload)
		if err != nil {
			t.Errorf("Failed to decode request body: %v", err)
		}

		// Check some fields in the payload
		if payload.From != "service@example.com" {
			t.Errorf("Expected 'From' field to be 'service@example.com', got %s", payload.From)
		}

		// Respond with a success status
		w.WriteHeader(http.StatusOK)
	}))

	// Close the server when done
	defer server.Close()

	// Create a Mailpace client with the fake server URL
	client := NewClient("fake_token", server.URL+"/send")

	// Create a payload for testing
	payload := Payload{
		From:     "service@example.com",
		To:       "user@example.com",
		Subject:  "Test",
		TextBody: "Hello, this is a test!",
	}

	// Create a context with a timeout of 1 nanosecond
	ctx, cancel := context.WithTimeout(context.Background(), 1)
	defer cancel()

	// Send the email using the client
	err := client.Send(ctx, payload)
	// Check if an error was returned
	if err == nil {
		t.Errorf("expected an error, got nil")
	} else if !errors.Is(err, context.DeadlineExceeded) {
		t.Errorf("expected a DeadlineExceeded error, got %v", err)
	}
}
