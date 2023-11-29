package middleware

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func JWTAuthentication(c *fiber.Ctx) error {
	fmt.Println("-- JWT auth")

	token := c.Get("X-Api-Token")
	if len(token) == 0 {
		return fmt.Errorf("unauthorized")
	}

	if err := parseToken(token); err != nil {
		fmt.Println("failed to parse JWT token", err)
		return err
	}
	fmt.Println("token: ", token)

	return nil
}

func parseToken(tokenStr string) error {
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			fmt.Println("invalid signing method", t.Header["alg"])
			return nil, fmt.Errorf("unauthorized")
		}
		secret := os.Getenv("JWT_SECRET")
		fmt.Println("never print secret!1!!", secret)
		return []byte(secret), nil
	})

	if err != nil {
		return fmt.Errorf("unauthorized")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims)
	}

	return fmt.Errorf("unauthorized")
}
