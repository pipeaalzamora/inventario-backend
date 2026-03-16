package apiservices

import (
	"fmt"
	"net/http"
	"sofia-backend/domain/facades"
	"sofia-backend/shared"

	"github.com/gin-gonic/gin"

	"strings"
)

func AuthMiddleware(authFacade *facades.AuthFacade, secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := extractToken(c)
		_, claims, err := shared.ParseToken(tokenString, secret)
		if err != nil {
			fmt.Println("Error parsing token:", err)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// puedes pasar el user_id al contexto
		uid, ok := claims[shared.UserIdKey()].(string)
		if ok {
			c.Set(shared.UserIdKey(), uid)
		} else {
			fmt.Println("User ID not found in claims")
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		user, powersPersisted, err := authFacade.GetPersistedUser(c)
		if err != nil {
			fmt.Println("Error getting persisted user:", err)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		powers := make([]string, len(powersPersisted))
		for index, value := range powersPersisted {
			powers[index] = value.PowerName
		}

		// puedes pasar los poderes al contexto
		c.Set(shared.UserPowersKeys(), powers)

		// puedes pasar el user completo al contexto
		c.Set(shared.UserKey(), user)
		c.Next()
	}
}

func extractToken(c *gin.Context) string {
	bearer := c.GetHeader("Authorization")
	parts := strings.Split(bearer, " ")
	if len(parts) == 2 && strings.ToLower(parts[0]) == "bearer" {
		return parts[1]
	}

	return ""
}
