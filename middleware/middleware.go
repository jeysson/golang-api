package middleware

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/didip/tollbooth"
	"github.com/didip/tollbooth/limiter"
	"github.com/gorilla/mux"
)

func JWTMiddleware(secretKey string) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			tokenString := r.Header.Get("Authorization")
			if !strings.HasPrefix(tokenString, "Bearer ") {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			tokenString = strings.TrimPrefix(tokenString, "Bearer ")

			claims := &jwt.StandardClaims{}
			token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
				return []byte(secretKey), nil
			})

			if err != nil || !token.Valid {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

func RateLimiterMiddleware(next http.Handler) http.Handler {
	lim := tollbooth.NewLimiter(1, &limiter.ExpirableOptions{
		DefaultExpirationTTL: time.Minute,
	})

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		httpError := tollbooth.LimitByRequest(lim, w, r)
		if httpError != nil {
			http.Error(w, httpError.Message, httpError.StatusCode)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func LogginMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Request: %s %s\n", r.Method, r.RequestURI)
		next.ServeHTTP(w, r)
	})
}
