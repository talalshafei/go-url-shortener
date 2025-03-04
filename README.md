# URL Shortener Generator

This project is a lightweight URL shortener built using Go and the Gin web framework. It was developed by following a comprehensive tutorial series that walks you through every component—from generating a short URL hash to handling HTTP redirection. By integrating previously built modules for short URL generation and storage, the final product allows users to submit a long URL via a `/create-short-url` endpoint and then redirect to the original URL when the generated short URL is accessed.

## Key Features

- **Short URL Creation:** Generates a compact, hash-based URL when provided with a long URL and a user ID.
- **Redirection:** Efficiently maps the short URL back to the original long URL for seamless navigation.
- **Modular Design:** Built by integrating separate components for URL generation and storage, making the project scalable and easy to maintain.

## Project Structure

```plaintext
├── go.mod
├── go.sum
├── handler
│   └── handlers.go
├── main.go
├── shortener
│   ├── shorturl_generator.go
│   └── shorturl_generator_test.go
└── store
    ├── store_service.go
    └── store_service_test.go
