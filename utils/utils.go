package utils

import (
    "math/rand"
    "net/url"
    "time"
)

func GenerateCode(length int) string {
    chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
    rand.Seed(time.Now().UnixNano())

    result := make([]byte, length)
    for i := range result {
        result[i] = chars[rand.Intn(len(chars))]
    }
    return string(result)
}

func IsValidURL(raw string) bool {
    _, err := url.ParseRequestURI(raw)
    return err == nil
}
