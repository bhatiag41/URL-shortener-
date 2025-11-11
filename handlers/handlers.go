package handlers

import (
    "encoding/json"
    "net/http"
    "strings"
    "url-shortener/storage"
    "url-shortener/utils"
)

type RequestBody struct {
    URL string `json:"url"`
}

type ResponseBody struct {
    ShortURL string `json:"short_url"`
}

func CreateShortURL(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
        return
    }

    var req RequestBody
    err := json.NewDecoder(r.Body).Decode(&req)
    if err != nil || req.URL == "" {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    if !utils.IsValidURL(req.URL) {
        http.Error(w, "Invalid URL format", http.StatusBadRequest)
        return
    }

    shortCode := utils.GenerateCode(6)
    storage.Save(shortCode, req.URL)

    resp := ResponseBody{
        ShortURL: "http://localhost:8080/" + shortCode,
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(resp)
}

func Redirect(w http.ResponseWriter, r *http.Request) {
    code := strings.TrimPrefix(r.URL.Path, "/")
    if code == "" {
        http.Error(w, "Missing short code", http.StatusBadRequest)
        return
    }

    originalURL := storage.Get(code)
    if originalURL == "" {
        http.Error(w, "URL not found", http.StatusNotFound)
        return
    }

    http.Redirect(w, r, originalURL, http.StatusTemporaryRedirect)
}
package handlers

import (
    "encoding/json"
    "net/http"
    "strings"
    "url-shortener/storage"
    "url-shortener/utils"
)

type RequestBody struct {
    URL string `json:"url"`
}

type ResponseBody struct {
    ShortURL string `json:"short_url"`
}

func CreateShortURL(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
        return
    }

    var req RequestBody
    err := json.NewDecoder(r.Body).Decode(&req)
    if err != nil || req.URL == "" {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    if !utils.IsValidURL(req.URL) {
        http.Error(w, "Invalid URL format", http.StatusBadRequest)
        return
    }

    shortCode := utils.GenerateCode(6)
    storage.Save(shortCode, req.URL)

    resp := ResponseBody{
        ShortURL: "http://localhost:8080/" + shortCode,
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(resp)
}

func Redirect(w http.ResponseWriter, r *http.Request) {
    code := strings.TrimPrefix(r.URL.Path, "/")
    if code == "" {
        http.Error(w, "Missing short code", http.StatusBadRequest)
        return
    }

    originalURL := storage.Get(code)
    if originalURL == "" {
        http.Error(w, "URL not found", http.StatusNotFound)
        return
    }

    http.Redirect(w, r, originalURL, http.StatusTemporaryRedirect)
}
