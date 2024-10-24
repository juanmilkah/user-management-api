package middleware

import (
    "context"
    "net/http"
    "strings"
    "go-server/utils"
    "github.com/golang-jwt/jwt/v5"
)

type AuthMiddleware struct {
    jwtSecret    string
    errorHandler *utils.ErrorHandler
}

func NewAuthMiddleware(jwtSecret string, errorHandler *utils.ErrorHandler) *AuthMiddleware {
    return &AuthMiddleware{
        jwtSecret:    jwtSecret,
        errorHandler: errorHandler,
    }
}

func (m *AuthMiddleware) RequireAuth(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        authHeader := r.Header.Get("Authorization")
        if authHeader == "" {
            m.errorHandler.RespondWithError(w, http.StatusUnauthorized, "No authorization header", nil)
            return
        }

        tokenString := strings.Replace(authHeader, "Bearer ", "", 1)
        token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
            return []byte(m.jwtSecret), nil
        })

        if err != nil || !token.Valid {
            m.errorHandler.RespondWithError(w, http.StatusUnauthorized, "Invalid token", nil)
            return
        }

        if claims, ok := token.Claims.(jwt.MapClaims); ok {
            ctx := context.WithValue(r.Context(), "userID", claims["user_id"])
            next.ServeHTTP(w, r.WithContext(ctx))
        } else {
            m.errorHandler.RespondWithError(w, http.StatusUnauthorized, "Invalid token claims", nil)
        }
    })
}
