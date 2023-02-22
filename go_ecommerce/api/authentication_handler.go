package api

import (
	"encoding/json"
	"log"
	"os"

	"github.com/anthdm/weavebox"
	jwt "github.com/golang-jwt/jwt/v4"
)

type AuthenticationRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthenticationHandler struct {
	// userStore
}

func (h *AuthenticationHandler) AuthenticateUser(ctx *weavebox.Context) error {
	authReq := &AuthenticationRequest{}
	if err := json.NewDecoder(ctx.Request().Body).Decode(authReq); err != nil {
		return err
	}

	os.Setenv("JWT_SECRET", "mysecret")
	token, err := createJWT(authReq.Password)
	if err != nil {
		return err
	}
	log.Println(token)

	return nil
}

func createJWT(pw string) (string, error) {
	claims := &jwt.MapClaims{
		"expiresAt": 15000,
		"password":  pw,
	}

	secret := os.Getenv("JWT_SECRET")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secret))
}
