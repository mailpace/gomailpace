package gomailpace

import (
	"encoding/json"
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
	err := client.Send(payload)
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
	err := client.Send(payload)
	if err == nil {
		t.Error("Expected an error, but got nil")
	}
}
