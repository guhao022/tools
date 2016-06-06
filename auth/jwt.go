package auth

import (
	jwt "github.com/dgrijalva/jwt-go"
	"net/http"
)

func RequireTokenAuthentication(rw http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
	authBackend := InitJWTAuthenticationBackend()

	token, err := jwt.ParseFromRequest(
		req,
		func(token *jwt.Token) (interface{}, error) {
			return authBackend.PublicKey, nil
		})

	if err == nil &&
	token.Valid &&
	!authBackend.IsInBlacklist(req.Header.Get("Authorization")) {
		next(rw, req)
	} else {
		rw.WriteHeader(http.StatusUnauthorized)
	}
}