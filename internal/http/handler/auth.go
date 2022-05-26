package handler

import (
	"net/http"
)

type Auth struct {
	privateToken string
}

func NewAuth(privateToken string) *Auth {
	return &Auth{
		privateToken: privateToken,
	}
}

func (a *Auth) Do(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("x-auth-token") != a.privateToken {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		next(w, r)
	}
}
