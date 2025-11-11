package main

import (
    "log"
    "net/http"
    "url-shortener/handlers"
)

func main() {
    mux := http.NewServeMux()

    mux.HandleFunc("/shorten", handlers.CreateShortURL)
    mux.HandleFunc("/", handlers.Redirect)

    log.Println("Server running on :8080")
    log.Fatal(http.ListenAndServe(":8080", mux))
}
