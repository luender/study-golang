package middleware

import (
	"api-project/src/internal/auth"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware() gin.HandlerFunc {
    return func(ctx *gin.Context) {
        tokenString := ctx.GetHeader("Authorization")
        if tokenString == "" {
            ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Token não fornecido"})
            ctx.Abort()
            return
        }

        tokenString = strings.TrimPrefix(tokenString, "Bearer ")

        userID, err := auth.VerifyJwtToken(tokenString)
        if err != nil {
            ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
            ctx.Abort()
            return
        }

        ctx.Set("userID", userID)
        ctx.Next()
    }
}
