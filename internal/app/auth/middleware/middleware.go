package auth_middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/sri2103/domain_DD_todo/internal/app/auth"
)

func AuthMiddleware(next http.Handler)http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		authHeader := r.Header.Get("Authorization")

		if authHeader ==  "" {
			//TODO: return error
			return
		}
		tokenParts := strings.Split(authHeader," ")

		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			// TODO handler unauthorized
			return
		}
		tokenString := tokenParts[1]

		claimsData,err := auth.VerifyToken(tokenString)

		if err != nil {
			// TODO Handle error verification
			return
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, "claimsData", claimsData)
		r = r.WithContext(ctx)

		next.ServeHTTP(w,r)
	})
}