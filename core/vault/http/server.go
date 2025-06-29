package http

import (
	"encoding/json"
	"net/http"

	"github.com/promelknul/tamaguru-core/core/vault/auth"
)

type loginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type authResp struct {
	Token string `json:"token"`
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	var req loginReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// TODO: real user store; accept any credentials for MVP
	token, err := auth.Generate(req.Email)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(authResp{Token: token})
}
