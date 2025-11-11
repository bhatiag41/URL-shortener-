URL Shortener – Go

A minimal, structured URL shortener built with Go’s standard libraries. The service accepts a long URL and returns a shortened version, then redirects short codes back to the original URL. This project focuses on clean routing, concurrency-safe storage, and basic validation.

Features

Simple REST API (POST /shorten)

Redirect handler (GET /<code>)

In-memory storage with concurrency protection

Random short code generator

URL validation using net/url

Clean folder structure (handlers, storage, utils)

Tech Stack

Go (net/http)

Go routines & sync.RWMutex

No external dependencies

Folder Structure
url-shortener/
│
├── main.go
├── handlers/
│   └── handlers.go
├── storage/
│   └── storage.go
└── utils/
    └── utils.go

API Usage
Create a short URL

Endpoint:
POST http://localhost:8080/shorten

Body:

{
  "url": "https://example.com"
}


Response:

{
  "short_url": "http://localhost:8080/Ab12Xy"
}

Redirect to original URL

Open the generated short URL in your browser:

http://localhost:8080/Ab12Xy


You will be redirected to the original URL.

How to Run
go run main.go handlers/handlers.go storage/storage.go utils/utils.go


Server starts on:

http://localhost:8080

Notes

Storage is in-memory, so data resets on restart.

This project prioritizes simplicity and clarity.

Ideal for beginners learning Go’s HTTP handling and basic concurrency patterns.
