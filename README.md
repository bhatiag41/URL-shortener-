URL Shortener in Go

This is a simple URL shortener built using Go’s standard net/http package. It generates a short code for any valid URL and redirects users back to the original link.

Features

REST endpoint to shorten URLs

Redirect handler for short codes

In-memory storage using a concurrency-safe map

Random short code generation

Basic URL validation

Minimal project structure

Notes

All data is stored in memory and cleared when the server restarts.

This project focuses on basic Go concepts: routing, h
andlers, validation, and concurrency.
