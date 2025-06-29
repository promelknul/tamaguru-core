package http

import (
	"encoding/json"
	"net/http"

	"github.com/promelknul/tamaguru-core/core/vault/auth"
)

type refreshReq struct {
	Refresh string `json:"refresh_token"`
}
func RefreshHandler(w http.ResponseWriter, r *http.Request) {
	cors(w)
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var req refreshReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "bad json", http.StatusBadRequest)
		return
	}
	uid, ok := auth.ValidateRefresh(req.Refresh)
	if !ok {
		http.Error(w, "invalid token", http.StatusUnauthorized)
		return
	}
	newToken, _ := auth.Generate(uid)
	json.NewEncoder(w).Encode(map[string]string{"token": newToken})
}
