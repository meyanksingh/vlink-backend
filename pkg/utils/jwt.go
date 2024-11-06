package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte(os.Getenv("JWT_SECRET"))

func GenerateToken(username string) (string, error) {
	claims := jwt.MapClaims{
		"authorized": true,
		"user_id":    username,
		"exp":        time.Now().Add(24 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

func ValidateJWT(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrTokenMalformed
		}
		return secretKey, nil
	})
}

// func JWTAuthMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		// Retrieve the token from the Authorization header
// 		tokenString := c.GetHeader("Authorization")
// 		if tokenString == "" {
// 			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token required"})
// 			c.Abort()
// 			return
// 		}

// 		// Remove "Bearer " prefix from token if it's present
// 		if strings.HasPrefix(tokenString, "Bearer ") {
// 			tokenString = tokenString[len("Bearer "):]
// 		}

// 		// Validate the token
// 		token, err := ValidateJWT(tokenString)
// 		if err != nil || !token.Valid {
// 			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
// 			c.Abort()
// 			return
// 		}

// 		// Extract claims and add them to the context
// 		claims, ok := token.Claims.(jwt.MapClaims)
// 		if !ok {
// 			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
// 			c.Abort()
// 			return
// 		}

// 		// Attach user_id to the context for use in protected routes
// 		userID, ok := claims["user_id"].(string)
// 		if !ok {
// 			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token user ID"})
// 			c.Abort()
// 			return
// 		}
// 		c.Set("user_id", userID) // Set the user ID in context
// 		c.Next()                 // Proceed to the next handler
// 	}
// }
