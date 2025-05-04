package middlewares

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

var userCtxKey = &contextKey{"user"}

type contextKey struct {
	name string
}

type User struct {
	Writer     http.ResponseWriter
	UserName string `json:"username"`
}

func MiddleWares() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			c, err := r.Cookie("auth-cookie")
			if err != nil || c == nil {
				ctx := context.WithValue(r.Context(), userCtxKey, &User{
					UserName: "",
					Writer:   w,
				})
				r = r.WithContext(ctx)
				next.ServeHTTP(w, r)
				return
			}
			tokenString := c.Value

			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				}

				return []byte(os.Getenv("SECRET")), nil
			})

			claims,ok := token.Claims.(jwt.MapClaims)
			if !ok {
				fmt.Printf("error")
				http.Error(w, "Invalid cookie", http.StatusForbidden)
				return
			}
			userName := claims["username"].(string)
			ctx := context.WithValue(r.Context(), userCtxKey, &User{
				UserName: userName,
				Writer:   w,
			})
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

func ForContext(ctx context.Context) *User {
	raw, _ := ctx.Value(userCtxKey).(*User)
	return raw
}