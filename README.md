# GoMailpace

[![GoDoc](https://pkg.go.dev/badge/github.com/mailpace/gomailpace)](https://pkg.go.dev/github.com/mailpace/gomailpace)
[![Build Status](https://circleci.com/gh/mailpace/gomailpace.svg?style=svg)](https://circleci.com/gh/mailpace/gomailpace)
[![Coverage Status](https://codecov.io/gh/mailpace/gomailpace/graph/badge.svg?token=7FP4G7OLY5)](https://codecov.io/gh/mailpace/gomailpace)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)
[![Go Version](https://img.shields.io/github/go-mod/go-version/mailpace/gomailpace)](https://golang.org/doc/go-get-installation)
[![GitHub release](https://img.shields.io/github/release/mailpace/gomailpace.svg)](https://github.com/mailpace/gomailpace/releases)

This Go package provides a client for interacting with the MailPace API for sending emails.

## Installation

To use this package in your Go project, you can import it using the following:

```go
import "github.com/mailpace/gomailpace"
```

## Usage

### Creating a GoMailpace Client

```go
emailClient := gomailpace.NewClient("MAILPACE_TOKEN")
```

Replace "MAILPACE_TOKEN" with your actual MailPace API token.

### Sending an Email

```go
emailPayload := gomailpace.Payload{
    From:     "service@example.com",
    To:       "user@example.com",
    Subject:  "MailPace Rocks!",
    TextBody: "MailPace is the best transactional email provider out there",
}

err := emailClient.Send(emailPayload)
if err != nil {
    // handle err
}
```


### Additional Options

You can include various options such as HTML body, CC, BCC, attachments, tags, etc. as specified in the MailPace API documentation: https://docs.mailpace.com/reference/send 

```go
emailPayload := gomailpace.Payload{
    From:        "service@example.com",
    To:          "user@example.com",
    Subject:     "MailPace Rocks!",
    HTMLBody:    "<html><body><p>Content</p></body></html>",
    CC:          "cc@example.com",
    Attachments: []gomailpace.Attachment{
        {
            Name:        "attachment.jpg",
            Content:     "base64_encoded_string_here",
            ContentType: "image/jpeg",
        },
    },
    Tags: []string{"password reset", "welcome"},
}

```

## Running Tests

To run the unit tests for this package, use the following command:

```bash
go test
```

## Contributions

Feel free to contribute to this project by opening issues or submitting pull requests.

## License

This project is licensed under the MIT License - see the LICENSE file for details