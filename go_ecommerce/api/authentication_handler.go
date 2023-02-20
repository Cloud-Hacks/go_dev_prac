package api

import (
	"encoding/json"
	"os"

	"github.com/anthdm/weavebox"
)

type AuthenticationRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthenticationHandler struct {
	// userStore
}

func (h *AuthenticationHandler) AuthenticateUser(ctx *weavebox.Context) (error, string) {
	authReq := &AuthenticationRequest{}
	if err := json.NewDecoder(ctx.Request().Body).Decode(authReq); err != nil {
		return err, ""
	}

	token, err := createJWT(authReq.Password)
	if err != nil {
		return err, ""
	}

	return nil, token
}

func createJWT(account *Account) (string, error) {
	claims := &jwt.MapClaims{
		"expiresAt":     15000,
		"accountNumber": account.Number,
	}

	secret := os.Getenv("JWT_SECRET")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secret))
}
