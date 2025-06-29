package main

import (
	"log"
	"net/http"

	vaulthttp "github.com/promelknul/tamaguru-core/core/vault/http"
)

func main() {
	http.HandleFunc("/auth/login", vaulthttp.LoginHandler)
	log.Println("vaultd HTTP :9300 up")
	log.Fatal(http.ListenAndServe(":9300", nil))
}
