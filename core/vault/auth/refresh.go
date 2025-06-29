package auth

import (
	"crypto/rand"
	"encoding/hex"
)

func GenRefresh() string {
	b := make([]byte, 32)
	rand.Read(b)
	return hex.EncodeToString(b)
}
