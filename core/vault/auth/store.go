package auth

import "sync"

// naive in‑mem store for MVP
var mu sync.Mutex
var refreshMap = map[string]string{} // refresh -> userID

func SaveRefresh(user, refresh string) {
	mu.Lock()
	defer mu.Unlock()
	refreshMap[refresh] = user
}
func ValidateRefresh(r string) (string, bool) {
	mu.Lock()
	defer mu.Unlock()
	uid, ok := refreshMap[r]
	return uid, ok
}
