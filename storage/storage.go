package storage

import "sync"

var urlMap = make(map[string]string)
var mu sync.RWMutex

func Save(code, url string) {
    mu.Lock()
    defer mu.Unlock()
    urlMap[code] = url
}

func Get(code string) string {
    mu.RLock()
    defer mu.RUnlock()
    return urlMap[code]
}
