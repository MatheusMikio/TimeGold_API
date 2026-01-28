package middlewares

import (
	"net/http"
	"strings"

	"github.com/MatheusMikio/config"
	"github.com/MatheusMikio/handlers"
	"github.com/MatheusMikio/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthRequired() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			handlers.SendError(ctx, http.StatusUnauthorized, models.CreateErrorMessage("Authorization", "Token not provided"))
			ctx.Abort()
			return
		}

		parts := strings.Fields(authHeader)
		if len(parts) != 2 || parts[0] != "Bearer" {
			handlers.SendError(ctx, http.StatusUnauthorized, models.CreateErrorMessage("Authorization", "Invalid token format"))
			ctx.Abort()
			return
		}

		tokenString := parts[1]

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return []byte(config.GetJwtSecret()), nil
		})

		if err != nil || !token.Valid {
			handlers.SendError(ctx, http.StatusUnauthorized, models.CreateErrorMessage("Token", "Invalid or expired token"))
			ctx.Abort()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			ctx.Set("user_id", uint(claims["user_id"].(float64)))
			ctx.Set("email", claims["email"].(string))
			ctx.Set("entity_type", claims["entity_type"].(string))
			ctx.Set("role", claims["role"].(string))
		}

		ctx.Next()
	}
}

func RoleRequired(allowedRoles ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		logger := config.GetLogger("AuthMiddleware")
		logger.Infof("AuthRequired called for %s", ctx.Request.URL.Path)
		role, exists := ctx.Get("role")
		if !exists {
			handlers.SendError(ctx, http.StatusForbidden, models.CreateErrorMessage("Authorization", "Role not found"))
			ctx.Abort()
			return
		}

		roleStr := role.(string)
		for _, allowed := range allowedRoles {
			if roleStr == allowed {
				ctx.Next()
				return
			}
		}

		handlers.SendError(ctx, http.StatusForbidden, models.CreateErrorMessage("Authorization", "Access denied for this role"))
		ctx.Abort()
	}

}
