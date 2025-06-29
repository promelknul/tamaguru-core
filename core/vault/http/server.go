package http

import (
	"encoding/json"
	"net/http"

	"github.com/promelknul/tamaguru-core/core/vault/auth"
)

func cors(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
}

type loginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type authResp struct {
	Token string `json:"token"`
	Refresh string `json:"refresh_token"`
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	cors(w)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var req loginReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "bad json", http.StatusBadRequest)
		return
	}
	token, err := auth.Generate(req.Email)
	if err != nil {
		http.Error(w, "token error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	refresh:=auth.GenRefresh();auth.SaveRefresh(req.Email,refresh);json.NewEncoder(w).Encode(authResp{Token:token,Refresh:refresh})
}
