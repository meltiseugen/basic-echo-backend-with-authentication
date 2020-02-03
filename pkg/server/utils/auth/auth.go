package auth

import "github.com/dgrijalva/jwt-go"

var (
	Excluded = []string{"/", "/login", "/logout", "/register", "/v1/message", "/metrics"}
)

type JWTClaims struct {
	Role string `json:"role"`
	jwt.StandardClaims
}
