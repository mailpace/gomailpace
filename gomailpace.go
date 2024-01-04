package gomailpace

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// Attachment represents an email attachment.
type Attachment struct {
	Name        string `json:"name"`
	Content     string `json:"content"`
	ContentType string `json:"content_type"`
	CID         string `json:"cid,omitempty"`
}

// Payload represents the email payload with additional options.
type Payload struct {
	From            string       `json:"from"`
	To              string       `json:"to"`
	HTMLBody        string       `json:"htmlbody,omitempty"`
	TextBody        string       `json:"textbody,omitempty"`
	CC              string       `json:"cc,omitempty"`
	BCC             string       `json:"bcc,omitempty"`
	Subject         string       `json:"subject,omitempty"`
	ReplyTo         string       `json:"replyto,omitempty"`
	ListUnsubscribe string       `json:"list_unsubscribe,omitempty"`
	Attachments     []Attachment `json:"attachments,omitempty"`
	Tags            interface{}  `json:"tags,omitempty"`
}

// Client represents the MailPace API client.
type Client struct {
	Token string
	URL string
}

const DefaultURL = "https://app.mailpace.com/api/v1/send"

// NewClient creates a new MailPace API client.
func NewClient(token string, urls ...string) *Client {
	client := &Client{
		Token: token,
	}

	if len(urls) > 0 && urls[0] != "" {
		client.URL = urls[0]
	} else {
		client.URL = DefaultURL
	}

	return client
}

// Send sends an email using the MailPace API.
func (c *Client) Send(payload Payload) error {
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", c.URL, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("MailPace-Server-Token", c.Token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to send email. Status code: %d", resp.StatusCode)
	}

	return nil
}
